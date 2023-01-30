package main

import "fmt"

//if <blabla> {function}
//else if <blabla> == blabla { blabla }

func main() {
	num := 3
	if num > 0 {
		fmt.Println("Number is greater than 0")
	} else if num < 0 {
		fmt.Println("Number is less than 0")
	} else if num == 3 {
		fmt.Println("Number equals 3")
	}
}
