package main

import (
	"fmt"
)

// Buffered channel
/*
	Antrian channel juga dapat digunakan pada golang dengan menambahkan nilai int untuk ukuran buffernya dan deklarasikan saat menggunakan make() pada parameter kedua
	contoh : make(chan int, 5)
	5 disana merukana ukuran buffernya
*/
func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 1
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

