package main

import (
	"fmt"
	"github.com/bmizerany/pat"
	"github.com/elazarl/go-bindata-assetfs"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"strings"
)

//go:generate go-bindata -debug templates static
func main() {
	mux := pat.New()
	fs := http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "static"})
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	go h.run()
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

func ExecuteScript(send chan string) {
	script := &Script{1, "list files", "ls && sleep 1 && ls && sleep 1 && ls && sleep 1 && ls"}
	host := &Host{"127.0.0.1", "user", "password", "22"}
	err := script.Execute(host, &ChannelWriter{send})
	if err != nil {
		panic(err)
	}
}

type ChannelWriter struct {
	c chan string
}

func (self *ChannelWriter) Write(p []byte) (n int, err error) {
	message := string(p)
	message = strings.Replace(message, "\n", "<br/>", -1)
	self.c <- message
	return len(p), nil
}
