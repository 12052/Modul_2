package main

import "fmt"

func main() {
	var angka [5]int
	var karakter [4]byte

	for i := 0; i < 5; i++ {
		fmt.Scan(&angka[i])
	}

	for i := 0; i < 4; i++ {
		fmt.Scanf("%c", &karakter[i])
	}

	for _, v := range angka {
		fmt.Printf("%c", v)
	}
	fmt.Print()

	for _, v := range karakter {
		fmt.Printf("%c", v+1)
	}
	fmt.Print()
}
