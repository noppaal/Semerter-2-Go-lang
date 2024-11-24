package main

import "fmt"

const NMAX int = 2022

type student struct {
	name, sid string
	gpa       float64
}
type tabMhs [NMAX]student

func main() {
	var A tabMhs
	var nData int

	isiArray(&A, &nData)
	fmt.Print("Berikut data yang sudah di input:\n")
	cetakData(A, nData)
	ascending(&A, nData)
	fmt.Print("Berikut data yang sudah di urutkan(Ascending):\n")
	cetakData(A, nData)
}

func isiArray(A *tabMhs, n *int) {
	var i int
	fmt.Scan(n)
	if *n > NMAX {
		*n = NMAX
	}

	for i = 0; i < *n; i++ {
		fmt.Scan(&A[i].sid, &A[i].name, &A[i].gpa)
	}
}

func ascending(A *tabMhs, n int) {
	var i, j, idx int

	for i = 0; i < n-1; i++ {
		idx = i
		for j = i + 1; j < n; j++ {
			if A[idx].gpa > A[j].gpa {
				idx = j
			}
		}
		A[i], A[idx] = A[idx], A[i]
	}
}

func cetakData(A tabMhs, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Println(A[i].sid, A[i].name, A[i].gpa)
	}
}
