package handler

import (
	"fmt"
	"net/http"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	//not done yet

	fmt.Fprintln(w, "Welcome to homepage")
}
