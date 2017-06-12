package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/y", hello)
	http.HandleFunc("/alex", hey)
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

	fmt.Fprintln(w, "<HTML>")
	fmt.Fprintln(w, "<FORM METHOD=\"LINK\" ACTION=\"alex\"><INPUT TYPE=\"submit\" VALUE=\"Clickable Button\"></FORM>")

	game := ""
	rand.Seed(time.Now().UTC().UnixNano())
	num := rand.Intn(3)
	if num == 0 {
		game = "noop\n <b style='color:springGreen'> HA HA HA HA HA HA HA HA HA HA HA</b>"
	} else if num == 1 {
		game = "yup"
	} else if num == 2 {
		game = "probobly"
	} else {
		game = "error"
	}
	fmt.Fprintln(w, game)
	fmt.Fprintln(w, "</HTML>")
}

func hey(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s\n", req.Proto, req.URL)
	fmt.Fprintln(w, "<HTML>")
	fmt.Fprintln(w, "ya boy alex did it")
}
