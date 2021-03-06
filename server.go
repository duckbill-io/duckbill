// duckbill 一个简单的博客引擎
package main

import (
	"github.com/duckbill-io/duckbill/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.DefaultRouter()
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}
	log.Println("starting duckbill ...")
	log.Fatal(server.ListenAndServe())
}
