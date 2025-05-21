package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func cryptoRandeSecure(max int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))

	if err != nil {
		fmt.Println(err)
	}

	return nBig.Int64()
}

func shufflePassword(shufflepass string) string {
	runes := []rune(shufflepass)

	for i := len(runes) - 1; i > 0; i-- {
		shuffle := cryptoRandeSecure(int64(i + 1))
		runes[i], runes[shuffle] = runes[shuffle], runes[i]
	}

	return string(runes)
}

func generatePassword(passwordLength int) string {
	password := ""
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	digits := "0123456789"
	special := "!@#$%^&*()_-"

	password += string(uppercase[cryptoRandeSecure(int64(len(uppercase)))])
	password += string(lowercase[cryptoRandeSecure(int64(len(lowercase)))])
	password += string(digits[cryptoRandeSecure(int64(len(digits)))])
	password += string(special[cryptoRandeSecure(int64(len(special)))])

	allChars := uppercase + lowercase + digits + special

	for i := 0; i < passwordLength-4; i++ {
		password += string(allChars[cryptoRandeSecure(int64(len(allChars)))])
	}

	return shufflePassword(password)
}

func main() {
	var passwordLength int

	fmt.Print("Entrer la longueur du mot de passe souhaité : ")
	fmt.Scanln(&passwordLength)

	if passwordLength < 4 {
		fmt.Println("Erreur : Le mot de passe doit contenir au moins 4 caractères")
		return
	}

	fmt.Println("Mot de passe générer : ", generatePassword(passwordLength))
}
