//Read line by scan
//fmt.Scan(&blabla) - & is pointer

package main

import "fmt"

func main() {

	//Variables
	var name string
	var age uint8

	//Change variables type in another var
	var x int8 = 2
	var y int64 = int64(x)
	fmt.Println(y)

	//Name
	fmt.Println("What is your name?")
	fmt.Scan(&name)
	fmt.Println("Hi, " + name + "!")

	//Age
	fmt.Println("So, how old are you?")
	fmt.Scan(&age)
	fmt.Println("Wow, you are " + fmt.Sprint(age) + " years old!")
}
