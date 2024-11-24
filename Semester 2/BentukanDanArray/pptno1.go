package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	p titik
	r int
}

func setTitik() titik {
	var p titik
	fmt.Scan(&p.x, &p.y)
	return p
}

func setLingkaran() lingkaran {
	var l lingkaran
	l.p = setTitik()
	fmt.Scan(&l.r)
	return l
}

func jarak2Titik(t1, t2 titik) float64 {
	return math.Sqrt((float64(((t1.x - t2.x) * (t1.x - t2.x)) + ((t1.y - t2.y) * (t1.y * t2.y)))))
}

func jarak2Lingkaran(l1, l2 lingkaran) {
	var jarak, jarakAntarapusat float64
	jarak = jarak2Titik(l1.p, l2.p)
	jarakAntarapusat = float64(l1.r + l2.r)
	if jarak < jarakAntarapusat {
		fmt.Println("Kedua lingkaran beririsan")
	} else if jarak == jarakAntarapusat {
		fmt.Println("kedua lingkaran bersinggungan")
	} else {
		fmt.Println("Kedua lingkarang berjauhan")
	}
}

func main() {
	var l1, l2 lingkaran
	l1 = setLingkaran()
	l2 = setLingkaran()

	fmt.Println("Lingkaran 1:", l1)
	fmt.Println("Lingkaran 2:", l2)
	jarak2Lingkaran(l1, l2)
}
