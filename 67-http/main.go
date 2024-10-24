package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	PORT string
)

func main() {
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root route."))
	})
	http.HandleFunc("/health", health)
	ping := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}
	http.HandleFunc("/ping", ping)

	println("Server is up and running on the port:", PORT)
	// 0.0.0.0
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatal(err) // os.Exit(2) along with println
	} // this ensure that your server is running on 8081 port
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}
