package handler

import (
	"fmt"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pong")
}

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Approved")
}
