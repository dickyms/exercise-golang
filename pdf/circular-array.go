package main

import (
	"fmt"
)

/*
indCon := int32(len(a))
	if indCon%k == 0 {
		return a
	} else {
		newArr := make([]int32, len(a))
		for i := int32(0); i < indCon; i++ {
			newArr[i] = a[()]
		}
		return newArr
	}
*/

func circularArrayRotation(a []int32, k int32, q []int32) []int32 {
    // Write your code here
    newArr := make([]int32, len(a))
	as := int32(len(a))
	for x, i := range q {
		fmt.Println((as-k+i)%as)
		newArr[x] = a[(as-k+i)%as]
	}
	return newArr
}


func main() {
	fmt.Println(circularArrayRotation([]int32{1, 2, 3}, 2, []int32{0, 1, 2}))
	//s := []int{1, 2, 3, 4}
	fmt.Println(1%3)
}

