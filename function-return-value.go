package main
import "fmt"

func getHello(name string) string {
	return "Hello, " + name
}

func main() {
	result := getHello("John")
	fmt.Println(result) // Output: Hello, John

	fmt.Println(getHello("Jane")) // Output: Hello, Jane
}