package server

import (
	"humoshop/api/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartRouter() {
	router := mux.NewRouter()
	//router.HandleFunc("/product/{product_name}", handlers.ProductHandler)
	//router.HandleFunc("/category/{category_name}", handlers.CategoryHandler)

	//роуты которые относятся к user-y
	router.HandleFunc("/user", handlers.GetUser).Methods("GET")
	router.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/user", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user", handlers.DeleteUser).Methods("DELETE")

	//роуты которые относятся к product
	router.HandleFunc("/product", handlers.ProductGET).Methods("GET")
	router.HandleFunc("/product", handlers.ProductCreate).Methods("POST")
	router.HandleFunc("/product", handlers.ProductUpdate).Methods("PUT")
	router.HandleFunc("/product", handlers.ProductDelete).Methods("DELETE")

	//роуты которые относятся к category
	router.HandleFunc("/category", handlers.CategoryGET).Methods("GET")

	// main rout
	router.HandleFunc("/", handlers.HomeHandler)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Println("SERVER listing ERROR")
	}
}
