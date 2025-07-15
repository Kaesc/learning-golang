package main

import "fmt"
func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "Alice"
	// person["age"] = "30"
	
	person := map[string]string{
		"name": "Tila",
		"age":  "30",
	} 

	fmt.Println(person["name"]) // Output: Alice
	fmt.Println(person["age"])  // Output: 30
	fmt.Println(person) // Output: map[age:30 name:Alice]
	fmt.Println(person["address"]) // Output: (empty string, key not found)

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Kesya"
	book["ups"] = "Salah"

	fmt.Println(book) // Output: map[author:Kesya title:Belajar Golang ups:Salah]
	delete(book, "ups") // Menghapus key "ups"
	fmt.Println(book) // Output: map[author:Kesya title:Belajar Golang
}