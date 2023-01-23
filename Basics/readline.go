//Read line by scan
//Sprint command to convert int to string

package main

import "fmt"

func main() {

	//Variables
	var name string
	var age uint8

	//Name
	fmt.Println("What is your name?")
	fmt.Scan(&name)
	fmt.Println("Hi, " + name + "!")

	//Age
	fmt.Println("So, how old are you?")
	fmt.Scan(&age)
	fmt.Println("Wow, you are " + fmt.Sprint(age) + " years!")
}
