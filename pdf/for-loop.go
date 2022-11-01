package main

import (
	"fmt"
	"time"
	"math/rand"
	"bufio"
	"os"
	"log"
	"strconv"
	"strings"
)

var pl = fmt.Println

func main() {
	//for loop
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(100) + 1
	for true {
		fmt.Print("Guess a number between 0 and 100 :")
		pl("Random Number is :", randNum)
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if iGuess > randNum {
			pl("Pick a Lower Value")
		} else if iGuess < randNum {
			pl("Pick a Higher Value")
		} else {
			pl("You Guessed it")
			break
		}
	}

	//basic
	for x := 5; x >= 1; x-- {
		pl(x)
	}

	//for loop untuk array
	arrNums := []int{1,2,3,4,5}
	for _, num := range arrNums {
		pl(num)
	}
}