package server

import (
	"context"
	"gitee.com/cristiane/micro-mall-comments/pkg/code"
	"gitee.com/cristiane/micro-mall-comments/proto/micro_mall_comments_proto/comments_business"
	"gitee.com/cristiane/micro-mall-comments/service"
)

type CommentsServer struct {
}

func NewCommentsServer() *CommentsServer {
	return new(CommentsServer)
}

func (c *CommentsServer) CommentsOrder(ctx context.Context, req *comments_business.CommentsOrderRequest) (*comments_business.CommentsOrderResponse, error) {
	result := &comments_business.CommentsOrderResponse{
		Common: &comments_business.CommonResponse{
			Code: comments_business.RetCode_SUCCESS,
			Msg:  "",
		}}
	retCode := service.CommentsOrder(ctx, req)
	if retCode != code.Success {
		switch retCode {
		case code.UserOrderStateInvalid:
			result.Common.Code = comments_business.RetCode_USER_ORDER_STATE_INVALID
		case code.TransactionFailed:
			result.Common.Code = comments_business.RetCode_TRANSACTION_FAILED
		case code.UserNotExist:
			result.Common.Code = comments_business.RetCode_USER_NOT_EXIST
		case code.UserOrderCodeNotExist:
			result.Common.Code = comments_business.RetCode_USER_ORDER_NOT_EXIST
		default:
			result.Common.Code = comments_business.RetCode_ERROR
		}
		return result, nil
	}
	return result, nil
}

func (c *CommentsServer) FindShopComments(ctx context.Context, req *comments_business.FindShopCommentsRequest) (*comments_business.FindShopCommentsResponse, error) {
	result := &comments_business.FindShopCommentsResponse{
		Common: &comments_business.CommonResponse{
			Code: comments_business.RetCode_SUCCESS,
			Msg:  "",
		}}
	commentList, retCode := service.FindShopComments(ctx, req)
	if retCode != code.Success {
		switch retCode {
		case code.ErrorServer:
			result.Common.Code = comments_business.RetCode_ERROR
		default:
			result.Common.Code = comments_business.RetCode_ERROR
		}
		return result, nil
	}
	result.CommentsList = commentList
	return result, nil
}

func (c *CommentsServer) FindCommentsTags(ctx context.Context, req *comments_business.FindCommentsTagRequest) (*comments_business.FindCommentsTagResponse, error) {
	result := &comments_business.FindCommentsTagResponse{
		Common: &comments_business.CommonResponse{
			Code: comments_business.RetCode_SUCCESS,
			Msg:  "",
		},
		Tags: nil,
	}
	tagList, retCode := service.FindCommentsTags(ctx, req)
	if retCode != code.Success {
		switch retCode {
		case code.CommentTagNotExist:
			result.Common.Code = comments_business.RetCode_COMMENT_TAG_NOT_EXIST
		default:
			result.Common.Code = comments_business.RetCode_ERROR
		}
		return result, nil
	}
	result.Tags = tagList
	return result, nil
}

func (c *CommentsServer) ModifyCommentsTags(ctx context.Context, req *comments_business.ModifyCommentsTagsRequest) (*comments_business.ModifyCommentsTagsResponse, error) {
	result := &comments_business.ModifyCommentsTagsResponse{
		Common: &comments_business.CommonResponse{
			Code: comments_business.RetCode_SUCCESS,
			Msg:  "",
		},
		TagCode: "",
	}
	retCode := service.ModifyCommentsTags(ctx, req)
	if retCode != code.Success {
		switch retCode {
		case code.CommentTagNotExist:
			result.Common.Code = comments_business.RetCode_COMMENT_TAG_NOT_EXIST
		case code.CommentTagExist:
			result.Common.Code = comments_business.RetCode_COMMENT_TAG_EXIST
		default:
			result.Common.Code = comments_business.RetCode_ERROR
		}
		return result, nil
	}
	return result, nil
}
