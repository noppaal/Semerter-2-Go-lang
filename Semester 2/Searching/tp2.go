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

	binarySearch(data, nData, x2, &idx)

	fmt.Print("Hasil pencarian: ")
	if idx != -1 {
		fmt.Printf("Bilangan %d ditemukan pada posisi ke-%d\n", x2, idx)
	} else {
		fmt.Println("Bilangan tidak ditemukan.")
	}
}

func bacaData(A *tabInt, n *int) {
	var i int = 0

	for i < NMAX {
		fmt.Scan(&A[i])

		if i == 0 || A[i] >= A[i-1] {
			i++
		} else {
			break
		}
	}
	*n = i
}

func cetakData(A tabInt, n int) {
	fmt.Print("Data Bilangan: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", A[i])
	}
	fmt.Println()
}

func binarySearch(A tabInt, n, x int, idx *int) {
	low := 0
	high := n - 1

	for low <= high {
		mid := (low + high) / 2

		if A[mid] == x {
			*idx = mid + 1
			return
		} else if A[mid] < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	*idx = -1
}
