package main

import "fmt"

func main() {
	var limit int

	fmt.Print("Enter the limit: ")
	fmt.Scanf("%d", &limit)
	arr := make([]int, limit) // Initialize the slice with the specified length
	getArray(arr, limit)
	displayArray(arr, limit)
}

func getArray(arr []int, limit int) {
	fmt.Print("Enter the elements of the array: ")
	for i := 0; i < limit; i++ {
		fmt.Scanf("%d", &arr[i])
	}
}

func displayArray(arr []int, limit int) {
	fmt.Printf("Array values are:")
	for i := 0; i < limit; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
