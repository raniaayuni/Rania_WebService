package main

import (
	"fmt"
	"log"
	"net/http"
	"uas_webser/controller/auth"
	"uas_webser/controller/order"
	"uas_webser/controller/product"
	"uas_webser/database"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	database.InitDB()
	fmt.Println("Hello World")

	router := mux.NewRouter()

	router.HandleFunc("/regis", auth.Registration).Methods("POST")
	router.HandleFunc("/login", auth.Login).Methods("POST")
	//Router handler product
	router.HandleFunc("/products", product.GetProduct).Methods("GET")
	router.HandleFunc("/products", auth.JWTAuth(product.PostProduct)).Methods("POST")
	router.HandleFunc("/products/{id}", auth.JWTAuth(product.PutProduct)).Methods("PUT")
	router.HandleFunc("/products/{id}", auth.JWTAuth(product.DeleteProduct)).Methods("DELETE")
	//Router handler Album
	router.HandleFunc("/orders", order.GetOrder).Methods("GET")
	router.HandleFunc("/orders", auth.JWTAuth(order.PostOrder)).Methods("POST")
	router.HandleFunc("/orders/{id}", auth.JWTAuth(order.PutOrder)).Methods("PUT")
	router.HandleFunc("/orders/{id}", auth.JWTAuth(order.DeleteOrder)).Methods("DELETE")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5500"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		Debug:          true,
	})

	handler := c.Handler(router)

	fmt.Println("Server is running on http://localhost:2000")
	log.Fatal(http.ListenAndServe(":2000", handler))

}
