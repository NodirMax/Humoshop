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
	routerUser := router.PathPrefix("/user").Subrouter()
	routerUser.HandleFunc("/", handlers.GetUser).Methods("GET")
	routerUser.HandleFunc("/", handlers.CreateUser).Methods("POST")
	routerUser.HandleFunc("/", handlers.UpdateUser).Methods("PUT")
	routerUser.HandleFunc("/", handlers.DeleteUser).Methods("DELETE")

	//роуты которые относятся к product
	routerProduct := router.PathPrefix("/product").Subrouter()
	routerProduct.HandleFunc("/", handlers.ProductGET).Methods("GET")
	routerProduct.HandleFunc("/", handlers.ProductCreate).Methods("POST")
	routerProduct.HandleFunc("/", handlers.ProductUpdate).Methods("PUT")
	routerProduct.HandleFunc("/", handlers.ProductDelete).Methods("DELETE")

	//роуты которые относятся к category
	routerCategory := router.PathPrefix("/category").Subrouter()
	routerCategory.HandleFunc("/", handlers.CategoryGET).Methods("GET")
	routerCategory.HandleFunc("/", handlers.CategoryCreate).Methods("POST")
	routerCategory.HandleFunc("/", handlers.CategoryUpdate).Methods("PUT")
	routerCategory.HandleFunc("/", handlers.CategoryDelete).Methods("DELETE")

	// main rout
	router.HandleFunc("/", handlers.HomeHandler)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Println("SERVER listing ERROR")
	}
}
