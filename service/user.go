package service

import (
	"errors"
	"mychat/cmd/base"
	"mychat/model"

	"github.com/sirupsen/logrus"
)

//获取数据库连接池
var Db = base.Db

type UserService struct {
}

func (u *UserService) Login(mobile, passwd string) (user model.User, err error) {
	tmp := model.User{}
	has, err := Db.Where("mobile=?", mobile).Get(&tmp)
	if err != nil {
		logrus.Error(err)
	}
	if !has {
		return tmp, errors.New("用户不存在！")
	}
	if passwd != tmp.Passwd {
		return tmp, errors.New("密码错误！")
	}
	return tmp, nil
}

func (u *UserService) Register(mobile, passwd string) (model.User, error) {
	tmp := model.User{}
	has, err := Db.Where("mobile=?", mobile).Get(&tmp)
	if err != nil {
		logrus.Error(err)
	}
	if has {
		return tmp, errors.New("该手机号已注册")
	}
	tmp.Mobile = mobile
	tmp.Passwd = passwd
	_, err = Db.Insert(tmp)
	if err != nil {
		logrus.Error(err)
	}
	return tmp, nil
}
