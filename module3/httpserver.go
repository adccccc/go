package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello World!")
	flag.Set("v", "4")
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/", healthz)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {

	for key, values := range r.Header {
		for _, v := range values {
			w.Header().Add(key, v)
		}
	}
	w.Header().Add("VERSION", os.Getenv("VERSION"))
	w.WriteHeader(http.StatusOK)

	remoteIp := strings.Split(r.RemoteAddr, ":")[0]
	fmt.Println("Client IP: " + remoteIp)
	fmt.Println("Response ResponseCode: 200")
}
