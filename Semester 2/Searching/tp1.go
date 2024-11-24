package main

import "fmt"

const NMAX int = 10

type tabInt [NMAX]int

func main() {
	var data tabInt
	var nData int
	var x2 int
	var idx int

	fmt.Scan(&x2)

	bacaData(&data, &nData)

	cetakData(data, nData)

	fmt.Print("Hasil pencarian: ")
	if binarySearch(data, nData, x2, &idx) {
		fmt.Printf("Bilangan %d ditemukan pada posisi ke-%d\n", x2, idx)
	} else {
		fmt.Println("Bilangan tidak ditemukan.")
	}
}

func bacaData(A *tabInt, n *int) {
	var input tabInt
	var i int = 0

	fmt.Scan(&input[i])
	for (i < NMAX) && (A[i] < input[i]) {
		A[i] = input[i]
		fmt.Scan(&input[i])
		i++
	}
	*n = i + 1
}

func cetakData(A tabInt, n int) {
	fmt.Print("Data Bilangan: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", A[i])
	}
	fmt.Println()
}

func frekuensiBilangan(A tabInt, n, x int) int {
	tot := 0
	for i := 0; i < n; i++ {
		if A[i] == x {
			tot += 1
		}
	}
	return tot
}

func sequentialSearch(A tabInt, n, x int) bool {
	var res bool
	res = false
	for i := 0; i < n; i++ {
		if A[i] == x {
			res = true
		}
	}
	return res
}
