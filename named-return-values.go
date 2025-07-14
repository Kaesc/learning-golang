package main
import "fmt"	

func names() (firstName, lastName string) {
	firstName = "Kaes"
	lastName = "Sukma"
	return // Menggunakan named return values
}

func main() {
	a, b := names() // Memanggil fungsi dengan named return values
	fmt.Println(a, b) // Output: Kaes Sukma
}