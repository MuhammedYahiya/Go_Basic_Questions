package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &number)
	flag := 0

	if number == 1 {
		flag = 1
	}

	for i := 2; i < number; i++ {
		if number%2 == 0 {
			flag = 1
			break
		}
	}

	if flag == 0 {
		fmt.Println("Its a prime")
	} else {
		fmt.Println("Its not a prime")
	}

}
