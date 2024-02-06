package main

import "fmt"

func main() {
	var totalMarkPercent float64

	fmt.Print("Enter the total mark percentage: ")
	fmt.Scanf("%f", &totalMarkPercent)

	if totalMarkPercent >= 90 {
		fmt.Printf("A grade")
	} else if totalMarkPercent >= 80 {
		fmt.Printf("B grade")
	} else if totalMarkPercent >= 70 {
		fmt.Println("C grade")
	} else if totalMarkPercent >= 60 {
		fmt.Println("D grade")
	} else if totalMarkPercent >= 50 {
		fmt.Printf("E grade")
	} else {
		fmt.Printf("Better luck next time")
	}
}
