package main

import "fmt"

const NMAX int = 1000

type tabBunga [NMAX]string

func main() {
	var A tabBunga
	var n int
	isiArray(&A, &n)
	mengurutkan(&A, n)
	tampilArray(A, n)
}

func panjang(bunga string) int {
	count := 0
	for i := 0; i < len(bunga); i++ {
		if bunga[i] != '.' && bunga[i] != '_' {
			count++
		}
	}
	return count
}

func mengurutkan(A *tabBunga, n int) {
	var i, j int
	var tmp string

	i = 1
	for i <= n-1 {
		j = i
		tmp = A[i]
		for j > 0 && panjang(tmp) < panjang(A[j-1]) {
			A[j] = A[j-1]
			j--
		}
		A[j] = tmp
		i++
	}
}

func isiArray(A *tabBunga, n *int) {
	var i int

	fmt.Scan(n)
	if *n > NMAX {
		*n = NMAX
	}
	for i = 0; i < *n; i++ {
		fmt.Scan(&A[i])
	}
}

func tampilArray(A tabBunga, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Println(A[i])
	}
}
