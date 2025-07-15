package main
import "fmt"

// Be careful with closures, they can capture variables from their surrounding scope and change them.
// This can lead to unexpected behavior if not handled properly. u might be confused about the output.

func main() {
	counter := 0




	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	increment()

	fmt.Println("Counter:", counter) // Output: Counter: 3
}

