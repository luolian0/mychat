package model

import "time"

type User struct {
	Id       int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`     //id  主键自增
	Mobile   string    `xorm:"unique varchar(20)" form:"mobile" json:"mobile"` // 手机号
	Passwd   string    `xorm:"varchar(40)" form:"passwd" json:"-"`             // 密码
	Avatar   string    `xorm:"varchar(150)" form:"avatar" json:"avatar"`       // 头像
	Sex      string    `xorm:"varchar(2)" form:"sex" json:"sex"`               // 性别
	Nickname string    `xorm:"varchar(20)" form:"nickname" json:"nickname"`    // 昵称
	Salt     string    `xorm:"varchar(10)" form:"salt" json:"-"`               // 随机数  用于密码加密
	Online   int       `xorm:"int(10)" form:"online" json:"online"`            // 是否在线 1在线  0离线
	Token    string    `xorm:"varchar(40)" form:"token" json:"token"`          // token  chat简单版就放在数据库
	Memo     string    `xorm:"varchar(140)" form:"memo" json:"memo"`           // 自我描述
	CreateAt time.Time `xorm:"datetime" form:"create_at" json:"createAt"`      // 创建时间
}
