package main

import (
	"fmt"
)

func hitungLuasKelilingLingkaran(r float64, l, k *float64) {
	const pi = 3.14
	*l = pi * r * r
	*k = 2 * pi * r
}

func hitungLuasKelilingPersegi(s float64, l, k *float64) {
	*l = s * s
	*k = 4 * s
}

func hitungTotal(lL, lP, kL, kP float64, totLuas, totKel *float64) {
	*totLuas = lL + lP
	*totKel = kL + kP
}

func main() {
	var radius, sisi, lL, lP, kL, kP, totLuas, totKel float64

	fmt.Scanln(&radius, &sisi)
	fmt.Printf("%7s %7s %7s %7s %7s %7s %7s %7s\n", "R", "S", "LL", "LP", "KL", "KP", "TL", "TP")
	for radius != 0 || sisi != 0 {
		hitungLuasKelilingLingkaran(radius, &lL, &kL)
		hitungLuasKelilingPersegi(sisi, &lP, &kP)
		hitungTotal(lL, lP, kL, kP, &totLuas, &totKel)
		fmt.Printf("%7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f\n", radius, sisi, lL, lP, kL, kP, totLuas, totKel)
		fmt.Scanln(&radius, &sisi)
	}
}
