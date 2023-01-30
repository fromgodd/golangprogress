//Typeof in Go, Converting
//var blabla type
//fmt.printf("%T", blabla)

package main

import "fmt"

func main() {
	fmt.Println("Check the types")

	var name string = "adidas"
	var age int32 = 55

	fmt.Printf("name has type of %T\n", name)
	fmt.Printf("age has type of %T", age)

	var ageConverted = int64(age)
	fmt.Printf("ageconverted is %T", ageConverted)
}
