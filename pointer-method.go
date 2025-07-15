package main
import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() { //method with pointer receiver
	man.Name = "Mr " + man.Name
}

func main() {
	Eko := Man{"Eko"}
	Eko.Married() // This will not change Eko's Name because it is passed by value
	fmt.Println(Eko.Name) // Output: Eko
}