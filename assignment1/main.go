package main

import (
	"fmt"
	"os"
	"strconv"
)

type student struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	student1 := []student{

		{nama: "Firman Aulia Insani",
			alamat:    "Belum tau",
			pekerjaan: "Mahasiswa",
			alasan:    "dengan course ini saya berharap bisa mempelajari bidang baru"},

		{nama: "Andri Nut Hidayatulloh",
			alamat:    "Belum tau",
			pekerjaan: "Mahasiswa",
			alasan:    "Saya ikut course ini berharap ada product yg bisa saya buat dan membantu masyarakat"},

		{nama: "I Komang Widnyana",
			alamat:    "Belum tau",
			pekerjaan: "Mahasiswa",
			alasan:    "dengan course ini ku berharap golang bisa memantapkan kemampuan ku dibidang fullstack"},

		{nama: "Erico",
			alamat:    "Belum tau",
			pekerjaan: "Mahasiswa",
			alasan:    "Belum tau"},

		{nama: "Jose Yolanda Purba",
			alamat:    "Belum tau",
			pekerjaan: "Mahasiswa",
			alasan:    "Belum tau"},

		{nama: "Andri Kuwito",
			alamat:    "Bogor",
			pekerjaan: "Mahasiswa",
			alasan:    "Ingin mempelajari bahasa golang"},

		{nama: "Sandy Budiman",
			alamat:    "Belum tau",
			pekerjaan: "Mahasiswa",
			alasan:    "Belum tau"},

		{nama: "Rafli Andreansyah",
			alamat:    "Belum tau",
			pekerjaan: "Mahasiswa",
			alasan:    "Belum tau"},

		{nama: "Taufiq Hidayah",
			alamat:    "Belum tau",
			pekerjaan: "Mahasiswa",
			alasan:    "Belum tau"},

		{nama: "Melvita Sari",
			alamat:    "Belum tau",
			pekerjaan: "Mahasiswa",
			alasan:    "Berkeinginan menjadi seorang yang memiliki daya ingat yang kuat"},
	}

	args := os.Args[1]
	number, _ := strconv.Atoi(args)
	fmt.Println(student1[number-1])

}
