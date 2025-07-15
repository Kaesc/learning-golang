package main
import "fmt"

type Customer struct {
	Name, Address string
	Age 		  int
}

func main() {
	var customer Customer
	customer.Name = "Berdine Rose"
	customer.Address = ", Rizal Street, Manila"
	customer.Age = 70

	fmt.Println(customer)

	Kuma := Customer{
		Name:    "Bartholomew Kumazi",
		Address: "1234 Elm Street, Zaddyville",
		Age:      18,
	}

	fmt.Println(Kuma)

	LysBien := Customer{"Lys Bien", "Puta's village, Rizal", 17}
	fmt.Println(LysBien)
}

