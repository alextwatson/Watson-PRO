package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)
	port := os.Getenv("PORT")
	if port == "" {
		port = "2005"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s\n", req.Proto, req.URL)
	fmt.Fprintln(w, "Welcome to alex thing!")
	fmt.Fprintln(w, "you are looking at my website.")
	fmt.Fprintln(w, "it's awesome I know")
	fmt.Fprintln(w, "and I'm just a kid")
	fmt.Fprintln(w, "my MOM and DAD are awesome")
	fmt.Fprintln(w, "OK so I love fruit and I like soccer and if you don't know what that is,it's a sport")
	fmt.Fprintln(w, "just saying more stuff are coming SOON")
}
