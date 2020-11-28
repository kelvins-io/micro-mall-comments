package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-comments/model/args"
	"gitee.com/cristiane/micro-mall-comments/model/mysql"
	"gitee.com/cristiane/micro-mall-comments/pkg/code"
	"gitee.com/cristiane/micro-mall-comments/pkg/util"
	"gitee.com/cristiane/micro-mall-comments/proto/micro_mall_comments_proto/comments_business"
	"gitee.com/cristiane/micro-mall-comments/proto/micro_mall_order_proto/order_business"
	"gitee.com/cristiane/micro-mall-comments/repository"
	"gitee.com/kelvins-io/kelvins"
	"github.com/google/uuid"
	"strings"
	"time"
)

func CommentsOrder(ctx context.Context, req *comments_business.CommentsOrderRequest) (retCode int) {
	retCode = code.Success
	if req.Uid > 0 {
		retCode = createOrderCommentsInspect(ctx, req.Uid, req.OrderInfo.ShopId, req.OrderInfo.OrderCode)
		if retCode != code.Success {
			return
		}
	} else {
		retCode = code.UserNotExist
		return
	}
	txCode := uuid.New().String()
	anonymity := 0
	if !req.Anonymity {
		anonymity = 1
	}
	orderComment := &mysql.CommentsOrder{
		CommentCode: txCode,
		Uid:         req.Uid,
		ShopId:      req.OrderInfo.ShopId,
		OrderCode:   req.OrderInfo.OrderCode,
		Star:        int(req.OrderInfo.StarLevel),
		Anonymity:   anonymity,
		Content:     req.OrderInfo.Content,
		ImgList:     strings.Join(req.OrderInfo.ImgList, "|"),
		State:       0,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}
	logisticsComment := &mysql.CommentsLogistics{
		LogisticsCode:          txCode,
		Uid:                    req.Uid,
		FedexPackStar:          int(req.LogisticsInfo.FedexPack),
		FedexPackContent:       strings.Join(req.LogisticsInfo.FedexLabel, "|"),
		DeliverySpeedStar:      int(req.LogisticsInfo.DeliverySpeed),
		DeliverySpeedContent:   strings.Join(req.LogisticsInfo.DeliverySpeedLabel, "|"),
		DeliveryService:        int(req.LogisticsInfo.DeliveryService),
		DeliveryServiceContent: strings.Join(req.LogisticsInfo.DeliveryServiceLabel, "|"),
		Comment:                req.LogisticsInfo.Comment,
		State:                  0,
		CreateTime:             time.Now(),
		UpdateTime:             time.Now(),
	}
	tx := kelvins.XORM_DBEngine.NewSession()
	err := tx.Begin()
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "CommentsOrder Begin err: %v", err)
		retCode = code.ErrorServer
		return
	}
	err = repository.CreateCommentLogisticsByTx(tx, logisticsComment)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			kelvins.ErrLogger.Errorf(ctx, "CreateCommentLogisticsByTx Rollback err: %v, model: %+v", errRollback, logisticsComment)
		}
		kelvins.ErrLogger.Errorf(ctx, "CreateCommentLogisticsByTx err: %v, model: %+v", err, logisticsComment)
		retCode = code.ErrorServer
		return
	}
	err = repository.CreateCommentOrderByTx(tx, orderComment)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			kelvins.ErrLogger.Errorf(ctx, "CreateCommentOrderByTx Rollback err: %v, model: %+v", errRollback, orderComment)
		}
		kelvins.ErrLogger.Errorf(ctx, "CreateCommentOrderByTx err: %v, model: %+v", err, orderComment)
		retCode = code.ErrorServer
		return
	}
	err = tx.Commit()
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "CommentsOrder Commit err: %v", err)
		retCode = code.TransactionFailed
		return
	}

	return
}

func createOrderCommentsInspect(ctx context.Context, uid, shopId int64, orderCode string) (retCode int) {
	retCode = code.Success
	serverName := args.RpcServiceMicroMallOrder
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return code.ErrorServer
	}
	defer conn.Close()
	orderClient := order_business.NewOrderBusinessServiceClient(conn)
	orderReq := &order_business.InspectShopOrderRequest{
		Uid:       uid,
		ShopId:    shopId,
		OrderCode: orderCode,
	}
	orderRsp, err := orderClient.InspectShopOrder(ctx, orderReq)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "InspectShopOrder %v,err: %v", serverName, err)
		return code.ErrorServer
	}
	if orderRsp.Common.Code == order_business.RetCode_SUCCESS {
		return
	}
	kelvins.ErrLogger.Errorf(ctx, "InspectShopOrder %v,req: %+v, rsp: %+v", serverName, orderReq, orderRsp)
	switch orderRsp.Common.Code {
	case order_business.RetCode_ORDER_STATE_INVALID:
		retCode = code.UserOrderStateInvalid
	case order_business.RetCode_ORDER_NOT_EXIST:
		retCode = code.UserOrderCodeNotExist
	default:
		retCode = code.ErrorServer
	}
	return
}

const sqlSelectFindShopComment = "shop_id,order_code,star,content,img_list,comment_code"

func FindShopComments(ctx context.Context, req *comments_business.FindShopCommentsRequest) (result []*comments_business.OrderCommentsInfo, retCode int) {
	retCode = code.Success
	result = make([]*comments_business.OrderCommentsInfo, 0)
	where := map[string]interface{}{}
	if req.Uid > 0 {
		where["uid"] = req.Uid
	}
	if req.ShopId > 0 {
		where["shop_id"] = req.ShopId
	}
	orderCommentList, err := repository.FindCommentOrder(sqlSelectFindShopComment, where)
	result = make([]*comments_business.OrderCommentsInfo, len(orderCommentList))
	if len(orderCommentList) == 0 {
		return
	}
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "FindCommentOrder err: %v, where: %+v", err, where)
		retCode = code.ErrorServer
		return
	}
	for i := 0; i < len(orderCommentList); i++ {
		row := orderCommentList[i]
		info := &comments_business.OrderCommentsInfo{
			ShopId:    row.ShopId,
			OrderCode: row.OrderCode,
			StarLevel: comments_business.StarLevel(row.Star),
			Content:   row.Content,
			ImgList:   strings.Split(row.ImgList, "|"),
			CommentId: row.CommentCode,
		}
		result[i] = info
	}

	return
}
