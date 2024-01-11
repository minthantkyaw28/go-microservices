package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	//chi router 
	router := chi.NewRouter()

	//Chi Logger 
	router.Use(middleware.Logger)
	
	//api routes 
	router.Get("/hello",basicHandler)

	//actual server
	server := &http.Server{
		Addr: ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil{
		fmt.Println("Failed to listen to the server", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World !"))
}
