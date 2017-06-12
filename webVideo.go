package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var vid []byte

func main() {
	http.HandleFunc("/motion/notification", notification)
	http.HandleFunc("/motion/video", video)
	http.HandleFunc("/", owner)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func notification(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s\n", req.Proto, req.URL)
	println(req)
}

func video(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s\n", req.Proto, req.URL)
	vid, _ = ioutil.ReadAll(req.Body)
	println(string(vid))
}
func owner(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s\n", req.Proto, req.URL)
	fmt.Fprintln(w, "Content-Type: video/x-msvideo\n")
	fmt.Fprintln(w, vid)
}
