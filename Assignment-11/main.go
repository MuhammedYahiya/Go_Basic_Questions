package main

import "fmt"

func main() {

	var arr [10]int
	var limit, i int
	count := 0
	fmt.Print("Enter the limit of the array: ")
	fmt.Scanf("%d", &limit)
	fmt.Print("Enter the elements of the array: ")
	for i = 0; i < limit; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	for i = 0; i < limit; i++ {
		if arr[i]%2 == 0 {
			count += 1
		}
	}

	fmt.Printf("Number of even numbers in the given array is: %d", count)

}
