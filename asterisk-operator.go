package main
import "fmt"

/* 
	
*/
type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "DKI Jakarta", "Indonesia"}
	address2 := &address1 // Pointer

	address2.City = "Bandung" // Changing the city in address2
	fmt.Println(address1) // Ikut berubah valuenya
	fmt.Println(address2) 

	// address2 = &Address{"Jakarta", "Jawa Barat", "Indonesia"} // Reassigning address2 to a new Address
	*address2 = Address{"Jakarta", "Jawa Barat", "Indonesia"}
	fmt.Println(address1) // Tidak berubah valuenya
	fmt.Println(address2) // Berubah valuenya
}