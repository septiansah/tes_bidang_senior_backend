package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	pass := ("TWaa@12")
	cekPassword := CekPassword(pass)
	if cekPassword == true {
		fmt.Println("Sukses membuat password :)")
	}
	// fmt.Println(cekPassword)
}

func CekPassword(pass string) bool {
	upper := 0
	lower := 0
	number := 0
	special := 0
	specialChar := []string{"@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "+"}
	// var result bool

	for _, ch := range pass {
		if unicode.IsUpper(ch) {
			upper++
		}
	}
	for _, ch := range pass {
		if unicode.IsLower(ch) {
			lower++
		}
	}
	for _, ch := range pass {
		if unicode.IsNumber(ch) {
			number++
		}
	}

	for _, sc := range specialChar {
		if strings.Contains(pass, sc) {
			special++
		}
	}

	if upper < 2 {
		fmt.Println("Memiliki minimal 2 huruf besar")
		return false
	} else if lower < 5 {
		fmt.Println("Memiliki minimal 5 huruf kecil ")
		return false
	} else if special == 0 {
		fmt.Println("Memiliki minimal special character (@#$%^&*()-+)")
		return false
	} else if number < 2 {
		fmt.Println("Memiliki minimal 2 angka")
		return false
	} else if len(pass) < 10 {
		fmt.Println("Salah karena kurang dari 10 karakter")
		return false
	} else if len(pass) > 20 {
		fmt.Println("Salah karena lebih dari 20 karakter")
		return false
	}

	return true
}
