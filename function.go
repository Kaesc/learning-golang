package main
import "fmt"

func sayHello() {
	fmt.Println("Hello, World!")
}

// func main() {
// 	sayHello() // Memanggil fungsi sayHello
// 	sayHello() 
// }

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello,", firstName, lastName)
}

func main() {
	sayHelloTo("John", "Doe") // Memanggil fungsi sayHelloTo dengan argumen
	sayHelloTo("Jane", "Smith")
}