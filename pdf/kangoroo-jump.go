package main

import (
	"fmt"
)

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
    // Write your code here
    if x1 < x2 && v1 < v2 {
        return "NO"
    } else {
        flag := false
        atkang1 := x1
        atkang2 := x2
        for i := 0; i < 10000; i++ {
            atkang1 += v1
            atkang2 += v2
            if (atkang1 == atkang2) {
                flag = true
            }
        }
        if flag == false {
            return "NO"
        } else {
            return "YES"
        }
    }
}


func main() {
	fmt.Println(kangaroo(0, 2, 5, 1))
}

