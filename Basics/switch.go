/*Conditions. Switch*/

package main

import "fmt"

func main() {
	name := "John"
	// if name == "Kate" {
	// 	fmt.Println("it is not john!")
	// }

	/*
		switch <var> {
			case "something":
				condition

			default: // like in C++
		}
	*/

	switch name {
	case "Jordan":
		fmt.Println("Name is JORDAN")
	case "Kate":
		fmt.Println("Name is KATE")
	default:
		fmt.Println("Name is Jordan")
	}

}
