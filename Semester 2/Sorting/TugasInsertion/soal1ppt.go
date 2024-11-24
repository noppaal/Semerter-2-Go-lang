package main

import "fmt"

const NMAX int = 1000

type tabInt [NMAX]int

func main() {
	var A tabInt
	var n int

	fmt.Scanln(&n)
	inputArray(&A, &n)
	pengurutan(&A, n)
	cariPosMed(A, n)
}

func inputArray(A *tabInt, n *int) {
	if *n > NMAX {
		*n = NMAX
	}

	for i := 0; i < *n; i++ {
		fmt.Scan(&A[i])
	}
}

func pengurutan(A *tabInt, n int) {
	var i, j, tmp int

	i = 1
	for i <= n-1 {
		j = i
		tmp = A[i]
		for j > 0 && tmp < A[j-1] {
			A[j] = A[j-1]
			j--
		}
		A[j] = tmp
		i++
	}
}

func cariPosMed(A tabInt, n int) {
	var med float64

	if n%2 == 0 {
		med = (float64(n/2) + float64((n/2)+1)) / 2
	} else {
		med = float64((n + 1) / 2)
	}
	fmt.Print(med)
}
