package main

import (
	"fmt"
	"math"
)

/*
basic math operations
+ - * /

to divide evenly use static variable type (float64)


% - remainder from division


math library

Ceil - round to highest (everything) 25.2 -> 26 25.0005 -> 26, 25.9999 -> 26

Floor


*/

func main() {
	var a float64 = 25.2135
	result := math.Ceil(a)
	fmt.Println(result)
	var b float64 = 25.322115
	result2 := math.Floor(b)
	fmt.Println(result2)

}
