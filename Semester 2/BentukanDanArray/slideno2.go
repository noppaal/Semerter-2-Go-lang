package main

import "fmt"

type geometry struct {
	area, perimeter int
}

type rectangle struct {
	p, l     int
	color    string
	property geometry
}

func isiData(persegi *rectangle) {
	fmt.Scanln(&persegi.p, &persegi.l, &persegi.color)
}

func hitung(persegi *rectangle) {
	isiData(persegi)
	persegi.property.area = persegi.p * persegi.l
	persegi.property.perimeter = persegi.p*2 + persegi.l*2
}

func main() {
	var persegi rectangle
	hitung(&persegi)
	fmt.Printf("Nilai luas dari persegi : %d\n", persegi.property.area)
	fmt.Printf("Nilai keliling dari persegi : %d\n", persegi.property.perimeter)
	fmt.Printf("Dengan warna dari persegi : %s\n", persegi.color)
}
