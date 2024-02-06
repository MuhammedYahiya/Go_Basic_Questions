package main

import "fmt"

func main() {
	var array1 [5]int
	var array2 [5]int
	var temp [5]int

	fmt.Print("Enter the elements of array one: ")
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &array1[i])
	}
	for i := 0; i < 5; i++ {
		fmt.Printf("%d%s", array1[i], " ")
	}

	fmt.Print("\nEnter the elements of array one: ")
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &array2[i])
	}
	for i := 0; i < 5; i++ {
		fmt.Printf("%d%s", array2[i], " ")
	}
	fmt.Print("\n")
	fmt.Println("Arrays after swapping")
	for i := 0; i < 5; i++ {
		temp[i] = array1[i]
		array1[i] = array2[i]
		array2[i] = temp[i]
	}
	fmt.Print("Array1: ")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d%s", array1[i], " ")
	}
	fmt.Print("\n")
	fmt.Print("Array2: ")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d%s", array2[i], " ")
	}

}
