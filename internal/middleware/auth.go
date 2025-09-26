package middleware

import (
	"fmt"
	"net/http"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Authentication middleware")
}
