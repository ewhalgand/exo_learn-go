package controllers

import (
	"encoding/json"
	"log"
	"os"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func LoadContact() ([]Person, error) {
	const fileName = "contacts.json"

	var people []Person

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatalln("Erreur lors de la création du fichier")
		}
		defer file.Close()

		file.Write([]byte("[]"))
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Erreur lors de la l'ouverture du fichier")
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&people)
	if err != nil {
		log.Fatalln("Erreur lors du décodage")
	}

	return people, err
}
