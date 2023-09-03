package main

import (
	"fmt"
)

func Main() {
	fmt.Println("Assalamualaikum Teman ?")
	fmt.Println("1. Waalaikum Salam ")
	fmt.Println("2. Exit ")

	var salam int8
	fmt.Scan(&salam)
	if salam == 1 {
		fmt.Println("Selamat Datang D Aplikasi Kalkulator Sederhana ")
		fmt.Println("Apakah Anda Ingin Melanjutkan ? ")
		fmt.Println("Ya ? ")
		fmt.Println("Tidak ? ")

		var lanjut string
		fmt.Scan(&lanjut)
		if lanjut == "Ya" {
			fmt.Println("Silahkan Pilih Operator Penjumlahan Yang Diinginkan !")
			fmt.Println("1. + (Penjumlahan)")
			fmt.Println("2. - (Pengurangan)")
			fmt.Println("3. * (Perkalian)")
			fmt.Println("4. / (Pembagian)")

			var operator string
			fmt.Scan(&operator)

			var num1, num2 int8
			fmt.Println("Masukan Angka Pertama :")
			fmt.Scan(&num1)
			fmt.Println("Masukan Angka Kedua :")
			fmt.Scan(&num2)

			switch operator {
			case "1":
				fmt.Println("Hasil Penjumlahan Dari", num1, "+", num2, "Adalah", num1+num2)
			case "2":
				fmt.Println("Hasil Penjumlahan Dari", num1, "-", num2, "Adalah", num1-num2)
			case "3":
				fmt.Println("Hasil Penjumlahan Dari", num1, "*", num2, "Adalah", num1*num2)
			case "4":
				fmt.Println("Hasil Penjumlahan Dari", num1, "/", num2, "Adalah", num1/num2)
			default:
				fmt.Println("Invalid operator")
			}
		} else {
		}

	} else {

	}

}

func main() {
	// agung.Calc()
	// Main()
}
