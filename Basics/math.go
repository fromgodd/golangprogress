package main

import (
	"fmt"
)

/*
basic math operations
+ - * /

to divide evenly use static variable type (float64)

% - remainder from division

math library

Ceil - round to highest (everything) 25.2 -> 26 25.0005 -> 26, 25.9999 -> 26

# Floor - lowest

Round - normal
*/

func main() {
	// var a float64 = 25.2135
	// result := math.Ceil(a)
	// fmt.Println(result)
	// var b float64 = 25.922115
	// result2 := math.Round(b)
	// fmt.Println(result2)

	fmt.Println("•CALC IN GO•")
	fmt.Println("Input operation + - * / %")

	var action string
	fmt.Scan(&action)

	fmt.Println("Input first number: ")
	var firstNumber float64
	fmt.Scan(&firstNumber)

	fmt.Println("Input second number: ")
	var secondNumber float64
	fmt.Scan(&secondNumber)

	switch {
	case action == "+":
		fmt.Println("Sum: " + fmt.Sprint(firstNumber+secondNumber))
	case action == "-":
		fmt.Println("Difference: " + fmt.Sprint(firstNumber-secondNumber))
	case action == "*":
		fmt.Println("Product: " + fmt.Sprint(firstNumber*secondNumber))
	case action == "/":
		fmt.Println("Quotient: " + fmt.Sprint(firstNumber/secondNumber))
	default:
		fmt.Println("Error")
	}

}
