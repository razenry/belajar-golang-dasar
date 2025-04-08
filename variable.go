package main

import "fmt"

func main() {
	
	var name string = "Kafkariela Prima Rasendriya"
	var age int = 20
	var isMarried bool = false
	var gender string = "Laki-laki"

	fmt.Println("Nama : ", name)
	fmt.Println("Umur : ", age)
	fmt.Println("Jenis Kelamin : ", gender)
	fmt.Println("Menikah : ", isMarriedCheck(isMarried))
	
	name = "Eko Ari Iswanto"
	age = 38
	isMarried = true
	
	fmt.Println("Nama : ", name)
	fmt.Println("Umur : ", age)
	fmt.Println("Jenis Kelamin : ", gender)
	fmt.Println("Menikah : ", isMarriedCheck(isMarried))
	
	name = "Ela Asmani"
	age = 48
	gender = "Perempuan"
	isMarried = true
	
	fmt.Println("Nama : ", name)
	fmt.Println("Umur : ", age)
	fmt.Println("Jenis Kelamin : ", gender)
	fmt.Println("Menikah : ", isMarriedCheck(isMarried))
}

func isMarriedCheck(isMarried bool) string {
	if isMarried {
		return "Ya"
	}
	return "Tidak"
}