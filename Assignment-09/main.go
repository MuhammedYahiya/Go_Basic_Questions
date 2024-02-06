package main

import "fmt"

func main() {
	var limit int
	fmt.Print("Enter the limit:")
	fmt.Scanf("%d", &limit)
	for i := 1; i <= limit; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
