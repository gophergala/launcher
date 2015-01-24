package main

import (
	"fmt"
	"github.com/bmizerany/pat"
	"github.com/elazarl/go-bindata-assetfs"
	"code.google.com/p/go.net/websocket"
	"log"
	"net/http"
)

//go:generate go-bindata -debug templates static
func main() {
	mux := pat.New()
	fs := http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "static"})
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/ws", websocket.Handler(wsHandler))

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
