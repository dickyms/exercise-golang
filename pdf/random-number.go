package main

import (
	"fmt"
	"math/rand"
	"time"
)

var pl = fmt.Println

func main() {
	pl(time.Now().Unix())
	rand.Seed(time.Now().UTC().Unix())
	numAcak := rand.Intn(50) + 1
	pl("random : ", numAcak)
}