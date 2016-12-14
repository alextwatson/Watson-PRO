package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var x string

func main() {
	flag.StringVar(&x, "x", "pokemon", "got to catch them all")
	flag.Parse()
	page, err := http.Get("http://duckduckgo.com/?q=" + x)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s/n", err.Error())
		return
	}
	io.Copy(os.Stdout, page.Body)
	page.Body.Close()
}
