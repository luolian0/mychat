package main

import (
	"mychat/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/user/login", controller.Login)
	http.HandleFunc("/user/register", controller.Register)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
