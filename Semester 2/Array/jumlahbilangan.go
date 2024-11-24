package main

import "fmt"

const NMax = 1000

func jumlahBilangan(bilangan []int) int {
	jumlah := 0
	for _, n := range bilangan {
		if n == 0 {
			break
		}
		jumlah += n
	}
	return jumlah
}

func main() {
	var bilangan []int
	var n int

	fmt.Scan(&n)
	for n != 0 {
		bilangan = append(bilangan, n)
		fmt.Scan(&n)
	}
	result := jumlahBilangan(bilangan)
	fmt.Println(result)
}
