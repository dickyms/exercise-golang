package main

import (
	"fmt"
	"math"
)

var pl = fmt.Println

func main() {
	pl("Abs(-10) = ", math.Abs(-10)) //nilai absolute bukan min
	pl("Pow(4,3) = ", math.Pow(4,3)) //nilai pangkat. 4 pangkat 3
	pl("Sqrt(16) = ", math.Sqrt(16)) //akar dari 16 = 4
	pl("Cbrt(8) = ", math.Cbrt(8)) // akar 3 dari 8
	pl("Ceil(4.4) = ", math.Ceil(4.4)) // pebulatan ke atas
	pl("Floor(4.4) = ", math.Floor(4.4))
	pl("Round(4.4) = ", math.Round(4.4))
	pl("Log2(8) = ", math.Log2(8))
	pl("Log10(100) = ", math.Log10(100))
	// The Log of e to the power of 2
	pl("Log(7.389) = ", math.Log(math.Exp(2)))
	pl("Max(5,4) = ", math.Max(5,4))
	pl("Min(5,4) = ", math.Min(5,4))

	//Convert 90 degress to radians
	r90 := 90 * math.Pi / 180
	d90 := r90 * (180 / math.Pi)
	fmt.Printf("%.2f radians = %.2f degress\n", r90, d90)
	pl("Sin(90) = ", math.Sin(r90))
}