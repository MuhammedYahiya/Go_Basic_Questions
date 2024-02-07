package main

import "fmt"

func main() {
	var written_text, lab_exams, assignments float64
	fmt.Print("Enter the written test mark: ")
	fmt.Scanf("%f", &written_text)
	fmt.Print("Enter the lab exams mark: ")
	fmt.Scanf("%f", &lab_exams)
	fmt.Print("Enter the assignment mark: ")
	fmt.Scanf("%f", &assignments)
	overall_grade := (written_text*70)/100 + (lab_exams*20)/100 + (assignments*10)/100
	fmt.Println("Overall grade is: ", overall_grade)
}
