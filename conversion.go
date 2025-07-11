package main
import("fmt")

func main() {
	// int 8 -> -128 to 127
	// int 16 -> -32768 to 32767
	// int 32 -> -2147483648 to 2147483647

	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Budi"
	var B uint8 = name[0] 
	var BString = string(B)
	fmt.Println(name)
	fmt.Println(B)
	fmt.Println(BString)
}