package main

import (
	"GestionDeStock/controllers"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func affichage_menu() {
	fmt.Printf("1 - Ajouter un produit\n")
	fmt.Printf("2 - Supprimer un produit\n")
	fmt.Printf("3 - Modifier un produit\n")
	fmt.Printf("4 - Recherchez un produit\n")
	fmt.Printf("5 - Listez les produit \n")
	fmt.Printf("6 - Quittez \n")
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
			fmt.Print("\033[H\033[2J")
			controllers.AddProduct()
			choix = 0
		case 2:
			fmt.Print("\033[H\033[2J")
			fmt.Printf("Quelle est le nom du produit à supprimer? \n")
			scanner.Scan()
			rep := scanner.Text()
			controllers.DelProduct(rep)
			choix = 0
		case 3:
			fmt.Print("\033[H\033[2J")
			fmt.Printf("Quelle est le nom du produit à modifier? \n")
			scanner.Scan()
			rep := scanner.Text()
			fmt.Printf("Que souhaitez vous modifier? \n")
			fmt.Printf("1 - Nom\n")
			fmt.Printf("2 - Prix\n")
			fmt.Printf("3 - Quantité\n")
			change := 0
			for change > 3 || change < 1 {
				scanner.Scan()
				change, _ = strconv.Atoi(scanner.Text())
			}
			controllers.UpdateProduct(rep, change)
			choix = 0
		case 4:
			fmt.Print("\033[H\033[2J")
			fmt.Printf("Quelle est le nom du produit recherchez? \n")
			scanner.Scan()
			rep := scanner.Text()
			controllers.FindProduct(rep)
			choix = 0
		case 5:
			fmt.Print("\033[H\033[2J")
			controllers.ListProduct()
			choix = 0
		case 6:
			break
		}
	}

}
