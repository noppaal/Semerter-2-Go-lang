package main

import "fmt"

func hitungMenang(g, k int, jm *int) {
	if g > k {
		*jm++
	}
}

func hitungDraw(g, k int, jd *int) {
	if g == k {
		*jd++
	}
}

func hitungKalah(g, k int, jkal *int) {
	if g < k {
		*jkal++
	}
}

func hitungJumlahGolKegolanSelisih(g, k int, jg, jk, jsg *int) {
	*jg += g
	*jk += k
	*jsg = *jg - *jk
}

func hitungJumlahPoint(jm, jd int, jp *int) {
	*jp = (jm * 3) + jd
}

func main() {
	var n, jm, jd, jkal, jg, jk, jsg, jp int
	var g, k int

	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		fmt.Scanln(&g, &k)

		hitungMenang(g, k, &jm)
		hitungDraw(g, k, &jd)
		hitungKalah(g, k, &jkal)
		hitungJumlahGolKegolanSelisih(g, k, &jg, &jk, &jsg)
	}
	hitungJumlahPoint(jm, jd, &jp)
	fmt.Printf("%d %d %d %d %d %d %d %d\n", n, jm, jd, jkal, jg, jk, jsg, jp)
}
