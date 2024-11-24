package main

import "fmt"

const NMAX int = 37

type tHimpunan struct {
	anggota [NMAX]int
	panjang int
}

func main() {
	var A1, A2 tHimpunan

	fmt.Print("Anggota himpunan 1: ")
	bacaMasukan(&A1)
	fmt.Print("Anggota himpunan 2: ")
	bacaMasukan(&A2)
	urut(&A1)
	urut(&A2)

	fmt.Printf("Himpunan 1 = Himpunan 2? %v", sama(A1, A2))
}

func bacaMasukan(A *tHimpunan) {
	var tmp, i int

	i = 0
	fmt.Scan(&tmp)
	for i < NMAX && !ada(*A, tmp) {
		A.anggota[i] = tmp
		i++
		A.panjang = i
		fmt.Scan(&tmp)
	}
}

func ada(A tHimpunan, x int) bool {
	var i int

	for i = 0; i < A.panjang; i++ {
		if A.anggota[i] == x {
			return true
		}
	}
	return false
}

func urut(A *tHimpunan) {
	var i, j, tmp int

	i = 1
	for i <= A.panjang-1 {
		j = i
		tmp = A.anggota[i]
		for j > 0 && tmp < A.anggota[j-1] {
			A.anggota[j] = A.anggota[j-1]
			j--
		}
		A.anggota[j] = tmp
		i++
	}
}

func sama(A1, A2 tHimpunan) bool {
	var i int

	if A1.panjang != A2.panjang {
		return false
	}

	for i = 0; i < A1.panjang; i++ {
		if A1.anggota[i] != A2.anggota[i] {
			return false
		}
	}
	return true
}
