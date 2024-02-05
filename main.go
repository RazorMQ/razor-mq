package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/RazorMQ/razor-mq/hub"
)

func main() {

	var port int
	if len(os.Args) > 1 {
		if p, err := strconv.Atoi(os.Args[0]); err == nil {
			port = p
		}
	} else {
		port = 8080
	}

	addr := flag.String("addr", fmt.Sprintf(":%d", port), "http service address")

	hub := hub.NewHub()
	go hub.Start()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hub.HandlePeer(w, r)
	})

	server := &http.Server{
		Addr:              *addr,
		ReadHeaderTimeout: 3 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Couldnt start broker: ", err)
	}
}
