package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Product struct {
	Nom      string
	Prix     float64
	Quantite int
}

func affichage_menu() {
	fmt.Printf("1 - Ajouter un produit\n")
	fmt.Printf("2 - Supprimer un produit\n")
	fmt.Printf("3 - Modifier un produit\n")
	fmt.Printf("4 - Recherchez un produit\n")
	fmt.Printf("5 - Listez les produit \n")
	fmt.Printf("6 - Quittez \n")
}
func add_product() {
	fmt.Println("Indiquer le nom du produit")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Indiquer le prix du produit")
	scanner.Scan()
	price, _ := strconv.ParseFloat(scanner.Text(), 64)
	price = verif_price(price)

	fmt.Println("Indiquer la quantité du produit")
	scanner.Scan()
	quantity, _ := strconv.Atoi(scanner.Text())
	quantity = verif_quantity(quantity)

	product := Product{name, price, quantity}
	fmt.Println(product)
	fmt.Println("Produit ajouté")

}

func del_product() {
	fmt.Println("Menu Supprimer")
}
func update_product() {
	fmt.Println("Menu Modifié")
}

func list_product(create_list bool) {
	if create_list {
		fmt.Println("Menu List")
	} else {
		fmt.Println("Menu recherchez")
	}
}
func verif_price(price float64) float64 {
	scanner := bufio.NewScanner(os.Stdin)
	for price == 0 {
		fmt.Println("Le prix ne peut pas être 0, pour les prix à virgule utiliser le point")
		scanner.Scan()
		price, _ = strconv.ParseFloat(scanner.Text(), 64)
	}
	return price
}
func verif_quantity(quantity int) int {
	scanner := bufio.NewScanner(os.Stdin)
	for quantity == 0 {
		fmt.Println("La quantité ne peut pas être 0, ne peut pas être du texte")
		scanner.Scan()
		quantity, _ = strconv.Atoi(scanner.Text())
	}
	return quantity
}

func main() {

	var text int
	choix := 0
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Bonjour vous voici dans votre inventaire. Que voulez vous faire ?\n")
	//Initialiser une seule fois
	for choix > 6 || choix < 1 {
		affichage_menu()
		scanner.Scan()                         //Permet d'écouter l'entrée standard
		text, _ = strconv.Atoi(scanner.Text()) // Permet de récuperer le text entrée par l'utilisateur
		choix = text
		switch choix {
		case 1:
			add_product()
			choix = 0
		case 2:
			del_product()
			choix = 0
		case 3:
			update_product()
			choix = 0
		case 4:
			list_product(false)
			choix = 0
		case 5:
			list_product(true)
			choix = 0
		case 6:
			break
		}
	}

}
