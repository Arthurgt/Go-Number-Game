package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func numbers() {
	var i int
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	guessNumber := random.Intn(100)
	tryCount := 0

	for {
		tryCount++
		fmt.Println("Guess the number from 1 to 100: ")
		fmt.Scan(&i)

		if i < guessNumber {
			fmt.Println("To small number")
		} else if i > guessNumber {
			fmt.Println("To high number")
		} else {
			fmt.Println("Correct number!")
			fmt.Printf("You guessed the number in: %v tries", tryCount)
			os.Exit(0)
		}
	}

}
