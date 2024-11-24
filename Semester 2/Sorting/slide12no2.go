package main

import (
	"fmt"
)

const MAX int = 30

type IntArray struct {
	tabInt [MAX]int
	N      int
}

var array1, array2 IntArray

func main() {
	inputArray(&array1)
	inputArray(&array2)
	sortArray(&array1)
	sortArray(&array2)
	fmt.Print(isSimilar(array1, array2))
}

func inputArray(T *IntArray) {
	var input, i int
	i = 0
	fmt.Scan(&input)
	for i < MAX && input != 0 {
		T.tabInt[i] = input
		i++
		fmt.Scan(&input)
	}
	T.N = i
}

func sortArray(T *IntArray) {
	var i, j, idx int

	for i = 0; i < T.N; i++ {
		idx = i
		for j = i + 1; j < T.N-1; j++ {
			if T.tabInt[idx] > T.tabInt[j] {
				idx = j
			}
		}
		T.tabInt[i], T.tabInt[idx] = T.tabInt[idx], T.tabInt[i]
	}
}

func isSimilar(T1, T2 IntArray) bool {
	var i int

	if T1.N != T2.N {
		return false
	}

	for i = 0; i < T1.N; i++ {
		if T1.tabInt[i] != T2.tabInt[i] {
			return false
		}
	}
	return true
}
