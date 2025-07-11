package main
import "fmt"
func main() {
	var names = [5]string{"Alice", "Bob", "Charlie", "Diana", "Ethan"}
	var ages = [6]int{
		25, 
		30, 
		22, 
		28, 
		35,
		31, //tanda koma wajib di akhir
	}

	fmt.Println("Names:", names[1])
	fmt.Println("Ages:", ages[2])

	var values = [3]int{
		1, 
		2, 
		3,
	}
	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 10
	fmt.Println(values)
}