package main
import "fmt"

func main() {
	name := "Kesya"

	switch name {
	case "Kesya":
		fmt.Println("Hello Kesya")
	case "Kaes":
		fmt.Println("Hello Kaes")
	default:
		fmt.Println("Hello Unknown")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Name is too long")
	case false:
		fmt.Println("Name is short enough")
	}

	//switch statement with multiple conditions
	
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Name is too long")
	case length > 5:
		fmt.Println("Name is short enough")
	default:
		fmt.Println("Name is just right")
	}
}