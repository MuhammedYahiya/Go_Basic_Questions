package main

import "fmt"

func main() {
	var limit int
	fmt.Print("Enter the limit of the array: ")
	fmt.Scanf("%d", &limit)
	arr := make([]int, limit)
	fmt.Print("Enter the elements of the array: ")
	for i := 0; i < limit; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	for i := 0; i < limit-1; i++ {
		arr[i] = arr[i] * arr[i+1]
		fmt.Print(arr[i], " ")
	}

}
