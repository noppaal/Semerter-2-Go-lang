package main

import "fmt"

func reamur(c float64) float64 {
	var result float64
	result = ((4 * c) / 5)
	return result
}

func fahrenheit(c float64) float64 {
	var result float64
	result = ((9 * c) / 5) + 32
	return result
}

func kelvin(c float64) float64 {
	var result float64
	result = c + 273
	return result
}

func main() {
	var awal, akhir, step float64
	var rea, fah, kel, i float64

	fmt.Scanln(&awal, &akhir, &step)
	fmt.Println("Celcius", "Reamur", "Fahrenheit", "Kelvin")
	for i = awal; i <= akhir; i += step {
		rea = reamur(i)
		fah = fahrenheit(i)
		kel = kelvin(i)
		fmt.Printf("%10.2f ", i)
		fmt.Printf("%10.2f ", rea)
		fmt.Printf("%10.2f ", fah)
		fmt.Printf("%10.2f\n ", kel)
	}
}
