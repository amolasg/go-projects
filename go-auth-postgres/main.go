package main

import (
	"go-authentication/api/controllers"
	"log"
	"net/http"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	log.Println("Server started on: http://localhost:8080")

	fs := http.FileServer(http.Dir("asset/"))
	http.Handle("/asset/", http.StripPrefix("/asset/", fs))

	http.HandleFunc("/", controllers.Login)
	http.HandleFunc("/loginsubmit", controllers.LoginProcess)
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/registeruser", controllers.RegisterUser)
	http.HandleFunc("/success", controllers.Success)

	http.ListenAndServe(":8080", nil)
}
