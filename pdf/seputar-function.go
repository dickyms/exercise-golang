package main

import (
	"fmt"
)

var pl = fmt.Println

// function as parameter
// type sayHello func(string) string => untuk jadi alias
func buat(name string, sayHello func(string) string) {
	pl(sayHello(name))
}

func sayHello(name string) string {
	kata := "hello " + name + ", apa kabar ?"
	return kata
} 
// end function parameter

// variadic function
func iniFungsinya (number ...int) {
	for _, x := range number {
		pl(x)
	}
}

// function multiple return
func multiReturn(anak1 string, anak2 string, anak3 string) (helloAnak1 string, helloAnak2 string, helloAnak3 string) {
	helloAnak1 = "Hello " + anak1 + ", Bagaimana kabarnya ?"
	helloAnak2 = "Hello " + anak2 + ", Bagaimana kabarnya ?"
	helloAnak3 = "Hello " + anak3 + ", Bagaimana kabarnya ?"
	return
}

//recursive function
func factorialRecursive (nilai int) int {
	if nilai == 1 {
		return 1
	} else {
		return nilai * factorialRecursive(nilai-1)
	}
}



func main() {
	fungsinya := sayHello // function menjadi value
	buat("dicky", fungsinya)
	iniFungsinya(1, 2, 3, 4, 5, 6, 10, 2)
	
	panggilan1, panggilan2, panggilan3 := multiReturn("Ramses", "George", "Buzz") // cara menerima hasil returnnya
	
	pl(panggilan1)
	pl(panggilan2)
	pl(panggilan3)

	//function anonim
	inifuncAnonim := func (saySomething string) {
		pl(saySomething + " from anonim function")
	}
	inifuncAnonim("Apa ya")
	pl(factorialRecursive(10))
}