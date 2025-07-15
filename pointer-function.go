package main
import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIndo(address *Address) { // Function that takes a pointer to Address
	address.Country = "Indonesia" // This change will not affect the original address
}

func main() {
	address := Address{}
	changeCountryToIndo(&address) // Passing the address by reference using &
	fmt.Println(address) // The country will not be "Indonesia" because it was passed by value
}