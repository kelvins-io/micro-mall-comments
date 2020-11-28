package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-comments/model/mysql"
	"gitee.com/cristiane/micro-mall-comments/pkg/code"
	"gitee.com/cristiane/micro-mall-comments/proto/micro_mall_comments_proto/comments_business"
	"gitee.com/cristiane/micro-mall-comments/repository"
	"gitee.com/kelvins-io/common/errcode"
	"gitee.com/kelvins-io/kelvins"
	"github.com/google/uuid"
	"strings"
	"time"
)

const sqlSelectFindCommentsTag = "tag_code,classification_major,classification_medium,classification_minor,content,state"

func FindCommentsTags(ctx context.Context, req *comments_business.FindCommentsTagRequest) (result []*comments_business.CommentsTags, retCode int) {
	retCode = code.Success
	where := map[string]interface{}{}
	if req.TagCode != "" {
		where["tag_code"] = req.TagCode
	}
	if req.ClassificationMajor != "" {
		where["classification_major"] = req.ClassificationMajor
	}
	if req.ClassificationMedium != "" {
		where["classification_medium"] = req.ClassificationMedium
	}
	tagsList, err := repository.FindCommentsTags(sqlSelectFindCommentsTag, where)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "FindCommentsTags err: %v,where: %+v", err, where)
		retCode = code.ErrorServer
		return
	}
	result = make([]*comments_business.CommentsTags, len(tagsList))
	if len(tagsList) == 0 {
		return
	}
	for i := 0; i < len(tagsList); i++ {
		tag := &comments_business.CommentsTags{
			TagCode:              tagsList[i].TagCode,
			ClassificationMajor:  tagsList[i].ClassificationMajor,
			ClassificationMedium: tagsList[i].ClassificationMedium,
			ClassificationMinor:  tagsList[i].ClassificationMinor,
			Content:              tagsList[i].Content,
			State:                false,
		}
		if tagsList[i].State == 0 {
			tag.State = true
		} else {
			tag.State = false
		}
		result[i] = tag
	}
	return
}

func ModifyCommentsTags(ctx context.Context, req *comments_business.ModifyCommentsTagsRequest) (retCode int) {
	retCode = code.Success
	if req.OpType == comments_business.OperationType_CREATE {
		tag := mysql.CommentsTags{
			TagCode:              uuid.New().String(),
			ClassificationMajor:  req.Tag.ClassificationMajor,
			ClassificationMedium: req.Tag.ClassificationMedium,
			ClassificationMinor:  req.Tag.ClassificationMinor,
			Content:              req.Tag.Content,
			State:                0,
			CreateTime:           time.Now(),
			UpdateTime:           time.Now(),
		}
		if req.Tag.State {
			tag.State = 0
		} else {
			tag.State = 1
		}
		err := repository.CreateCommentsTags([]mysql.CommentsTags{tag})
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "CreateCommentsTags err: %v, model: %+v", err, tag)
			if strings.Contains(err.Error(), errcode.GetErrMsg(code.DBDuplicateEntry)) {
				retCode = code.CommentTagExist
				return
			}
			retCode = code.ErrorServer
			return
		}
		return
	} else if req.OpType == comments_business.OperationType_UPDATE {
		findWhere := map[string]interface{}{
			"tag_code": req.Tag.TagCode,
		}
		tags, err := repository.FindCommentsTags("tag_code", findWhere)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "FindCommentsTags err: %v, findWhere: %+v", err, findWhere)
			retCode = code.ErrorServer
			return
		}
		if len(tags) == 0 {
			retCode = code.CommentTagNotExist
			return
		}
		updateWhere := map[string]interface{}{
			"tag_code": req.Tag.TagCode,
		}
		updateMaps := map[string]interface{}{}
		if req.Tag.Content != "" {
			updateMaps["content"] = req.Tag.Content
		}
		if req.Tag.State {
			updateMaps["state"] = 0
		} else {
			updateMaps["state"] = 1
		}
		_, err = repository.UpdateCommentsTag(updateWhere, updateMaps)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "UpdateCommentsTag err: %v, updateWhere: %+v,updateMaps:%+v", err, updateWhere, updateMaps)
			retCode = code.ErrorServer
			return
		}
		return
	}
	return
}
