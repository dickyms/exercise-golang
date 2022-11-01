package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	// %d : integer
	// %c : Character
	// %f : float , %.2f untuk set angka dibelakang koma
	// %t : Boolean 
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses based on data type
	// %T : Type of supplied value

	fmt.Printf("%s %d %c %f %t %o %x \n", 
		"stuff", 1, 'A', 3.14, true, 1, 1)
}