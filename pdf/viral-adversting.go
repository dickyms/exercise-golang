package main

import (
	"fmt"
)

func viralAdvertising(n int32) int32 {
    // Write your code here
    var sum int32 = 0
    var people int32 = 5
    
    for i := int32(0); i < n; i++ {
        nextPeople := people/2
        sum += nextPeople
        people = nextPeople*3
    }
    return sum
}

func main() {
	fmt.Println(viralAdvertising(8))
}

