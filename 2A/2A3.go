package main

import "fmt"

func main() {
	var jariJari int
	const pi = 3.1415926535

	fmt.Print("Jejari = ")
	fmt.Scanln(&jariJari)

	jariJariFloat := float64(jariJari)

	pangkatTiga := jariJariFloat * jariJariFloat * jariJariFloat
	pangkatDua := jariJariFloat * jariJariFloat

	volume := (4.0 / 3.0) * pi * pangkatTiga
	luasPermukaan := 4 * pi * pangkatDua

	fmt.Printf("Bola dengan jari-jari %d memiliki volume %.2f dan luas permukaan %.2f\n", jariJari, volume, luasPermukaan)
}
