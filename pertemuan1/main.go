package main

import "fmt"

func main() {
	fmt.Print("Hello World - 2")
	fmt.Println("Hello World - 1")

	//variable
	var name string = "Andri"
	fmt.Println(name)

	var tanggal = 24
	fmt.Println(tanggal)

	umur := 21
	fmt.Println(umur)
	umur = 99
	fmt.Println(umur)

	var kota, universitas string = "Jakarta", "UNTAR"
	fmt.Println(kota, universitas)

	firstname, lastname, year := "Andri", "Kuwito", 2021
	fmt.Println(firstname, lastname, year)

	_, lastName, Year := "Andri", "Kuwito", 2021
	fmt.Println(lastName, Year)

	hacktiv := new(string)
	fmt.Println(hacktiv)

	var zeroInt int
	var zeroStr string
	var zeroBool bool
	fmt.Println(zeroInt, zeroStr, zeroBool)

	var formatInt int = 10
	var formatStr string = "Fundamental GO"
	var formatBool bool = true
	var formatFloat float32 = 3.14
	fmt.Println(formatInt, formatStr, formatBool, formatFloat)

	const phi float32 = 3.14
	fmt.Println(phi)

	//aritmatika
	x := 10
	y := 20
	penjumlahan := x + y
	pengurangan := x - y
	perkalian := x * y
	pembagian := x / y
	modulo := y % x

	fmt.Printf("Penjumlahan : %v\n", penjumlahan)
	fmt.Printf("Pengurangan : %v\n", pengurangan)
	fmt.Printf("Perkalian : %v\n", perkalian)
	fmt.Printf("Pembagian : %v\n", pembagian)
	fmt.Printf("Modulo : %v\n", modulo)

	//perbandingan
	fmt.Println(x == y)
	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)
	fmt.Println(x != y)

	//logika
	kondisiTrue := true
	kondisiFalse := false

	fmt.Println(kondisiTrue && kondisiFalse)
	fmt.Println(kondisiTrue || kondisiFalse)
	fmt.Println(!kondisiFalse)
}
