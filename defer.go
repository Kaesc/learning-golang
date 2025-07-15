package main 
import "fmt"

func logging() {
	fmt.Println("Function executed")
}

func runApp() {
	defer logging() // This will execute after runApp finishes, even if it panics
	fmt.Println("Running application...")
}

func main() {
	runApp()
}