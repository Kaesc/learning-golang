package main 
import "fmt"
func main() {
	type NoKTP string

	var ktpKaes NoKTP = "1234567890"
	var contoh string = "1234567222"
	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println(ktpKaes)
	fmt.Println(contohKTP)
}