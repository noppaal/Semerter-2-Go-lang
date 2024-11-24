package main

import "fmt"

const NMAX int = 1000

type peserta struct {
	nama    string
	g, s, b int
}

type tabPes [NMAX]peserta

func main() {
	var A tabPes
	var n int
	isiArray(&A, &n)
	sorting(&A, n)
	tampilArray(A, n)
}

func isiArray(A *tabPes, n *int) {
	var i int

	fmt.Scan(n)
	if *n > NMAX {
		*n = NMAX
	}
	for i = 0; i < *n; i++ {
		fmt.Scan(&A[i].nama, &A[i].g, &A[i].s, &A[i].b)
	}
}

func tampilArray(A tabPes, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Println(A[i].nama, A[i].g, A[i].s, A[i].b)
	}
}

func poin(g, s, b int) int {
	return (4 * g) + (3 * s) + b
}

func sorting(A *tabPes, n int) {
	var i, j int
	var tmp peserta

	i = 1
	for i <= n-1 {
		j = i
		tmp = A[i]
		for (j > 0 && poin(tmp.g, tmp.s, tmp.b) > poin(A[j-1].g, A[j-1].s, A[j-1].b)) || (j > 0 && poin(tmp.g, tmp.s, tmp.b) == poin(A[j-1].g, A[j-1].s, A[j-1].b) && tmp.g > A[j-1].g) || (j > 0 && poin(tmp.g, tmp.s, tmp.b) == poin(A[j-1].g, A[j-1].s, A[j-1].b) && tmp.g == A[j-1].g && tmp.s > A[j-1].s) || (j > 0 && poin(tmp.g, tmp.s, tmp.b) == poin(A[j-1].g, A[j-1].s, A[j-1].b) && tmp.g == A[j-1].g && tmp.s == A[j-1].s && tmp.b > A[j-1].b) {
			A[j] = A[j-1]
			j--
		}
		A[j] = tmp
		i++
	}
}
