package main

func main() {
	for i := 1; i < 100; i++ {
		x := i % 3
		y := i % 5
		if x == 0 && y == 0 {
			println("fizzbuzz")
		} else if y == 0 {
			println("buzz")
		} else if x == 0 {
			println("fizz")
		} else {
			println(i)
		}
	}
}
