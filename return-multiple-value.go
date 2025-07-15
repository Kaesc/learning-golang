package main 
import "fmt"

func getFullName() (string, string) {
	return "Kaes", "Sukma"
}

func main() {
	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName) // Output: Full Name: Kaes Sukma

	firstName, _ := getFullName()
	fmt.Println(firstName) // Output: Kaes 
}