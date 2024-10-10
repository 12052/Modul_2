package main

import "fmt"

func main() {
	var warna [5][4]string
	var warnaIdeal = []string{"merah", "kuning", "hijau", "ungu"}
	var berhasil bool

	for i := 0; i < 5; i++ {
		fmt.Printf("Percobaan %d: ", i+1)
		for j := 0; j < 4; j++ {
			fmt.Scan(&warna[i][j])
		}
	}

	berhasil = true
	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			if warna[i][j] != warnaIdeal[j] {
				berhasil = false
				break
			}
		}
		if !berhasil {
			break
		}
	}

	fmt.Println("BERHASIL:", berhasil)
}
