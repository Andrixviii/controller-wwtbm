package homecontroller

import (
	"html/template"
	"net/http"
	"os"
)

func Home(w http.ResponseWriter, r *http.Request) {
	asciiArt, _ := os.ReadFile("view/assets/ascii-art.txt")

	temp, err := template.ParseFiles("view/home/index.html")
	if err != nil {
		panic(err)
	}

	data := map[string]any{
		"AsciiArt": string(asciiArt),
	}
	temp.Execute(w, data)
}
