package main
import "fmt"

func random() any {
	return "ok"
}

// tidak bisa di ubah dua kali
func main() {
	var result any = random()
	// var resultString string = result.(string) // Type assertion to convert any to string
	// fmt.Println(resultString) // Output: ok

	// var resultInt int = result.(int) // This will panic if result is not an int
	// fmt.Println(resultInt) // This line will not be reached if the above line panics

	
	// Type assertion to check the type of result
	switch value := result.(type) {
	case string:
		fmt.Println("Result is a string:", value) // Output: Result is a string: ok
	case int:
		fmt.Println("Result is an int:", value) 
	default:
		fmt.Println("Result is of a different type:", value) 
	}

}