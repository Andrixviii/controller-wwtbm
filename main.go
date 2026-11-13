package main

import (
	"go-web/config"
	"go-web/controllers/categorycontroller"
	"go-web/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {

	// pemangilan koneksi
	config.ConnectionDB()

	// routing path
	http.HandleFunc("/", homecontroller.Home)
	http.HandleFunc("/actions", categorycontroller.Index)
	http.HandleFunc("/actions/create", categorycontroller.Create)
	http.HandleFunc("/actions/edit", categorycontroller.Edit)
	http.HandleFunc("/actions/delete", categorycontroller.Delete)

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)

}
