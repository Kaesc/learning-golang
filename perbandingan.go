package main
import "fmt"
func main() {
	var name1 = "Kaes"
	var name2 = "Kaes"
	var result bool 
if name1 == name2 {
	result = true
} else {
	result = false
}
	fmt.Println("Apakah nama1 sama dengan nama2?", result)
}