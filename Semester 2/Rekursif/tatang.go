for i := 0; i < NMAX; i++ {
	if T.tabInt[i] == 'a' {
		total = total + 1
	} else if T.tabInt[i] == 'b' {
		total = total + 2
	} else {
		total = total + 3
	}
}