package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	 //deklarasi array atau slice

	 var arr1 [5] int
	 arr1[0] = 1
	 arr2 := [5]int{1,5,6,2,9}

	 pl("Index 0 ", arr2[0])
	 pl("arr length : ", len(arr2))
	 for i := 0; i < len(arr2); i++ {
		pl(arr2[i])
	 }
	 for i, v := range arr2 {
		fmt.Printf("%d : %d\n", i, v)
	 }
}