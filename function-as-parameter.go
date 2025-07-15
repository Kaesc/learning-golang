package main
import "fmt"

type Filter func(string) string

// Function as parameter example
func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello,", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" || name == "Babi" {
		return "***"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Kaes", spamFilter) // Output: Hello, Kaes
	sayHelloWithFilter("Anjing", spamFilter) // Output: Hello, ***
	
	filter := spamFilter
	sayHelloWithFilter("Babi", filter) // Output: Hello, ***

}