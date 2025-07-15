package main
import "fmt"

// variable arguments (variadic function) allow a function to accept a variable number of arguments. varargs
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println(sumAll(1, 2, 3)) // Output: 6
	fmt.Println(sumAll(4, 5, 6, 7, 8)) // Output: 30

	numbers := []int{10, 20, 30, 40}
	fmt.Println(sumAll(numbers...)) // Output: 100
}