package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("put 2 numbers that equal 67")
	if os.Args[1]+os.Args[2] == 67 {
		fmt.Println("Your smart ")
	} else {
		fmt.Println("sorry man try again")
	}

}
