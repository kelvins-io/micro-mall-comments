package repository

import (
	"gitee.com/cristiane/micro-mall-comments/model/mysql"
	"xorm.io/xorm"
)

func CreateCommentLogisticsByTx(tx *xorm.Session, model *mysql.CommentsLogistics) (err error) {
	_, err = tx.Table(mysql.TableCommentsLogistics).Insert(model)
	return
}
