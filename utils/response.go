package utils

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

type Resp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg"`
}

func (r *Resp) RespFail(w http.ResponseWriter, msg string) {
	r.response(w, -1, nil, msg)
}

func (r *Resp) RespOk(w http.ResponseWriter, data interface{}, msg string) {
	r.response(w, 0, data, msg)
}

func (r *Resp) response(w http.ResponseWriter, code int, data interface{}, msg string) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	r.Code = code
	r.Msg = msg
	r.Data = data
	result, err := json.Marshal(r)
	if err != nil {
		logrus.Error(err)
	}
	_, err = w.Write(result)
	if err != nil {
		logrus.Error(err)
	}
}
