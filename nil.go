package main
import "fmt"

// nil hanya bisa di gunakan di beberapa tipe data -> interface, function, map, slice, pointer dan channel

func NewMap(name string) map[string]string {
	if name == "" {
		return nil // return nil if name is empty
	}else {
		return map[string]string{
			"name": name,
			} // return a map with the name
	}
}

func main() {
	data := NewMap("kaes")
	if data == nil {
		fmt.Println("Data is nil") // Check if data is nil
	} else {
		fmt.Println(data["name"]) // Print the data if it's not nil
	}
}