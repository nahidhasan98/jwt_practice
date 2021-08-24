package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nahidhasan98/jwt_practice/handler"
)

func main() {
	fmt.Println("Program is running...")

	http.HandleFunc("/home", handler.Home)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/refresh", handler.Refresh)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
