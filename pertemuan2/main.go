package main

import "fmt"

func main() {

	// //condition
	// var point = 8
	// if point == 10 {
	// 	fmt.Println("Lulus dengan nilai sempurna")
	// } else if point > 5 {
	// 	fmt.Println("Lulus")
	// } else if point == 4 {
	// 	fmt.Println("Hampir lulus")
	// } else {
	// 	fmt.Printf("Tidak lulus nilai anda %v \n", point)
	// }

	// var points = 8840.0
	// if percent := points / 100; percent >= 100 {
	// 	fmt.Printf("%.1f%s perfect!\n", percent, "%")
	// } else if percent >= 70 {
	// 	fmt.Printf("%.1f%s good!\n", percent, "%")
	// } else {
	// 	fmt.Printf("%.1f%s not bad!\n", percent, "%")
	// }

	// switch point {
	// case 6:
	// 	fmt.Println("perfect")
	// case 7:
	// 	fmt.Println("awesome")
	// default:
	// 	fmt.Println("not bad")
	// }

	// switch point {
	// case 7:
	// 	fmt.Println("perfect")
	// case 6, 5, 4, 3:
	// 	fmt.Println("awesome")
	// default:
	// 	fmt.Println("not bad")
	// }

	// switch point {
	// case 7:
	// 	fmt.Println("perfect")
	// case 6, 5, 4, 3:
	// 	fmt.Println("awesome")
	// default:
	// 	{
	// 		fmt.Println("not bad")
	// 		fmt.Println("you can be better!")
	// 	}
	// }

	// switch {
	// case point == 8:
	// 	fmt.Println("perfect")
	// case (point < 8) && (point > 3):
	// 	fmt.Println("awesome")
	// default:
	// 	{
	// 		fmt.Println("not bad")
	// 		fmt.Println("you need learn more")
	// 	}
	// }

	// switch {
	// case point == 8:
	// 	fmt.Println("perfect")
	// case (point < 8) && (point > 3):
	// 	fmt.Println("awesome")
	// case point < 5:
	// 	fmt.Println("you need learn more")
	// default:
	// 	{
	// 		fmt.Println("not bad")
	// 		fmt.Println("you need learn more")
	// 	}
	// }

	// if point > 7 {
	// 	switch point {
	// 	case 10:
	// 		fmt.Println("perfect!")
	// 	default:
	// 		fmt.Println("nice!")
	// 	}

	// } else {
	// 	if point == 5 {
	// 		fmt.Println("not bad")
	// 	} else if point == 3 {
	// 		fmt.Println("keep trying")
	// 	} else {
	// 		fmt.Println("you can do it")
	// 		if point == 0 {
	// 			fmt.Println("try harder!")
	// 		}
	// 	}
	// }

	// //looping
	// for i :=0; i<5; i++{
	// 	fmt.Println("angka",i)
	// }

	// var i = 0

	// for i<5{
	// 	fmt.Println("Angka",i)
	// 	i++
	// }

	// for{
	// 	fmt.Println("Angka",i)
	// 	i++
	// 	if i==5{
	// 		break
	// 	}
	// }

	// for i:= 1;i<=10;i++{
	// 	if i % 2==1{
	// 		continue
	// 	}
	// 	if i>8{
	// 		break
	// 	}
	// 	fmt.Println("Angka",i)
	// }

	// for i:=0; i<5; i++{
	// 	for j:= i; j<5; j++{
	// 		fmt.Println(j," ")
	// 	}
	// 	fmt.
	// 	Println()
	// }

	// outerLoop:
	// 	for i := 0; i < 5; i++ {

	// 		for j := 0; j < 5; j++ {
	// 			if i == 3 {
	// 				break outerLoop
	// 			}
	// 			fmt.Print("matriks[", i, "][", j, "]", "\n")
	// 		}
	// 	}

	//arrays

	// var names [4]string
	// names[0] = "trafalgar"
	// names[1] = "d"
	// names[2] = "water"
	// names[3] = "law"

	// fmt.Println(names[0], names[1], names[2], names[3])

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// fmt.Println("Jumlah element \t\t", len(fruits))
	// fmt.Println("isi semua element \t", fruits)

	// var buah [4]string
	// // buah=[4]string{"apple","grape","banana","melon"}

	// buah = [4]string{
	// 	"apple",
	// 	"grape",
	// 	"banana",
	// 	"melon",
	// }
	// fmt.Println(buah)

	// var numbers = [...]int{2, 3, 2, 4, 3}

	// fmt.Println("data array\t:", numbers)
	// fmt.Println("jumlah elemen \t", len(numbers))

	// var z = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	// z[0][1] = 20
	// fmt.Println(z)

	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("elemen %d : %s\n", i, fruits[i])
	// }

	// for i, fruit := range fruits {
	// 	fmt.Printf("elemen %d : %s\n", i, fruit)
	// }

	// for _, fruit := range fruits {
	// 	fmt.Printf("nama buah : %s\n", fruit)
	// }

	//slice

	// var fruits = []string{"apple", "grape", "banana", "melon"}
	// fmt.Println((fruits[0]))
	// var newFruits = fruits[0:2]
	// fmt.Println(newFruits)

	// var x []int
	// fmt.Println(x == nil)

	// var aFruits = fruits[0:3]
	// var bFruits = fruits[1:4]

	// var aaFruits = fruits[1:2]
	// var baFruits = fruits[0:1]

	// fmt.Println(fruits)
	// fmt.Println(aFruits)
	// fmt.Println(bFruits)
	// fmt.Println(aaFruits)
	// fmt.Println(baFruits)

	// baFruits[0] = "pinnaple"

	// fmt.Println(fruits)
	// fmt.Println(aFruits)
	// fmt.Println(bFruits)
	// fmt.Println(aaFruits)
	// fmt.Println(baFruits)

	// fmt.Println(len(fruits))
	// fmt.Println(cap(fruits))

	// fmt.Println(len(aFruits))
	// fmt.Println(cap(aFruits))

	// fmt.Println(len(bFruits))
	// fmt.Println(cap(bFruits))

	// var cFruits = append(fruits, "papaya")

	// fmt.Println(fruits)
	// fmt.Println(cFruits)

	// fmt.Println(cap(bFruits))
	// fmt.Println(len(bFruits))

	// fmt.Println(fruits)
	// fmt.Println(bFruits)

	// var dFruits = append(bFruits, "papaya")

	// fmt.Println(fruits)
	// fmt.Println(bFruits)
	// fmt.Println(dFruits)

	// dst := make([]string, 3)
	// src := []string{"watermelon", "pinnaple", "apple", "orange"}
	// n := copy(dst, src)

	// fmt.Println(dst)
	// fmt.Println(src)
	// fmt.Println(n)

	// var eFruits = fruits[0:2:2]

	// fmt.Println(fruits)
	// fmt.Println(len(fruits))
	// fmt.Println(cap(fruits))

	// fmt.Println(eFruits)
	// fmt.Println(len(eFruits))
	// fmt.Println(cap(eFruits))

	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40
	chicken["maret"] = 34
	chicken["april"] = 67

	fmt.Println("januari", chicken["januari"])
	fmt.Println("mei", chicken["mei"])

	for key, val := range chicken {
		fmt.Println(key, "\t:", val)
	}

	delete(chicken, "januari")
	// var chicken3 = map[string]int{}
	// var chicken4=make(map[string]int)
	// var chicken5 = *new(map[string]int)

	var value, isExist = chicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, ayam := range chickens {
		fmt.Println(ayam["gender"], ayam["name"])
	}
}
