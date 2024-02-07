package main

import "fmt"

func main() {
	var num1, num2, choice int
	fmt.Print("Enter two Number: ")
	fmt.Scanf("%d %d", &num1, &num2)
	fmt.Println("Which Operation Do you want to Perform.\n1)Addition\n2)Subtraction\n3)Multiplication\n4)Division")
	fmt.Print("Enter ur Choice: ")
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		sum := addition(num1, num2)
		fmt.Println("Sum is: ", sum)
	case 2:
		diff := subtraction(num1, num2)
		fmt.Println("Difference is: ", diff)
	case 3:
		mul := multiplication(num1, num2)
		fmt.Println("The product is: ", mul)
	case 4:
		quo := division(num1, num2)
		fmt.Println("The quotient is: ", quo)
	default:
		fmt.Println("Enter invalid option!")
	}

}

func addition(num1 int, num2 int) int {
	return num1 + num2
}

func subtraction(num1 int, num2 int) int {
	return num1 - num2
}

func multiplication(num1 int, num2 int) int {
	return num1 * num2
}

func division(num1 int, num2 int) int {
	return num1 / num2
}
