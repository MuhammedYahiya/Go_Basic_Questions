package main

import "fmt"

func main() {
	k := 0
	for i := 1; i <= 4; i++ {
		for j := 1; j <= i; j++ {
			k += 1
			fmt.Print(k, " ")
		}
		fmt.Println()

	}
}
