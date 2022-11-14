package main

import (
	"fmt"
)

type Profil struct {
	alamat string
	kelurahan string
}

func changeDirectory(profil *Profil) { //tanda * berarti struct profil diambil menggunakan pointer
	profil.alamat = "tidak ada"
}

func main ()  {
	data1 := Profil{
		alamat : "Jl. Jawa",
		kelurahan : "MargaAsaih",
	}
	data2 := "salaa"
	data3 := &data2 // ketika line ini dijalankan maka akan nge print 0xc000052250 alamat dari memori
	changeDirectory(&data1) // operator & untuk tanda pake ponter
	fmt.Println(data1.alamat)
	fmt.Println(data2)
	fmt.Println(data3)
}