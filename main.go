package main

import (
	"log"
	"net/http"
	"os"
	"packform-project-test/controllers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error load your configuration:", err.Error())
		return
	}

	port := os.Getenv("PORT")

	router := http.NewServeMux()

	router.HandleFunc("/orders", controllers.OrderHandler())

	if err := http.ListenAndServe(":"+port, router); err != nil { // Checked for error
		log.Fatal("The server application is error:", err.Error())
	}
}
