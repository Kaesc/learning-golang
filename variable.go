package main
import ("fmt")

func main() {
	var Name string

	Name = "John Doe"
	fmt.Println("Name:", Name)

	var age = 17
	fmt.Println("Age:", age)

	age = 20
	fmt.Println("Updated Age:", age)

	height := 1.75
	fmt.Println("Height:", height)

	var(
		firstName = "Cesya"
		lastName = "Apridita"
	)

	fmt.Println(firstName, lastName)
}