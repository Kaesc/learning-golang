package main
import (
	"belajar-golang/database" // Corrected import path
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase()) // Output: MySQL
}