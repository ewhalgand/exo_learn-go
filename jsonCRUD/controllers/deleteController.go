package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func DeleteContact() {
	people, err := LoadContact()
	if err != nil {
		log.Println("Erreur lors de chargement des contacts")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Quelle contact souhaite tu supprimés ? (par nom) ")
	nameToDelete, _ := reader.ReadString('\n')
	nameToDelete = strings.TrimSpace(nameToDelete)
	nameToDelete = strings.ToLower(nameToDelete)

	var deletePeople []Person
	found := false

	for _, person := range people {
		if strings.ToLower(person.Name) != nameToDelete {
			deletePeople = append(deletePeople, person)
		} else {
			found = true
		}
	}

	if !found {
		fmt.Println("Aucun contact avec ce nom n'a été trouvé")
		return
	}

	file, err := os.Create("contacts.json")
	if err != nil {
		log.Println("Erreur lors de la réouverture du fichier")
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	if err := encoder.Encode(deletePeople); err != nil {
		log.Println("Erreur lors de l'encodage")
		return
	}

	fmt.Println("Suppression réussi !")
}
