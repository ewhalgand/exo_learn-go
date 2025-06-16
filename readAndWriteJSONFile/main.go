package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	persons, _ := jsonDecoder()
	person, _ := jsonEncoder(persons)

	for _, data := range person {
		fmt.Printf("Nom du contact : %s\nAge du contact : %d\nEmail du contact : %s\n\n", data.Name, data.Age, data.Email)
	}
}

func jsonEncoder(persons []Person) ([]Person, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Veuillez entrer le nom du nouveau contact : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Veuillez entrer l'age du nouveau contact : ")
	ageStr, _ := reader.ReadString('\n')
	ageStr = strings.TrimSpace(ageStr)
	age, _ := strconv.Atoi(ageStr)

	fmt.Print("Veuillez entrer l'email du nouveau contact : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	file, err := os.Create("contacts.json")
	if err != nil {
		fmt.Println("Erreur lors de la creation du fichier")
		return nil, err
	}
	defer file.Close()

	persons = append(persons, Person{
		Name:  name,
		Age:   age,
		Email: email,
	})

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	err = encoder.Encode(&persons)
	if err != nil {
		fmt.Println("Erreur lors de l'encodage du fichier")
		return nil, err
	}

	return persons, nil
}

func jsonDecoder() ([]Person, error) {
	file, err := os.Open("contacts.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture", err)
		return nil, err
	}
	defer file.Close()

	var persons []Person

	err = json.NewDecoder(file).Decode(&persons)
	if err != nil {
		fmt.Println("Erreur lors du décodage", err)
		return nil, err
	}

	return persons, nil
}
