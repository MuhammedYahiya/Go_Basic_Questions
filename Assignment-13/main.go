package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string
	var compare string
	fmt.Print("Enter a string: ")
	fmt.Scanf("%s", &name)
	loweStr := strings.ToLower(name)

	for i := len(loweStr) - 1; i >= 0; i-- {
		compare += string(loweStr[i])
	}

	if loweStr == compare {
		fmt.Println("It is a palindrome")
	} else {
		fmt.Println("It is not a palindrome")
	}

}
