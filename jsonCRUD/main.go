package main

import (
	"bufio"
	"fmt"
	"jsonCRUD/controllers"
	"log"
	"os"
	"strings"
)

func main() {
	readContacts, err := controllers.LoadContact()
	if err != nil {
		log.Printf("Erreur de décodage")
		return
	}

	if len(readContacts) == 0 {
		fmt.Println("Aucun contact trouvé")
	}

	loadedData(readContacts)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Souaitez vous supprimez un contact ? (O/N) ")
	isDelete, _ := reader.ReadString('\n')
	isDelete = strings.TrimSpace(isDelete)
	isDelete = strings.ToLower(isDelete)

	if isDelete == "o" {
		controllers.DeleteContact()
	}
}

func loadedData(readContacts []controllers.Person) {
	for i, data := range readContacts {
		fmt.Printf("Contact %d: \nNom: %s \nAge: %d \nEmail: %s\n\n", i+1, data.Name, data.Age, data.Email)
	}
}
