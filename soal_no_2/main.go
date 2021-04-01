package main

import (
	"fmt"

	// "log"
	"sort"
)

func main() {
	angkaAcak := []int{1, 2, 3, 4, 5, 1, 4, 6, 8, 10}

	uniqueAngka := cekDuplikat(angkaAcak)

	hitungAngkaBerulang := hitungAngkaBerulang(angkaAcak)

	angka2 := []int{1, 2, 3, 4, 5, 1, 4, 6, 8, 10}
	sort.Sort(sort.Reverse(sort.IntSlice(angka2)))

	fmt.Println("(1) Hapus angka berulang di array tersebut sehingga menghasilkan ", uniqueAngka)
	fmt.Println("============================================")

	fmt.Println("(3) Jumlah angka berulang di array ", hitungAngkaBerulang)
	fmt.Println("============================================")

	fmt.Println("(4) Urutan bilangan secara descending ", angka2)

}

//Hapus angka berulang di array tersebut sehingga menghasilkan 1,2,3,4,5,6,8,1
func cekDuplikat(angkaAcak []int) []int {
	cek := make(map[int]bool)
	angkas := []int{}
	for _, angka := range angkaAcak {
		if _, value := cek[angka]; !value {

			cek[angka] = true

			angkas = append(angkas, angka)
		}
	}
	return angkas
}

// Jumlah angka berulang di array tersebut
func hitungAngkaBerulang(angkaAcak []int) int {
	cek := make(map[int]bool)
	angkas := []int{}
	for _, angka := range angkaAcak {
		if _, value := cek[angka]; !value {
			cek[angka] = true
		} else {
			angkas = append(angkas, angka)
		}
	}
	hasil := len(angkas)
	return hasil
}
