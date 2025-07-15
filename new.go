package main
import "fmt"

type Address struct {
	City, Province, Country string
}

func main(){
	// without using &
	var alamat1 *Address = new(Address)
	var alamat2 *Address = alamat1

	alamat2.City = "Bandung" 

	fmt.Println(alamat1) // Ikut berubah valuenya
	fmt.Println(alamat2) // Ikut berubah valuenya
}