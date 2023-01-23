package main

import (
	"fmt"
	"math"
)

func main() {

	//ax^2 + bx + c = 0

	var a, b, c, x1, x2 float64

	fmt.Println("Find square of equation")
	fmt.Println("Input A ")
	fmt.Scan(&a)
	fmt.Println("Input B ")
	fmt.Scan(&b)
	fmt.Println("Input C")
	fmt.Scan(&c)

	D := (b * b) - 4*(a*c)

	if D > 0 {
		//var x1 float64
		//var x2 float64
		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)
		// fmt.Println("The equation has two roots.\nX1 -> " + fmt.Sprint(x1) + "\nX2 -> " + fmt.Sprint(x2))
		fmt.Println("X1: " + fmt.Sprint(x1) + "\nX2: " + fmt.Sprint(x2) + "\nDiscriminant = " + fmt.Sprint(D))
	} else if D == 0 {
		//var x1 float64
		//var x2 flaot64
		var x float64
		x = (-b) / (2 * a)
		fmt.Println("Equation has one root")
		fmt.Println("X: " + fmt.Sprint(x))
	} else if D < 0 {
		fmt.Println("Equation has no roots! D<0")
		fmt.Println("Discriminant is " + fmt.Sprint(D))
	}

}
