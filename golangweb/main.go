package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	aboutHandler := func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("about Page"))
	}

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/mario", handler.MarioHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Profile"))
	})


	log.Println("starting web on port 8080")
	err := http.ListenAndServe(":8080",mux)
	log.Fatal(err)
}




