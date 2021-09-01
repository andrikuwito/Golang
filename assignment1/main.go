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

		{nama: "Firman Aulia Insani \n",
			alamat:    "Belum tau \n",
			pekerjaan: "Mahasiswa \n",
			alasan:    "Dengan course ini saya berharap bisa mempelajari bidang baru"},

		{nama: "Andri Nut Hidayatulloh \n",
			alamat:    "Belum tau \n",
			pekerjaan: "Mahasiswa \n",
			alasan:    "Saya ikut course ini berharap ada product yg bisa saya buat dan membantu masyarakat"},

		{nama: "I Komang Widnyana \n",
			alamat:    "Belum tau \n",
			pekerjaan: "Mahasiswa \n",
			alasan:    "Dengan course ini ku berharap golang bisa memantapkan kemampuan ku dibidang fullstack"},

		{nama: "Erico \n",
			alamat:    "Belum tau \n",
			pekerjaan: "Mahasiswa \n",
			alasan:    "Belum tau \n"},

		{nama: "Jose Yolanda Purba \n",
			alamat:    "Belum tau \n",
			pekerjaan: "Mahasiswa \n",
			alasan:    "Belum tau"},

		{nama: "Andri Kuwito \n",
			alamat:    "Bogor \n",
			pekerjaan: "Mahasiswa \n",
			alasan:    "Ingin mempelajari bahasa golang"},

		{nama: "Sandy Budiman \n",
			alamat:    "Belum tau \n",
			pekerjaan: "Mahasiswa \n",
			alasan:    "Belum tau"},

		{nama: "Rafli Andreansyah \n",
			alamat:    "Belum tau \n",
			pekerjaan: "Mahasiswa \n",
			alasan:    "Belum tau"},

		{nama: "Taufiq Hidayah \n",
			alamat:    "Belum tau \n",
			pekerjaan: "Mahasiswa \n",
			alasan:    "Belum tau"},

		{nama: "Melvita Sari \n",
			alamat:    "Belum tau \n",
			pekerjaan: "Mahasiswa \n",
			alasan:    "Berkeinginan menjadi seorang yang memiliki daya ingat yang kuat"},
	}

	args := os.Args[1]
	number, _ := strconv.Atoi(args)
	fmt.Println(student1[number-1])

}
