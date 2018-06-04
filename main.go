package main

import (
	"net/http"
	"log"

	"github.com/adizim/webdemo/handlers"	
)

func main() {
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/index", handlers.IndexHandler)
	http.HandleFunc("/view/", handlers.MakeHandler(handlers.ViewHandler))
	http.HandleFunc("/edit/", handlers.MakeHandler(handlers.EditHandler))
	http.HandleFunc("/save/", handlers.MakeHandler(handlers.SaveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}