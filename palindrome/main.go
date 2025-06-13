package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var text string
	var isPalindrom string

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Entrer un mot : ")
	text, err := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, " ", "")

	if err != nil {
		fmt.Println("Erreur : ce n'es pas un mot")
		return
	}

	reverse := reverseString(text)
	if reverse == text {
		isPalindrom = "est un Palindrome"
	} else {
		isPalindrom = "n'est pas un Palindrome"
	}

	fmt.Printf("Le mot %s %s.\n", text, isPalindrom)
}

func reverseString(text string) string {
	rns := []rune(text)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}
