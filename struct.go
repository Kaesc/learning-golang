package main
import "fmt"

type Customer struct {
	Name, Address string
	Age 		  int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name) // Method function 
}

func main() {
	var customer Customer
	customer.Name = "Bernard Rose"
	customer.Address = ", Rizal Street, Manila"
	customer.Age = 70

	fmt.Println(customer)

	Kuma := Customer{
		Name:    "Bartholomew Kumazi",
		Address: "1234 Elm Street, Zaddyville",
		Age:      18,
	}

	fmt.Println(Kuma)

	LysBien := Customer{"LysBien", "Puta's village, Rizal", 17}
	fmt.Println(LysBien)

	Kuma.sayHello("LysBien") // Call method function of Kuma. Cant do this " sayHello() "
}

