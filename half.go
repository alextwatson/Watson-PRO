package main

import "flag"

func main() {
	var n = flag.Int("n", 0, "123flag")
	flag.Parse()
	num, isEven := half(*n)
	println("num=", num, " isEven=", isEven)
}

func half(n int) (int, bool) {
	if n%2 == 0 {
		return n / 2, true
	} else {
		return n / 2, false
	}

}
