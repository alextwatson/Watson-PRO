package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the play")
	now := time.Now()
	fmt.Println("The time now=", time.Now())
	then := now.AddDate(0, 0, -1)

	fmt.Println("then:", then)

	fmt.Println("alex is good")

}
