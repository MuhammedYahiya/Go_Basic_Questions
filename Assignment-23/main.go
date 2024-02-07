package main

import "fmt"

func main() {
	var limit int
	fmt.Print("Enter the limit of the array: ")
	fmt.Scanf("%d", &limit)
	arr := make([][]int, limit)
	for i := 0; i < limit; i++ {
		arr[i] = make([]int, limit)
	}
	getArray(arr, limit)
	displayArray(arr, limit)
}

func getArray(arr [][]int, limit int) {
	fmt.Println("Enter the elements of the array:")
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Scanf("%d", &arr[i][j])
		}
	}
}
func displayArray(arr [][]int, limit int) {
	fmt.Println("The elements of arrays are:")
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}
