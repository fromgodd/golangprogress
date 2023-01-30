/* MAIN TYPES



string

bool

int -
	int int8 int16 int32 int64

	uint uint8 uint16 uint32 uint64 uinitptr

	byte - alias for uint8

	rune - alias for int32

	float32 float64

	complex64 complex128

*/

/*

Strings

str := "hello"
str := 'multiline
string'


Numbers
num := 3 //int
num := 3. //float64
num := 3 + 4i //complex128
num := byte('a') (alias for uint8)


Get the type
fmt.Printf("%T\n", name)

%T - Shows type of variable (needs Printf)

TODO: *CHECK GO FMT LIBRARY FOR MORE*



*/

package main

import "fmt"

func main() {
	var name = "Alex"
	var age = 12
	var isCool = false //bool type
	fmt.Println(name, age, isCool)
	fmt.Printf("%T %T %T\n", name, age, isCool)
}
