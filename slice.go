package main 

import "fmt"

func main() {
	// Contoh penggunaan slice di Go
	// Slice adalah potongan dari array yang dapat berubah ukurannya
	names:= []string{"Alice", "Bob", "Charlie", "Diana", "Ethan", "Frank"}
	
	slice1 := names[4:6] // Mengambil elemen dari indeks 4 sampai 6
	fmt.Println(slice1) // Output: [Ethan Frank]

	slice2 := names[:3] // Mengambil elemen dari awal sampai indeks 3
	fmt.Println(slice2) // Output: [Alice Bob Charlie]

	slice3 := names[3:] // Mengambil elemen dari indeks 3 sampai akhir
	fmt.Println(slice3) // Output: [Diana Ethan Frank]

	var slice4 []string = names[:] // Mengambil semua elemen
	fmt.Println(slice4)

	days := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:]
	fmt.Println(daySlice1) // Output: [Sabtu Minggu]

	daySlice1[0] = "Sabtu Baru" // Mengubah elemen pertama dari daySlice1
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1) // Output: [Sabtu Baru Minggu Baru]
	
	// Append Slice
	daySlice2 := append(daySlice1, "Hari Baru")
	daySlice2[0] = "Sabtu New"
	fmt.Println(daySlice2) // Output: [Sabtu New Minggu Baru Hari Baru]
	fmt.Println(days) // Output: [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]

	// Make New SLice
	var newSlice []string = make([]string, 2, 5) // (2 -> length, 5 -> capacity)
	newSlice[0] = "Kaes"
	newSlice[1] = "Kaes1"
	// newSlice[2] = "Kaes2" // Ini akan menyebabkan panic karena index 2 tidak ada harus append
	fmt.Println(newSlice) 
	fmt.Println(len(newSlice)) 
	fmt.Println(cap(newSlice)) 

	newSlice2 := append(newSlice, "Kaes2")
	fmt.Println(newSlice2) // Output: [Kaes Kaes1 Kaes2]
	fmt.Println(len(newSlice2)) // Output: 3 (panjang baru)
	fmt.Println(cap(newSlice2)) // Output: 5 (kapasitas tetap 5

	newSlice2[0] = "Kaes New"
	fmt.Println(newSlice2) // Output: [Kaes New Kaes1 Kaes2]
	fmt.Println(newSlice)

	// Copy Slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice) 
	fmt.Println(toSlice) 

	// Slice vs Array (The Difference)
	iniArray := [...]string{"A", "B", "C"}
	iniSlice := []string{"A", "B", "C"}
	fmt.Println(iniArray) // Output: [A B C]
	fmt.Println(iniSlice) // Output: [A B C]
}