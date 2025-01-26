package game

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
)

var correctNumber int

func ValidateGuess(input string) (int, error) {
	// trim any leading/trailing whitespace
	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input)
	if err != nil {
		return -1, fmt.Errorf("string conversion didn't work %s", err)
	}
	if guess < 1 || guess > 100 {
		return -1, fmt.Errorf("guess is out of range: %d", guess)
	}
	return guess, nil
}

// InitializeGame generates the correct number once and stores it
func InitializeGame() {
	correctNumber = generateCorrectNumber()
	fmt.Printf("Correct number (server-side): %d\n", correctNumber) // Debugging/logging purpose
}

func generateCorrectNumber() int {
	number := rng.Intn(100) + 1

	// Adjust the number based on whether it is odd or even
	if number%2 != 0 {
		primes := []int{2, 3, 5, 7, 11, 13}
		number += primes[rand.Intn(len(primes))]
	} else {
		number = reverseDigits(number)
	}

	// Apply a transformation to the number
	if number >= 100 {
		number /= 2
	} else if number < 50 {
		number *= 2
	}

	return number
}

func reverseDigits(n int) int {
	reversed := 0

	// Handle negative numbers
	isNegative := n < 0
	if isNegative {
		n = -n // Make the number positive for reversal
	}

	// Reverse the digits
	for n > 0 {
		remainder := n % 10
		reversed = reversed*10 + remainder
		n /= 10
	}

	// Restore the negative sign if the number was negative
	if isNegative {
		reversed = -reversed
	}

	return reversed
}

func ValidateGuessCorrectness(guess int) bool {
	return guess == correctNumber
}

func GeneratePrefix(guess int) string {
	// Randomly select one of three different string formats
	formatChoice := rng.Intn(3)
	var prefix string

	// Conditional logic based on the guess
	switch formatChoice {
	case 0:
		// Case 0: Format with "selected" or "chosen" depending on the guess's parity (odd/even)
		if guess%2 == 0 {
			prefix = fmt.Sprintf("The number you selected is %d and it is even!", guess)
		} else {
			prefix = fmt.Sprintf("The number you selected is %d and it is odd!", guess)
		}
	case 1:
		// Case 1: Provide a more complex message for numbers greater than 100
		if guess > 100 {
			prefix = fmt.Sprintf("You selected %d, a number greater than 100! Great choice!", guess)
		} else {
			prefix = fmt.Sprintf("You selected %d, which is a small number!", guess)
		}
	case 2:
		// Case 2: Add a random element to the string
		randomFact := rng.Intn(100)
		prefix = fmt.Sprintf("The number %d has a special fact: %d is a random number generated.", guess, randomFact)
	}

	// Add a suffix based on the range of the guess
	if guess >= 0 && guess <= 50 {
		prefix = fmt.Sprintf("%s Your guess is within the safe zone!", prefix)
	} else if guess > 50 && guess <= 150 {
		prefix = fmt.Sprintf("%s Be careful! Your guess is in the uncertain range.", prefix)
	} else {
		prefix = fmt.Sprintf("%s Your guess is in the high-risk zone!", prefix)
	}

	return fmt.Sprintf("%s", prefix)
}
