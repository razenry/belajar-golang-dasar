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
	fmt.Println(sayHello(name, "Siang"))

	name = "Eko Ari Iswanto"
	age = 38
	isMarried = true

	fmt.Println("Nama : ", name)
	fmt.Println("Umur : ", age)
	fmt.Println("Jenis Kelamin : ", gender)
	fmt.Println("Menikah : ", isMarriedCheck(isMarried))
	fmt.Println(sayHello(name, "Siang"))

	name = "Ela Asmani"
	age = 48
	gender = "Perempuan"
	isMarried = true

	fmt.Println("Nama : ", name)
	fmt.Println("Umur : ", age)
	fmt.Println("Jenis Kelamin : ", gender)
	fmt.Println("Menikah : ", isMarriedCheck(isMarried))
	fmt.Println(sayHello(name, "Siang"))

	name = "Razenry Aretha Silverine"
	fmt.Println(sayHello(name, "Siang"))	
}

func isMarriedCheck(isMarried bool) string {
	if isMarried {
		return "Ya"
	}
	return "Tidak"
}

func sayHello(name string, time string) string {
	greating := "Selamat " + time	

	return "Hello " + name + " " + greating
}
