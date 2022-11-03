package main

import (
	"fmt"
)

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
    // Write your code here
    resultApple := 0
    resultOrange := 0
    for i := 0; i < len(apples); i++ {
        apples[i] += a
		if apples[i] >= s && apples[i] <= t {
            resultApple++
        }
    }
    for i := 0; i < len(oranges); i++ {
        oranges[i] += b
        if oranges[i] >= s && oranges[i] <= t {
            resultOrange++
        }
    }
    fmt.Println(resultApple)
    fmt.Println(resultOrange)
    
}


func main() {
	apples := []int32{-2, 2, 1}
	oranges := []int32{5, -6}
	countApplesAndOranges(7, 11, 5, 15, apples, oranges)
}

