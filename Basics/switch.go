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

	/*
		If we want check if number is lower or higher than X in switch case, then we shouldn't write <var> after switch word. Directly open case and condition.

		switch {
			number > 10 {
				...
			}
		}

	*/
	//TODO: CHECK FALLTHROUGH
	number := 2
	switch {
	case number > 5:
		fmt.Println("Number is higher than 5")
		fallthrough
	case number < 5:
		fmt.Println("Number is lower than 5")
	}
}
