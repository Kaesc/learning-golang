package main
import "fmt"
func main() {
	name := "Kesyaa"

	if name == "Kesyaa" {
		fmt.Println("Hello Kesyaa")
	} else if name == "Kaes"{
		fmt.Println("Hello Kaes")
	} else {
		fmt.Println("Hello Unknown")
	}

	// IF short Statement
	if length := len(name); length > 5 {
		fmt.Println("Name is too long")
	}else{
		fmt.Println("Name is short enough")
	}
}