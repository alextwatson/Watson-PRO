package main

import "fmt"

func main() {
	y := "yo"
	a := fmt.Println
	for i := 0; i < 101; i++ {
		for j := 0; j < 5; j++{
			
			a(i, y)
		}
		
	}
}
