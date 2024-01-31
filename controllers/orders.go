package controllers

import (
	"fmt"
	"net/http"
)

func OrderHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getOrder(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("testing order endpoint")
}
