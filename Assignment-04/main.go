package main

import "fmt"

func main() {
	var mark float64
	fmt.Print("Enter your mark: ")
	fmt.Scanf("%f", &mark)

	if mark > 100 {
		fmt.Println("Please Enter mark between 0 and 100")

	} else if mark >= 50 {
		fmt.Println("You Passed")
	} else {

		fmt.Println("Better luck next time")

	}
}
