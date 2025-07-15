package main
import "fmt"
func main() {
	// counter := 1

	// for counter <= 5 {
	// 	fmt.Println("Counter is: ", counter)
	// 	counter++ // Increment counter by 1
	// }

	for counter := 1; counter <= 10; counter++ {
		if counter%2 == 0 {
			fmt.Println("Counter is even: ", counter)
		} else {
			fmt.Println("Counter is odd: ", counter)
		}
	}

	fmt.Println("Finish")

	// For Range Loop
	names := []string{"Alice", "Bob", "Charlie", "Diana"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Printf("Index: %d, Name: %s\n", index, name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
	
}