package main

import (
	"github.com/vstanchev/wafty/lib"
	"log"
	"net/http"
)


func handleProxyRequest(config lib.WafConfig) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		lib.ServeReverseProxy(config, res, req)
	}
}

func main() {
	config := lib.LoadConfig("config.toml")
	http.HandleFunc("/", handleProxyRequest(config))

	log.Printf("Listening on %s\n", config.ListenAddress)
	if err := http.ListenAndServe(config.ListenAddress, nil); err != nil {
		panic(err)
	}
}
