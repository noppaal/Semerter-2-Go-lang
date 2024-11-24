package main

import "fmt"

const NMax int = 100

type tabChar [NMax]byte

func masukanArray(T *tabChar, n *int) {
	/*I.S. Data tersedia pada piranti masukan
	  Proses: Masukan akan terus berlangsung dan berhenti ketika pengguna
	  memasukkan '.'
	  F.S. Sekumpulan karakter sebanyak n berada dalam array T
	  Petunjuk: Gunakan while-loop untuk melakukan proses input
	*/
	var char byte
	i := 0
	fmt.Scanf("%c", &char)
	for char != '.' {
		(*T)[i] = char
		fmt.Scanf("%c", &char)
		i++
	}
	*n = i
}

func upperCase(T *tabChar, n int) {
	/*I.S. Terdefinisi sekumpulan n karakter pada array T
	  F.S. Seluruh anggota karakter pada array T dikonversi menjadi huruf kapital
	  Petunjuk: Gunakan ASCII, apabila sudah huruf kapital, tidak perlu diubah
	*/
	for i := 0; i < n; i++ {
		if (*T)[i] >= 'a' && (*T)[i] <= 'z' {
			(*T)[i] -= 32
		}
	}
}

func cetakArray(T tabChar, n int) {
	/*I.S. Terdefinisi sekumpulan n karakter pada array T
	  F.S. Menampilkan seluruh elemen array T pada layar
	*/
	for i := 0; i < n; i++ {
		fmt.Printf("%c", T[i])
	}
	fmt.Println()
}

func main() {
	var T tabChar
	var n int
	masukanArray(&T, &n)
	upperCase(&T, n)
	cetakArray(T, n)
}
