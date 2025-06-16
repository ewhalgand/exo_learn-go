package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	file, err := os.Open("./contacts.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouvertur du fichier")
		return
	}
	defer file.Close()

	var contacts []Person

	err = json.NewDecoder(file).Decode(&contacts)
	if err != nil {
		fmt.Println("Erreur de décodage")
		return
	}

	for _, data := range contacts {
		fmt.Printf("Mon nom est %s, j'ai %d ans et voici mon contact : %s\n", data.Name, data.Age, data.Email)
	}
}
