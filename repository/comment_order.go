package repository

import (
	"gitee.com/cristiane/micro-mall-comments/model/mysql"
	"gitee.com/kelvins-io/kelvins"
	"xorm.io/xorm"
)

func CreateCommentOrderByTx(tx *xorm.Session, model *mysql.CommentsOrder) (err error) {
	_, err = tx.Table(mysql.TableCommentsOrder).Insert(model)
	return
}

func FindCommentOrder(sqlSelect string, where interface{}) (result []mysql.CommentsOrder, err error) {
	result = make([]mysql.CommentsOrder, 0)
	err = kelvins.XORM_DBEngine.Table(mysql.TableCommentsOrder).Select(sqlSelect).Where(where).Find(&result)
	return
}
