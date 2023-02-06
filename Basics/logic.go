package main

import "fmt"

func main() {
	a := 3
	b := 10

	/*
		AND - &&
		OR - ||
		NOR - !=

	*/
	if a > 1 || b > 11 {
		fmt.Println("True")
	}
	if a < 10 && b < 24 {
		println("Second")
	}
	if a != 5 {
		println("hello, world")
	}
}
