package utils

import (

	"fmt"
	"math/rand"
	"time"
	"os"
)

func PrintRules(guessingCase string) {

	if guessingCase == "numbers" {

		fmt.Println("\n\t ======== Game Rules ========\n")
		fmt.Println("1. You're going to guess numbers from 10 to 20")
		fmt.Println("2. You have 10 chances to play")
		fmt.Println("3. There is no time limit \n")

	}else if guessingCase == "words" {
		
		fmt.Println("\n\t ======== Game Rules ========\n")
		fmt.Println("1. You're going to guess words ")
		fmt.Println("2. You have 10 chances to play")
		fmt.Println("3. There is no time limit")
		fmt.Println("4. You'll not be receiving hints for words to guess\n")
		fmt.Println("5. This is the most difficult mode.\n")

	}else {
		fmt.Println("\n Invalid game mode\n")
	}

}

func GetRandomNumber() int {

	// Create a new random number generator with a specific seed
	source := rand.NewSource(time.Now().UnixNano())
	randomNum := rand.New(source)

	// Define the range
	min := 10
	max := 20

	// Generate a random number within the range
	randomNumInRange := randomNum.Intn(max-min+1) + min

	return randomNumInRange
	
}

func CheckChances(chances int, gameMode string) {

	MAX_CHANCES := 10
	MIN_CHANCES := 0 
	
	if chances == MAX_CHANCES && gameMode == "numbers" { 

		PrintRules(gameMode)

	} else if chances == MAX_CHANCES && gameMode == "words" {

		PrintRules(gameMode)

	}else {

		if chances <= MIN_CHANCES {
			fmt.Println("Out of chances, Try again later.")
			os.Exit(1)
		}

	}

}