package main

import "fmt"

func main() {
	type club struct {
		nama              string
		point, selisihGol int
	}

	var c1 club
	var c2 club
	var c3 club

	fmt.Scanln(&c1.nama, &c1.point, &c1.selisihGol)
	fmt.Scanln(&c2.nama, &c2.point, &c2.selisihGol)
	fmt.Scanln(&c3.nama, &c3.point, &c3.selisihGol)

	pointTerbesar := c1.point

	if c2.point > pointTerbesar {
		fmt.Println(c2.nama)
	} else if c3.point > pointTerbesar {
		fmt.Println(c3.nama)
	} else {
		fmt.Println(c1.nama)
	}

	selisihGolTerbesar := c1.selisihGol
	if c2.selisihGol > selisihGolTerbesar {
		fmt.Println(c2.nama)
	} else if c3.selisihGol > selisihGolTerbesar {
		fmt.Println(c3.nama)
	} else {
		fmt.Println(c1.nama)
	}
}
