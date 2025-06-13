package main

import "fmt"

func main() {
	var number int

	fmt.Print("Combien de note souhaite tu rentrés : ")
	_, err := fmt.Scan(&number)

	if err != nil || number < 1 {
		fmt.Println("Erreur de saisie")
		return
	}

	var notes []float64

	for i := 1; i < number+1; i++ {
		var note float64

		for {
			fmt.Printf("Entrez la note %d : ", i)
			_, err := fmt.Scan(&note)

			if err == nil && note >= 0 && note <= 20 {
				break
			}

			fmt.Println("Erreur de saisie. Veuillez entrer une note entre 0 et 20")
		}

		notes = append(notes, note)
	}

	result := totalAverage(notes)

	fmt.Printf("Votre moyenne est de : %.2f sur 20\n", result)
}

func totalAverage(notes []float64) float64 {
	var average float64

	for _, note := range notes {
		average += float64(note)
	}

	return average / float64(len(notes))
}
