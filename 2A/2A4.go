package main

import "fmt"

func main() {
	var celsius float32

	fmt.Print("Temperatur Celsius : ")
	fmt.Scanln(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32
	reamur := celsius * 4 / 5
	kelvin := (fahrenheit + 459.67) * 5 / 9

	hasilFahrenheit := int(fahrenheit)
	hasilReamur := int(reamur)
	hasilKelvin := int(kelvin)

	fmt.Printf("Derajat Reamur : %d \n", hasilReamur)
	fmt.Printf("Derajat Fahrenheit : %d \n", hasilFahrenheit)
	fmt.Printf("Derajat Kelvin : %d \n", hasilKelvin)
}
