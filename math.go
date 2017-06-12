package main

import (
	"flag"
	"fmt"
	"math"
)

var x int
var y int

func main() {
	flag.IntVar(&x, "x", 12, "defult vaule for x is 12")
	flag.IntVar(&y, "y", 17, "defult vaule for y is 17")
	flag.Parse()
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	fmt.Println(x, y, z)
}
