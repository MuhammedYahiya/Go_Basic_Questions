package main

import "fmt"

func main() {
	var principalAmount int
	var interestRate, NumOfYears, SimpleInterest float64

	fmt.Print("Enter the principal amount: ")
	fmt.Scanf("%d", &principalAmount)
	fmt.Print("Enter the Interest rate: ")
	fmt.Scanf("%f", &interestRate)
	fmt.Print("Enter the Number of years: ")
	fmt.Scanf("%f", &NumOfYears)

	SimpleInterest = (float64(principalAmount) * interestRate * NumOfYears) / 100
	fmt.Printf("Simple Interest is %f", SimpleInterest)

}
