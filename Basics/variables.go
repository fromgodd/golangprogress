package main

import "fmt"

func main() {
	fmt.Println("Variables")

	//Vars
	//var "var_name" type

	//String variable

	var name string = "Alex"
	fmt.Println("name variable = " + name)

	//Int8
	var age int8
	fmt.Println("age variable = " + fmt.Sprint(age)) //sprint is used to convert number to string

	//Float32
	var floatvar float32 = 12
	fmt.Println("floatvar variable = ", floatvar)

}
