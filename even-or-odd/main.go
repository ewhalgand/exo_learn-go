package main

import "fmt"

func main() {
	var number int
	var result string

	fmt.Print("Entrer un nombre : ")
	_, err := fmt.Scan(&number)

	if err != nil {
		fmt.Println("Erreur : Nombre Invalide")
		return
	}

	if (number % 2) == 0 {
		result = "Pair"
	} else {
		result = "Impair"
	}

	fmt.Printf("Le nombre %d est %s.\n", number, result)
}
