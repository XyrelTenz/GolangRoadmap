package Playground

import (
	"errors"
	"fmt"
	"log"
)

func validateStudents(students []int) error {
	// Check if students is Empty
	if len(students) == 0 {
		return errors.New("No Students ID Found")
	}
	// Check if Invalid Value
	for _, studentList := range students {
		if studentList <= 0 {
			return fmt.Errorf("Invalid Student ID: %d", studentList)
		}
	}
	return nil
}

func Validation() {

	studentsID := [...]int{
		1, 2, 3, 4, 5, 6, 7,
	}

	err := validateStudents(studentsID[:])
	if err != nil {
		log.Fatalf("Validation Error: %v", err)
	}
	fmt.Println("List of StudentID")

	for _, studentID := range studentsID {
		fmt.Println("Student ID:", studentID)
	}
}
