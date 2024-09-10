package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandom6DigitNumber() int {
	rand.New(rand.NewSource(time.Now().UnixNano())) // Seed the random number generator
	return rand.Intn(900000) + 100000               // Generate a number between 100000 and 999999
}


func AttachRandomNumber(txID string) string {
	randomNumber := GenerateRandom6DigitNumber()
	return fmt.Sprintf("%s-%06d", txID, randomNumber) // Format as TXID-RNNNNNN
}