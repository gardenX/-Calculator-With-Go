package main

import "fmt"

func penjumlahan(num1 int8, num2 int8) {
	fmt.Println("Hasil Penjumlahan Dari", num1, "+", num2, "Adalah", num1+num2)
}

func pengurangan(num1 int8, num2 int8) {
	fmt.Println("Hasil Pengurangan Dari", num1, "-", num2, "Adalah", num1-num2)
}

func perkalian(num1 int8, num2 int8) {
	fmt.Println("Hasil Perkalian Dari", num1, "*", num2, "Adalah", num1*num2)
}

func pembagian(num1 int8, num2 int8) {
	fmt.Println("Hasil Pembagian Dari", num1, "/", num2, "Adalah", num1/num2)
}

func operatirProses(operator string) {
	var num1, num2 int8
	fmt.Println("Masukan Angka Pertama :")
	fmt.Scan(&num1)
	fmt.Println("Masukan Angka Kedua :")
	fmt.Scan(&num2)

	switch operator {
	case "1":
		penjumlahan(num1, num2)
	case "2":
		pengurangan(num1, num2)
	case "3":
		perkalian(num1, num2)
	case "4":
		pembagian(num1, num2)
	default:
		fmt.Println("Invalid operator")
	}
}

func lanjut(salam int) {
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
			operatirProses(operator)

		} else {
		}

	}
}

func main() {
	fmt.Println("Assalamualaikum Teman ?")
	fmt.Println("1. Waalaikum Salam ")
	fmt.Println("2. Exit ")

	var salam int
	fmt.Scan(&salam)
	lanjut(salam)
}
