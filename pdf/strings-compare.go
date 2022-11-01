package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println

func main() {
	var result = strings.Compare("dicky", "c++") // bernilai 1
	var result2 = strings.Compare("dicky", "dicky") // bernilai 0 
	var result3 = strings.Compare("dicky", "sSs") // bernilai -1
	pl("ini jawabannya : ", result)
	pl("ini jawaban ke 2 : ", result2)
	pl("ini jawaban ke 3 : ", result3)
}