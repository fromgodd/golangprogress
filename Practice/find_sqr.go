package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Println("Find roots of square equation")

	fmt.Println("Enter A: ")
	fmt.Scan(&a)

	fmt.Println("Enter B: ")
	fmt.Scan(&b)

	fmt.Println("Enter C: ")
	fmt.Scan(&c)

	D := (b * b) - 4*(a*c)
	fmt.Println("Discriminant is " + fmt.Sprint(D))

	if D > 0 {
		var x1 float64
		var x2 float64
		fmt.Println("Equation has 2 roots")
		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)
		fmt.Println("X1 = " + fmt.Sprint(x1))
		fmt.Println("X2 = " + fmt.Sprint(x2))
	} else if D == 0 {
		var x float64
		x = (-b) / (2 * a)
		fmt.Println("Equation has single root")
		fmt.Println("X = " + fmt.Sprint(x))

	} else if D < 0 {
		fmt.Println("Equation has no solution!")
	}
}
