package main
import "fmt"

func getGoodbye(name string) string {
	return "Goodbye, " + name
}

func main() {
	contoh := getGoodbye
	contoh2 := getGoodbye

	fmt.Println(contoh("Kaesy")) // Output: Goodbye, Kaes
	fmt.Println(contoh2("Kaes")) // Output: Goodbye, Kaes
}