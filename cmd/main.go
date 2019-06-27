package main

import (
	"net/http"

	_ "./base"
)

func main() {
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
