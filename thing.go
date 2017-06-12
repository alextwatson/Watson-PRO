package main

import (
	"flag"
	"fmt"
)

func main() {
	var flagB = flag.Int("n1", 0, "123flag")
	var flagA = flag.Int("n2", 0, "321flag")
	var flagC = flag.Int("n3", 0, "321flag")
	flag.Parse()
	fmt.Println("2=", *flagB)
	fmt.Println("3=", *flagC)
	fmt.Println("1=", *flagA)
	fmt.Println(*flagB + *flagA + *flagC)
}
