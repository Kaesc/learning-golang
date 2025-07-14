package main
import "fmt"

func endApp() {
	fmt.Println("End of application")
	message := recover() // Recover from panic
	fmt.Println("Recovered from panic:", message)
}

func runApp(error bool) {
	defer endApp() // This will execute at the end of runApp, even if it panics
	if error {
		panic("Application encountered an error")
	}
}

func main() {
	runApp(true) // Change to true to see panic handling
	fmt.Println("Application is running smoothly")
}