package main

import "fmt"

func main() {
	var str1, str2 string
	fmt.Print("Enter your fist string: ")
	fmt.Scan(&str1)
	fmt.Print("Enter your second string: ")
	fmt.Scan(&str2)
	compare := compareString(str1, str2)
	if compare == -1 {
		fmt.Printf("%s is less than %s", str1, str2)
	} else if compare == 1 {
		fmt.Printf("%s is greater than %s", str1, str2)
	} else {
		fmt.Printf("%s is equal to %s", str1, str2)
	}
}

func compareString(str1 string, str2 string) int {
	if str1 < str2 {
		return -1
	} else if str1 > str2 {
		return 1
	}
	return 0
}
