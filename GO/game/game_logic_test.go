package game

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateGuess(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		err      bool
	}{
		{"10", 10, false},
		{"007", 7, false},
		{"81", 81, false},
		{" 10  ", 10, false},
		{"$", -1, true},
		{"-15", -1, true},
		{" ", -1, true},
	}

	for _, test := range tests {
		guess, err := ValidateGuess(test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, guess)
		}
	}
}

func TestValidateGuessCorrectness(t *testing.T) {
	correctNumber = 42

	tests := []struct {
		guess    int
		expected bool
	}{
		{42, true},
		{10, false},
		{100, false},
	}

	for _, test := range tests {
		isCorrect := ValidateGuessCorrectness(test.guess)
		assert.Equal(t, test.expected, isCorrect)
	}

	isCorrect := ValidateGuessCorrectness(10)
	assert.Equal(t, isCorrect, false)
}

func TestInitializeGame(t *testing.T) {
	// Mock the random number generator for testing
	rng = rand.New(rand.NewSource(1))
	InitializeGame()
	assert.Equal(t, 56, correctNumber) // Expected value based on the fixed seed
}

func TestGenerateCorrectNumber(t *testing.T) {
	// Mock the random number generator for testing
	rng = rand.New(rand.NewSource(1))
	number := generateCorrectNumber()
	assert.Equal(t, 56, number) // Expected value based on the fixed seed
}

func TestReverseDigits(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{123, 321},
		{-123, -321},
		{100, 1},
		{-100, -1},
	}

	for _, test := range tests {
		result := reverseDigits(test.input)
		assert.Equal(t, test.expected, result)
	}
}

func TestGeneratePrefix(t *testing.T) {
	// Mock the random number generator for testing
	rng = rand.New(rand.NewSource(1))

	tests := []struct {
		guess    int
		expected string
	}{
		{42, "The number 42 has a special fact: 87 is a random number generated. Your guess is within the safe zone!"},
	}

	for _, test := range tests {
		result := GeneratePrefix(test.guess)
		assert.Equal(t, test.expected, result)
	}
}
