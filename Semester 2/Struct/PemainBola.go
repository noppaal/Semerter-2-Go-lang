package main

import "fmt"

func main() {
	type skor struct {
		nama        string
		gol, assist int
	}

	var s1 skor
	var s2 skor
	var s3 skor

	fmt.Scanln(&s1.nama, &s1.gol, &s1.assist)
	fmt.Scanln(&s2.nama, &s2.gol, &s2.assist)
	fmt.Scanln(&s3.nama, &s3.gol, &s3.assist)

	if s1.gol < s2.gol && s1.gol < s3.gol {
		fmt.Println(s1.nama)
	} else if s2.gol < s1.gol && s2.gol < s3.gol {
		fmt.Println(s2.nama)
	} else if s3.gol < s2.gol && s3.gol < s1.gol {
		fmt.Println(s3.nama)
	}

	if s1.assist < s2.assist && s1.assist < s3.assist {
		fmt.Println(s1.nama)
	} else if s2.assist < s1.assist && s2.assist < s3.assist {
		fmt.Println(s2.nama)
	} else if s3.assist < s2.assist && s3.assist < s1.assist {
		fmt.Println(s3.nama)
	}

}
