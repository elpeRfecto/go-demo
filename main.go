package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("Enter how many passwords you want to generate: ")
	var inputPass int
	fmt.Scanln(&inputPass)
	res := make([]string, inputPass)
	for i := 0; i < inputPass; i++ {
		res[i] = generatePassword(userInput())
	}
	fmt.Println(res)
}

func generatePassword(length int) string{
	var possibleChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	res := make([]rune, length)
	for i := range res {
		res[i] = possibleChars[rand.IntN(len(possibleChars))]
	}
	return string(res)
}

func userInput() int{
	fmt.Println("Enter a password length: ")
	var input int
	fmt.Scanln(&input)
	return input
}
