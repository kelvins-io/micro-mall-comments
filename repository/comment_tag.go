package repository

import (
	"gitee.com/cristiane/micro-mall-comments/model/mysql"
	"gitee.com/kelvins-io/kelvins"
)

func FindCommentsTags(sqlSelect string, where interface{}) (result []mysql.CommentsTags, err error) {
	result = make([]mysql.CommentsTags, 0)
	err = kelvins.XORM_DBEngine.Table(mysql.TableCommentsTags).Select(sqlSelect).Where(where).Find(&result)
	return
}

func CreateCommentsTags(models []mysql.CommentsTags) (err error) {
	_, err = kelvins.XORM_DBEngine.Table(mysql.TableCommentsTags).Insert(models)
	return
}

func UpdateCommentsTag(where, maps map[string]interface{}) (rowAffected int64, err error) {
	return kelvins.XORM_DBEngine.Table(mysql.TableCommentsTags).Where(where).Update(maps)
}
