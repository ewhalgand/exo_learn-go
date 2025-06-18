package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func UpdateController() {
	people, err := LoadContact()
	if err != nil {
		log.Println("Erreur lors de chargement des contacts")
		return
	}

	index, _ := loadedData(people)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nouveau nom du contact (laisser vide pour ne pas changer) : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name != "" {
		people[index].Name = name
	}

	fmt.Println("Nouvelle age du contact (laisser vide pour ne pas changer) : ")
	ageStr, _ := reader.ReadString('\n')
	ageStr = strings.TrimSpace(ageStr)
	if ageStr != "" {
		age, err := strconv.Atoi(ageStr)
		if err == nil {
			people[index].Age = age
		}
	}

	fmt.Print("Nouvelle email du contact (laisser vide pour ne pas changer) : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if email != "" {
		people[index].Email = email
	}

	file, err := os.Create("contacts.json")
	if err != nil {
		log.Fatalln("Erreur lors de la creation du fichier")
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	err = encoder.Encode(&people)
	if err != nil {
		log.Println("Erreur lors de l'encodage")
		return
	}

	fmt.Println("Modification réussi !")
}

func loadedData(readContacts []Person) (int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Entrez l'id du contact à modifier : ")
	updatedContactStr, _ := reader.ReadString('\n')
	updatedContactStr = strings.TrimSpace(updatedContactStr)
	updatedContact, err := strconv.Atoi(updatedContactStr)

	if err != nil || updatedContact < 1 || updatedContact > len(readContacts) {
		log.Println("Veuillez saisir un id valide")
		return -1, err
	}

	return updatedContact - 1, nil
}
