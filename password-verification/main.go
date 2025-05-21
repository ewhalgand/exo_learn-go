package main

import "fmt"

func isSpecialChars(char rune) bool {
	special := "!@#$%^&*()-_=+[]{}<>?/\\|"

	for _, specialChar := range special {
		if char == specialChar {
			return true
		}
	}
	return false
}

func isValidPassword(password string) bool {
	hasNumber := false
	hasUpper := false
	hasLower := false
	hasSpecial := false

	for _, char := range password {
		if char >= '0' && char <= '9' {
			hasNumber = true
		}

		if char >= 'A' && char <= 'Z' {
			hasUpper = true
		}

		if char >= 'a' && char <= 'z' {
			hasLower = true
		}

		if isSpecialChars(char) {
			hasSpecial = true
		}
	}

	return len(password) >= 8 && hasNumber && hasUpper && hasLower && hasSpecial
}

func main() {
	var password string

	fmt.Print("Entrez votre mot de passe : ")
	fmt.Scanln(&password)

	if isValidPassword(password) {
		fmt.Println("Mot de passe valide ✅")
	} else {
		fmt.Println("Mot de passe invalide ❌")
	}

}
