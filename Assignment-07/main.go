package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &num)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", i, num, i*num)
	}
}
