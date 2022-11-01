package main

import "fmt"

var pl = fmt.Println

func main() {
	var jawaban int

	pl("Masukan nomornya : ")
	fmt.Scanf("%d", &jawaban)

	switch jawaban {
	case 1 :
		pl("jawabannya adalah 1")
		break
	case 2 :
		pl("jawabannya adalah 2")
		break
	case 3 :
		pl("jawabannya adalah 3")
		break
	default :
		pl("bukan 3, 2 dan 1")
		break
	}
}