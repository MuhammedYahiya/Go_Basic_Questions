package main

import "fmt"

func main() {
	var limit int
	fmt.Print("Enter a limit: ")
	fmt.Scanf("%d", &limit)
	sum := 0
	for i := 1; i <= limit; i++ {
		if i%2 != 0 {
			sum = sum + i
		}
	}

	fmt.Println("Sum of odd numbers:", sum)
}
