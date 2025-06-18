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

	fmt.Print("Que souhaitez vous faire :\nCréer un contact (1) - Modifier un contact (2) - Supprimé un contact (3) - Ne rien faire (4) : ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	choice = strings.ToLower(choice)

	switch choice {
	case "1":
		controllers.CreateController()
	case "2":
		controllers.UpdateController()
	case "3":
		controllers.DeleteContact()
	default:
		fmt.Println("Aucun choix choisi")
	}
}

func loadedData(readContacts []controllers.Person) {
	for i, data := range readContacts {
		fmt.Printf("Contact %d: \nNom: %s \nAge: %d \nEmail: %s\n\n", i+1, data.Name, data.Age, data.Email)
	}
}
