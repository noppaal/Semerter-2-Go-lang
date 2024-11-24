package main

import "fmt"

type pelari struct {
	asal  string
	waktu float32
}

func main() {
	var p1, p2, p3 pelari
	pemenang(&p1, &p2, &p3)
}

func pemenang(p1, p2, p3 *pelari) {
	fmt.Scanln(&p1.asal, &p1.waktu)
	fmt.Scanln(&p2.asal, &p2.waktu)
	fmt.Scanln(&p3.asal, &p3.waktu)
	if p1.waktu < p2.waktu && p1.waktu < p3.waktu {
		fmt.Printf("Atlet asal %s menang", p1.asal)
	} else if p2.waktu < p1.waktu && p2.waktu < p3.waktu {
		fmt.Printf("Atlet asal %s menang", p2.asal)
	} else if p3.waktu < p1.waktu && p3.waktu < p2.waktu {
		fmt.Printf("Atlet asal %s menang", p3.asal)
	} else {
		fmt.Print("seri")
	}
}
