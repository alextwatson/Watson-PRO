package main

import (
	"flag"
	"fmt"
)

var a string

func main() {
	g := fmt.Printf
	flag.StringVar(&a, "a", "racecar2", "defult vaule for a is racecar2")
	flag.Parse()
	x := len(a)
	fmt.Println(x)
	for i := x - 1; i > -1; i-- {
		g("%c", a[i])
	}
	fmt.Print("  ")

}
