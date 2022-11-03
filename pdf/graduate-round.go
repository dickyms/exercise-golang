package main

import (
	"fmt"
)

func hitungNilai(grades []int32) []int32 {
	result := make([]int32, grades[0])
	for i := 0; i < len(result); i++ {
		result[i] = grades[i+1]
	}
	for s, x := range result {
		multiple5 := ((x/5)+1)*5
		if multiple5 > 38 {
			new := multiple5 - x
			if new < 3 {
				result[s] = multiple5
			}
		} 
	}
	return result
}

func main() {
	arr := []int32{4, 73, 67, 38, 33}
	arr2 := hitungNilai(arr)
	fmt.Println(len(arr))
	fmt.Println(arr2)
}