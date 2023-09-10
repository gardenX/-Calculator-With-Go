package main

import "fmt"

func add(num1 int, num2 int) (result int) {
	result = num1 + num2
	return
}

func sub(num1 int, num2 int) (result int) {
	result = num1 - num2
	return
}

func times(num1 int, num2 int) (result int) {
	result = num1 * num2
	return
}

func div(num1 int, num2 int) (result int) {
	result = num1 / num2
	return
}

func operatorProcces(operator string) {
	var num1, num2 int

	fmt.Println("Masukan Angka Pertama :")
	fmt.Scan(&num1)
	fmt.Println("Masukan Angka Kedua :")
	fmt.Scan(&num2)

	switch operator {
	case "1":
		fmt.Println("Hasil Penjumlahan Dari", num1, "+", num2, " Adalah ", add(num1, num2))
		mathOperator("Ya")
	case "2":
		fmt.Println("Hasil Pengurangan Dari", num1, "-", num2, " Adalah ", sub(num1, num2))
	case "3":
		res := times(num1, num2)
		fmt.Println("Hasil Perkalian Dari", num1, "*", num2, "Adalah", res)
	case "4":
		res := div(num1, num2)
		fmt.Println("Hasil Pembagian Dari", num1, "/", num2, "Adalah", res)
	default:
		fmt.Println("Invalid operator")
		return
	}

}

func mathOperator(par1 string) {
	if par1 == "Ya" {
		fmt.Println("Silahkan Pilih Operator Penjumlahan Yang Diinginkan !")
		fmt.Println("1. + (Penjumlahan)")
		fmt.Println("2. - (Pengurangan)")
		fmt.Println("3. * (Perkalian)")
		fmt.Println("4. / (Pembagian)")

		var operator string
		fmt.Scan(&operator)
		operatorProcces(operator)
	} else {
	}
}

func abtApps(salam int) {
	if salam == 1 {
		fmt.Println("Selamat Datang D Aplikasi Kalkulator Sederhana ")
		fmt.Println("Apakah Anda Ingin Melanjutkan ? ")
		fmt.Println("Ya ? ")
		fmt.Println("Tidak ? ")

		var params string
		fmt.Scan(&params)
		mathOperator(params)
	}
}

func main() {
	fmt.Println("Assalamualaikum Teman ?")
	fmt.Println("1. Waalaikum Salam ")
	fmt.Println("2. Exit ")

	var params int
	fmt.Scan(&params)
	abtApps(params)
}
