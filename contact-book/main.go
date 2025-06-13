package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var numbersOfContacts int

	fmt.Print("Entrez le nombre de contact à ajouter : ")
	_, err := fmt.Scan(&numbersOfContacts)

	if err != nil || numbersOfContacts <= 0 {
		fmt.Println("Erreur de saisie : Ce n'es pas un nombre ou le nombre est inférieur à 1 !")
		return
	}

	contacts := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)

	for i := 1; i <= numbersOfContacts; i++ {
		var name string
		var phoneNumber string

		fmt.Printf("Entrer le nom du contact %d : ", i)
		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(name)

		fmt.Printf("Entrer le numéro de téléphone du contact %d : ", i)
		phoneNumber, _ = reader.ReadString('\n')
		phoneNumber = strings.TrimSpace(phoneNumber)

		contacts[name] = strings.ReplaceAll(phoneNumber, " ", "-")
	}

	for name, phoneNumber := range contacts {
		fmt.Printf("\nContact : %s - %s\n", name, phoneNumber)
	}

}
