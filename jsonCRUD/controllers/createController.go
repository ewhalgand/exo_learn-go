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

func CreateController() {
	people, err := LoadContact()
	if err != nil {
		log.Println("Erreur lors de chargement des contacts")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nom du contact : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "" {
		log.Println("Erreur le nom est vide")
		return
	}

	fmt.Print("Age du contact : ")
	ageStr, _ := reader.ReadString('\n')
	ageStr = strings.TrimSpace(ageStr)
	age, err := strconv.Atoi(ageStr)
	if err != nil || age <= 0 || age > 130 {
		log.Println("Veuiller saisir un age valide")
		return
	}

	fmt.Print("Email du contact : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if email == "" {
		log.Println("Veuiller saisir un email valide")
		return
	}

	file, err := os.Create("contacts.json")
	if err != nil {
		log.Fatalln("Erreur lors de la creation du fichier")
	}
	defer file.Close()

	people = append(people, Person{
		Name:  name,
		Age:   age,
		Email: email,
	})

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	err = encoder.Encode(&people)
	if err != nil {
		log.Println("Erreur lors de l'encodage")
		return
	}

	fmt.Println("Ajout du contact réussi !")
}
