package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	jsonData := []byte(`[
		{
			"name": "Alice",
			"age": 30,
			"email": "alice@example.com"
		},
		{
			"name": "Bob",
			"age": 18,
			"email": "bob@example.com"
		}
	]`)

	var persons []Person

	err := json.Unmarshal(jsonData, &persons)
	if err != nil {
		fmt.Println("Erreur de parsing:", err)
		return
	}

	for _, person := range persons {
		fmt.Printf("Nom: %s, Age: %d, Email: %s\n", person.Email, person.Age, person.Email)
	}
}
