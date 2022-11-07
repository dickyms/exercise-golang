package main
//Belajar GoRoutines
/*
    *concurrency dan parallel programming
    Parallel programming adalah memecahkan suatu masalah dengan membaginya jadi kecil
    dan dijalankan secara bersamaan pada waktu yang bersamaan

    *Process vs Thread
    Process adalah sebuah eksekusi program sedangkan thread adalah segmen dari process itu
    Process mengkonsumsu memory besar sedangkan thread menggunakan memory kecil
    Porcess saling terisolasi dengan process lain dan thread bisa saling berhubungan jika dalam process yang sama
    Process lama untuk dijalankan dihentikan sedangkan thread cepat untuk dijalankan dan dihentikan

    *Perbedaan paralel dan concurrency adalah dalam melaksanakan pekerjaan. Paralel dilakukan secara bersamaan sedangkan concurrency dilakukan secara bergantian
    paralel (> thread), concurrency (< thread). 

    *CPU-bound dan I/O Bound
    Aplikasi golang apabila diterapkan untuk CPU bound, akan sulit atau kurang bermanfaat
    algoritma yang berjalan membutuhkan kecepatan CPU
    I/O Bound maksudnya algoritma atau aplikasi yang sangat tergantung dengan kecepatan input output devices yang digunakan. Contohnya aplikasi seperti membaca data dari file, database dan lain-lain.
    
    Goroutine adalah sebuah thread ringan yang dikelila oleh Go Runtime
    Ukurannya sekitar 2kb, jauh lebih ringan dari thread 1mb atau 1000kb
    goroutine berjalan secara concurrent. go routine akan dijalankan dalam thread secara concurrent. Goroutine dijalankan oleg Go Scheduler dalam thread, dimana jumlah thread nya sebanyak GOMAXPROCS (sejumlah core CPU)

    //cara kerja Go Scheduler
    
*/

import (
    "testing"
    "fmt"
    "time"
)

func RunHelloWorld() {
    fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
    go RunHelloWorld()
    fmt.Println("UPS")

    time.Sleep(10 * time.Nanosecond)
}

func DisplayNumber(number int ) {
    fmt.Println("Display ", number)
}

func TestManyGoroutine(t *testing.T) {
    for i := 0 ; i < 10000; i++ {
       go DisplayNumber(i)
    }

    time.Sleep(5 * time.Second)
}