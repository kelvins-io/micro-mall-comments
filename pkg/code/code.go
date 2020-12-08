package code

import "gitee.com/kelvins-io/common/errcode"

const (
	Success               = 29000000
	ErrorServer           = 29000001
	DecimalParseErr       = 29000002
	TransactionFailed     = 29000003
	UserNotExist          = 29000005
	UserOrderCodeNotExist = 29000004
	DBDuplicateEntry      = 29000007
	CommentNotExist       = 29000006
	CommentExist          = 29000008
	CommentTagNotExist    = 29000009
	CommentTagExist       = 29000010
	UserOrderStateInvalid = 29000011
)

var ErrMap = make(map[int]string)

func init() {
	dict := map[int]string{
		Success:               "OK",
		ErrorServer:           "服务器错误",
		UserNotExist:          "用户不存在",
		DBDuplicateEntry:      "Duplicate entry",
		TransactionFailed:     "事务执行失败",
		DecimalParseErr:       "浮点数解析错误",
		UserOrderCodeNotExist: "用户订单不存在",
		CommentExist:          "评论已存在",
		CommentNotExist:       "评论不存在",
		CommentTagNotExist:    "评论标签不存在",
		CommentTagExist:       "评论标签已存在",
		UserOrderStateInvalid: "用户订单无效",
	}
	errcode.RegisterErrMsgDict(dict)
	for key, _ := range dict {
		ErrMap[key] = dict[key]
	}
}
