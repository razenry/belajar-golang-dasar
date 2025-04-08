package main

import (
	"fmt"
	"time"
)

// isMarriedCheck checks if a person is married and returns "Ya" or "Tidak".
func isMarriedCheck(isMarried bool) string {
	if isMarried {
		return "Ya"
	}
	return "Tidak"
}

// sayHello generates a greeting message based on the name and time of day.
func sayHello(name string, time string) string {
	greeting := "Selamat " + time
	return "Hello " + name + " " + greeting
}

// valueCheck evaluates a score and returns a corresponding description.
func valueCheck(value int) string {
	switch {
	case value > 95:
		return "Sempurna"
	case value > 80:
		return "Sangat baik"
	case value > 70:
		return "Lulus"
	case value > 60:
		return "Cukup"
	case value > 50:
		return "Kurang"
	case value > 40:
		return "Sangat kurang"
	case value > 30:
		return "Tidak lulus"
	case value > 20:
		return "Sangat tidak lulus"
	case value > 10:
		return "Gagal"
	case value > 0:
		return "Sangat gagal"
	default:
		return "Nol"
	}
}

func main() {
	var name string
	var age int
	var isMarried bool
	var genderInput string
	var gender string
	var value int

	// Input data from user
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&name)

	fmt.Print("Masukkan Umur: ")
	fmt.Scanln(&age)

	fmt.Print("Masukkan Jenis Kelamin (L/P): ")
	fmt.Scanln(&genderInput)
	if genderInput == "L" || genderInput == "l" {
		gender = "Laki-laki"
	} else if genderInput == "P" || genderInput == "p" {
		gender = "Perempuan"
	} else {
		gender = "Tidak diketahui"
	}

	fmt.Print("Apakah sudah menikah? (Y/n, default Y): ")
	var isMarriedInput string
	fmt.Scanln(&isMarriedInput)
	if isMarriedInput == "n" || isMarriedInput == "N" {
		isMarried = false
	} else {
		isMarried = true
	}

	fmt.Print("Masukkan Nilai: ")
	fmt.Scanln(&value)

	// Display data
	fmt.Println("\nData yang dimasukkan:")
	fmt.Println("Nama :", name)
	fmt.Println("Umur :", age)
	fmt.Println("Jenis Kelamin :", gender)
	fmt.Println("Menikah :", isMarriedCheck(isMarried))

	// Determine time of day
	var timeOfDay string
	hour := time.Now().Hour()
	switch {
	case hour >= 5 && hour < 12:
		timeOfDay = "Pagi"
	case hour >= 12 && hour < 18:
		timeOfDay = "Siang"
	case hour >= 18 && hour < 21:
		timeOfDay = "Sore"
	default:
		timeOfDay = "Malam"
	}

	// Display greeting and value evaluation
	fmt.Println(sayHello(name, timeOfDay))
	fmt.Println("Nilai " + valueCheck(value))
}
