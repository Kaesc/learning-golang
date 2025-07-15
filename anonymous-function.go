package main
import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked from registering:", name)
	} else {
		fmt.Println("You registered successfully:", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Doggy" 
	}

	registerUser("Alice", blacklist) // Output: User registered successfully: Alice

	registerUser("Doggy", func(name string) bool{
		return name == "Doggy" // Anonymous function to check if name is "Doggy"
	})
}