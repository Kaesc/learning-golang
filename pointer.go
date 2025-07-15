package main
import "fmt"

/* 
	Passing by value example (so it is not a pass by reference)
	In Go, when you assign a struct to another variable, it creates a copy of the struct.
	This means that changes made to the new variable do not affect the original variable.
*/
type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// address2 := address1 // Copying the value

	// address2.City = "Bandung" // Changing the city in address2
	// fmt.Println(address1) // Tidak berubah valuenya
	// fmt.Println(address2) // Berubah valuenya

	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address2 := &address1 // Pointer

	address2.City = "Bandung" // Changing the city in address2
	fmt.Println(address1) // Ikut berubah valuenya
	fmt.Println(address2) 
}