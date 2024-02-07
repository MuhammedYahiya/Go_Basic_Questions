package main

import "fmt"

func main() {
	var limit int
	fmt.Print("Enter the limit of the array: ")
	fmt.Scanf("%d", &limit)
	arr1 := make([][]int, limit)
	arr2 := make([][]int, limit)
	sum := make([][]int, limit)

	//initialize each inner slice of arr1 and arr2 with slices of integers of length limit
	for i := 0; i < limit; i++ {
		arr1[i] = make([]int, limit)
		arr2[i] = make([]int, limit)
		sum[i] = make([]int, limit)
	}
	getArray(arr1, arr2, limit)
	addArray(arr1, arr2, sum, limit)
	displayArray(sum, limit)
}

func getArray(arr1 [][]int, arr2 [][]int, limit int) {
	fmt.Println("Enter the elements of the first array:")
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Scanf("%d", &arr1[i][j])
		}
	}
	fmt.Println("Enter the elements of the second array:")
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Scanf("%d", &arr2[i][j])
		}
	}

}

func addArray(arr1 [][]int, arr2 [][]int, sum [][]int, limit int) {

	fmt.Println("Sum of array1 and array2:")
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			sum[i][j] = arr1[i][j] + arr2[i][j]
		}
	}
}

func displayArray(sum [][]int, limit int) {
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Print(sum[i][j], " ")
		}
		fmt.Println()
	}
}
