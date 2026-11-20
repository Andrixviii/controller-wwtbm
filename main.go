package main

import (
	"go-web/config"
	"go-web/controllers/categorycontroller"
	"go-web/controllers/homecontroller"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// pemangilan koneksi
	config.ConnectionDB()

	// routing path
	http.HandleFunc("/", homecontroller.Home)
	http.HandleFunc("/actions", categorycontroller.Index)
	http.HandleFunc("/actions/create", categorycontroller.Create)
	http.HandleFunc("/actions/edit", categorycontroller.Edit)
	http.HandleFunc("/actions/delete", categorycontroller.Delete)

	addr := ":" + port
	// log.Printf("Berjalan di http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
