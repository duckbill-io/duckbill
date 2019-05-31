// duckbill 一个简单的博客引擎
package main

import (
	"github.com/duckbill-io/duckbill/router"
	"log"
	"net/http"
)

func main() {
	r := router.New()
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}
	log.Println("starting duckbill ...")
	log.Fatal(server.ListenAndServe())
}
