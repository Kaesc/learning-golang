package main

import (
	"fmt"
)

type ValidationError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}

type NotFoundError struct {
	Message string
}

func (n *NotFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any)error {
	if id == "" {
		return &ValidationError{"ID cannot be empty"}
	}

	if id != "Kaes" {
		return &NotFoundError{"Data not found"}
	}
	return nil
}

func main() {
	err := SaveData("Kaes", nil)
	if err != nil {
		// error

		// Menggunakan if else 
		// if ValidationErr, ok := err.(*ValidationError); ok {
		// 	fmt.Println("Validation Error: ", ValidationErr.Error())
		// } else if NotFoundErr, ok := err.(*NotFoundError); ok {
		// 	fmt.Println("Not Found Error: ", NotFoundErr.Error())
		// }else {
		// 	fmt.Println("Unknown Error: ", err.Error())
		// }

		// Menggunakan switch case
		switch finalError := err.(type) {
			case *ValidationError:
				fmt.Println("Validation Error: ", finalError.Error())
			case *NotFoundError:
				fmt.Println("Not Found Error: ", finalError.Error())
			default:
				fmt.Println("Unknown Error: ", finalError.Error())
		}
	} else {
		// success
		fmt.Println("Success")
	}
}