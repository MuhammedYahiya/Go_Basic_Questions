package main

import "fmt"

func main() {
	var arr [10]int
	var limit int

	fmt.Print("Enter the size of the array: ")
	fmt.Scanf("%d", &limit)
	fmt.Println("Enter the elements of the array: ")
	for i := 0; i < limit; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("The elements of the array: ")
	for i := 0; i < limit; i++ {
		fmt.Printf("%d%s", arr[i], " ")
	}
	fmt.Print("\n")
	fmt.Println("Sorted array: ")
	for i := 0; i < limit; i++ {
		for j := i + 1; j < limit; j++ {
			if arr[i] < arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	for i := 0; i < limit; i++ {
		fmt.Printf("%d%s", arr[i], " ")
	}

}
