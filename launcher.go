package main

import (
	"fmt"
	"github.com/bmizerany/pat"
	"log"
	"net/http"
)

//go:generate go-bindata templates
func main() {
	mux := pat.New()
	http.Handle("/", mux)
	mux.Get("/", http.HandlerFunc(Home))
	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func GetAsset(path string) []byte {
	data, err := Asset(path)
	if err != nil {
		panic(err)
	}
	return data
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := GetAsset("templates/homepage.html.tmpl")
	fmt.Fprint(w, string(data))
}
