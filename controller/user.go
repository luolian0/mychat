package controller

import (
	"mychat/service"
	"mychat/utils"
	"net/http"

	"github.com/sirupsen/logrus"
)

var s service.UserService
var res utils.Resp

func Login(w http.ResponseWriter, r *http.Request) {
	mobile := r.PostFormValue("mobile")
	passwd := r.PostFormValue("password")
	user, err := s.Login(mobile, passwd)
	if err != nil {
		logrus.Warn(err)
		res.RespFail(w, err.Error())
	} else {
		logrus.Info(user.Mobile + " 登录成功！")
		res.RespOk(w, user, "操作成功！")
	}

}

func Register(w http.ResponseWriter, r *http.Request) {

	mobile := r.PostFormValue("mobile")
	passwd := r.PostFormValue("password")
	if mobile == "" || passwd == "" {
		logrus.Error("手机号或密码为空")
		res.RespFail(w, "手机号或密码为空")
	} else {
		user, err := s.Register(mobile, passwd)
		if err != nil {
			logrus.Error(err)
			res.RespFail(w, err.Error())
		} else {
			res.RespOk(w, user, "操作成功！")
		}
	}
}
