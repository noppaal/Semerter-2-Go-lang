package main

import "fmt"

func main() {
	type waktu struct {
		Jam   int
		Menit int
	}
	type jadwalKeretaApi struct {
		normorKA, kotaAsal, kotaTujuan string
		berangkat, tiba                waktu
	}

	var k1, k2, k3 jadwalKeretaApi

	fmt.Scanln(&k1.normorKA, &k1.kotaAsal, &k1.berangkat.Jam, &k1.berangkat.Menit, &k1.kotaTujuan, &k1.tiba.Jam, &k1.tiba.Menit)
	fmt.Scanln(&k2.normorKA, &k2.kotaAsal, &k2.berangkat.Jam, &k2.berangkat.Menit, &k2.kotaTujuan, &k2.tiba.Jam, &k2.tiba.Menit)
	fmt.Scanln(&k3.normorKA, &k3.kotaAsal, &k3.berangkat.Jam, &k3.berangkat.Menit, &k3.kotaTujuan, &k3.tiba.Jam, &k3.tiba.Menit)

	b1 := k1.berangkat.Jam*60 + k1.berangkat.Menit
	t1 := k1.tiba.Jam*60 + k1.tiba.Menit
	p1 := t1 - b1

	b2 := k2.berangkat.Jam*60 + k2.berangkat.Menit
	t2 := k2.tiba.Jam*60 + k2.tiba.Menit
	p2 := t2 - b2

	b3 := k3.berangkat.Jam*60 + k3.berangkat.Menit
	t3 := k3.tiba.Jam*60 + k3.tiba.Menit
	p3 := t3 - b3

	if p1 > p2 && p1 > p3 {
		fmt.Print(k1.normorKA)
	} else if p2 > p1 && p2 > p3 {
		fmt.Print(k2.normorKA)
	} else if p3 > p1 && p3 > p2 {
		fmt.Print(k3.normorKA)
	}
}
