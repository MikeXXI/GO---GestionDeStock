package controllers

import (
	"GestionDeStock/models"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var list_product []models.Product

func remove(tab []models.Product, s int) []models.Product {
	return append(tab[:s], tab[s+1:]...)
}

func AddProduct() {

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

	product := models.Product{
		Name:     name,
		Price:    price,
		Quantity: quantity}
	fmt.Println(product)
	list_product = append(list_product, product)
	fmt.Println("Produit ajouté")
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

func FindProduct(name string) {
	for i, product := range list_product {
		if product.Name == name {
			fmt.Println(i, product)
		}
	}
}

func DelProduct(name string) {
	for i, product := range list_product {
		if product.Name == name {
			list_product = remove(list_product, i)
			fmt.Println("Produit supprimé")
		}
	}
}
func UpdateProduct(rep string, change int) {
	scanner := bufio.NewScanner(os.Stdin)
	for i, product := range list_product {
		if product.Name == rep {
			if change == 1 {
				fmt.Printf("Quel est le nouveau nom du produit? \n")
				scanner.Scan()
				rep = scanner.Text()
				list_product[i].Name = rep
				fmt.Println("Produit modifié")
			}
			if change == 2 {
				fmt.Printf("Quel est le nouveau prix du produit? \n")
				scanner.Scan()
				rep, _ := strconv.ParseFloat(scanner.Text(), 64)
				rep = verif_price(rep)
				list_product[i].Price = rep
				fmt.Println("Produit modifié")
			}
			if change == 3 {
				fmt.Printf("Quel est la nouvelle quantité du produit? \n")
				scanner.Scan()
				rep, _ := strconv.Atoi(scanner.Text())
				rep = verif_quantity(rep)
				list_product[i].Quantity = rep
				fmt.Println("Produit modifié")
			}
		}
	}

}

func ListProduct() {
	for i, product := range list_product {
		fmt.Println(i, product)
	}
}
