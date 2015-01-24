package main

import (
	"fmt"
	"github.com/bmizerany/pat"
	"github.com/elazarl/go-bindata-assetfs"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"strings"
	"os"
)

var config *Config
//go:generate go-bindata -debug templates static
func main() {
	var err error
	config, err = ParseConfig("launcher.toml")
	if err != nil {
		fmt.Printf("Error while parsing launcher.toml config: %v\n", err.Error())
		os.Exit(1)
	}
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

func ExecuteScript(name string, send chan string) {
	var script *Script
	var host *Host
	for scriptName, configScript := range config.Scripts {
		if scriptName == name {
			script = configScript
		}
	}
	if script != nil {
		for hostName, configHost := range config.Hosts {
			if hostName == script.Host {
				host = configHost
			}
		}
	}
	if script != nil && host != nil {
		err := script.Execute(host, &ChannelWriter{send})
		if err != nil {
			panic(err)
		}
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
