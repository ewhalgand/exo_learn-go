package main

import "fmt"

func calculTotal(price float64, quantity int) float64 {
	newPrice := price * float64(quantity)

	if quantity > 10 {
		return newPrice * 0.90
	}

	return newPrice
}

func main() {
	var productName string
	var price float64
	var quantity int

	fmt.Print("Entrer le nom du produit : ")
	fmt.Scanln(&productName)

	fmt.Print("Entrer le prix unitaire : ")
	fmt.Scanln(&price)

	fmt.Print("Entrer la quantité : ")
	fmt.Scanln(&quantity)

	totalResult := calculTotal(price, quantity)

	if quantity > 10 {
		fmt.Println("🎉 Remise de 10% appliquée !")
	}

	fmt.Printf("Le total pour %d %s est de %.2f €.\n", quantity, productName, totalResult)
}
