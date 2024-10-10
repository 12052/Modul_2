package main

import "fmt"

func main() {

	/* variabel satu, dua, dan tiga dideklarasikan sebagai string, yang akan menyimpan input pengguna
	   temp adalah variabel string sementara yang digunakan untuk menukar nilai-nilai*/
	var (
		satu, dua, tiga string
		temp            string
	)

	//User menginput nilai dari variabel satu, dua, tiga
	fmt.Print("Masukkan input string : ")
	fmt.Scanln(&satu)
	fmt.Print("Masukkan input string : ")
	fmt.Scanln(&dua)
	fmt.Print("Masukkan input string : ")
	fmt.Scanln(&tiga)

	//Menampilkan hasil inputan user (Tanpa perubahan urutan)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)

	//Melakukan perubahan urutan setiap variabel
	temp = satu //Nilai satu diubah ke temp
	satu = dua  //Nilai dua diubah ke satu
	dua = tiga  //Nilai tiga diubah ke dua
	tiga = temp //Nilai temp (Awalnya satu) diubah ke tiga

	//Menampilkan hasil perubahan nilai pada variabel satu, dua, dan tiga
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
