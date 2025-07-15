package main
import "fmt"

// interface{} or any
func Oops() any {
	return true
	// return 12 
	// return "This is an example of a function that returns an empty interface."
}

func main() {
	var kosong any = Oops()
	fmt.Println(kosong)
}