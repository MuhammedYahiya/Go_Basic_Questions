package main

import "fmt"

func main() {
	var arr1 [3][3]int
	var arr2 [3][3]int
	var arr3 [3][3]int
	fmt.Print("Enter the value of array 1: ")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scanf("%d", &arr1[i][j])
		}
	}

	fmt.Print("Enter the value of array 2: ")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scanf("%d", &arr2[i][j])
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			arr3[i][j] = arr1[i][j] + arr2[i][j]
		}
	}

	fmt.Print("Sum of 2 array is: \n")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d%s", arr3[i][j], " ")
		}
		fmt.Println()
	}

}
