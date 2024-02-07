package main

import "fmt"

func main() {
	var annual_income float64
	fmt.Print("Enter the annual income: ")
	fmt.Scanf("%f", &annual_income)

	if annual_income <= 250000 {
		fmt.Println("No Tax")
	} else if annual_income <= 500000 {
		fmt.Printf("Income tax amount is: %.2f", (annual_income*5)/100)
	} else if annual_income <= 1000000 {
		fmt.Printf("Income tax amount is: %.2f", (annual_income*20)/100)
	} else if annual_income <= 5000000 {
		fmt.Printf("Income tax amount is: %.2f", (annual_income*30)/100)
	}
}
