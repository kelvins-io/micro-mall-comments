package models

import (
	"time"
)

type CommentsLogistics struct {
	Id                     int64     `xorm:"pk autoincr comment('自增ID') BIGINT"`
	LogisticsCode          string    `xorm:"not null comment('物流单号') index CHAR(40)"`
	Uid                    int64     `xorm:"not null comment('用户ID') index BIGINT"`
	FedexPackStar          int       `xorm:"comment('物流包装星级') TINYINT"`
	FedexPackContent       string    `xorm:"comment('物流包装评价') TEXT"`
	DeliverySpeedStar      int       `xorm:"comment('送货速度星级') TINYINT"`
	DeliverySpeedContent   string    `xorm:"comment('送货速度评价') TEXT"`
	DeliveryService        int       `xorm:"comment('配送服务星级') TINYINT"`
	DeliveryServiceContent string    `xorm:"comment('配送服务评价') TEXT"`
	Comment                string    `xorm:"comment('评价内容') TEXT"`
	Anonymity              int       `xorm:"not null default 0 comment('是否匿名，0-匿名，1-实名') TINYINT"`
	State                  int       `xorm:"not null default 0 comment('状态，0-有效，1-无效') TINYINT"`
	CreateTime             time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') index DATETIME"`
	UpdateTime             time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') DATETIME"`
}

type CommentsOrder struct {
	Id          int64     `xorm:"pk autoincr comment('自增ID') BIGINT"`
	CommentCode string    `xorm:"not null comment('评论code') index CHAR(40)"`
	Uid         int64     `xorm:"comment('用户ID') index BIGINT"`
	ShopId      int64     `xorm:"comment('店铺id') index BIGINT"`
	OrderCode   string    `xorm:"comment('订单code') CHAR(40)"`
	Star        int       `xorm:"comment('星级') INT"`
	Content     string    `xorm:"comment('评价类容') TEXT"`
	ImgList     string    `xorm:"comment('评价图片or视频') TEXT"`
	Anonymity   int       `xorm:"not null default 0 comment('是否匿名，0-匿名，1-实名') TINYINT"`
	State       int       `xorm:"not null default 0 comment('状态，0-有效，1-无效') TINYINT"`
	CreateTime  time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') index DATETIME"`
	UpdateTime  time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') DATETIME"`
}

type CommentsTags struct {
	Id                   int64     `xorm:"pk autoincr comment('自增ID') BIGINT"`
	TagCode              string    `xorm:"not null comment('标签code') index CHAR(40)"`
	ClassificationMajor  string    `xorm:"comment('主要分类') index VARCHAR(255)"`
	ClassificationMedium string    `xorm:"comment('中等分类') VARCHAR(255)"`
	ClassificationMinor  string    `xorm:"comment('次要分类') VARCHAR(255)"`
	Content              string    `xorm:"comment('内容') TEXT"`
	State                int       `xorm:"not null default 0 comment('状态，0-有效，1-无效') TINYINT"`
	CreateTime           time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime           time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') DATETIME"`
}
