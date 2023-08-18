package main

func main() {
	for i := 1; i < 100; i *= 2 {
		i += 4

		if i%2 == 0 {
			println(i)
		}
	}

	// i := 1
	// i *= 2
	// i += 4
	// fmt.Println(i)
	// ia := 6
	// ia *= 2
	// ia += 4
	// fmt.Println(ia)
	// d := 16
	// d *= 2
	// d += 4
	// fmt.Println(d)
	// w := 36
	// w *= 2
	// w += 4
	// fmt.Println(w)
}
