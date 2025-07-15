package main
import (
	_"belajar-golang/internal" // Corrected import path
	"belajar-golang/database" // Corrected import path
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase()) // Output: MySQL
}