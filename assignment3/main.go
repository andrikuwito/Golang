package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type status struct {
	Water int `json: "water"`
	Wind  int `json: "wind"`
}

type input struct {
	Status status `json :"status"`
}

var data = input{}
var air string
var angin string

func ulang() {
	for {
		data = input{
			status{Water: randNumber(), Wind: randNumber()},
		}
		fmt.Println(data)
		file, _ := json.MarshalIndent(data, "", " ")
		_ = ioutil.WriteFile("data.json", file, 0644)
		time.Sleep(5 * time.Second)
	}
}

func main() {
	go ulang()
	http.HandleFunc("/tes", func(w http.ResponseWriter, r *http.Request) {

		jsonFile, err := os.Open("data.json")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully Opened users.json")
		t, err := template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for {
			if data.Status.Water < 5 {
				fmt.Println("AMAN")
				air = "AMAN"
				break
			} else if data.Status.Water >= 6 && data.Status.Water <= 8 {
				fmt.Println("SIAGA")
				air = "SIAGA"
				break
			} else if data.Status.Water > 8 {
				fmt.Println("BAHAYA")
				air = "BAHAYA"
				break
			}
		}

		for {
			if data.Status.Wind < 6 {
				fmt.Println("AMAN")
				angin = "AMAN"
				break
			} else if data.Status.Wind >= 7 && data.Status.Wind <= 15 {
				fmt.Println("SIAGA")
				angin = "SIAGA"
				break
			} else if data.Status.Wind > 15 {
				fmt.Println("BAHAYA")
				angin = "BAHAYA"
				break
			}
		}
		hasil := map[string]interface{}{
			"Water": data.Status.Water,
			"Wind":  data.Status.Wind,
			"Air":   air,
			"Angin": angin,
		}
		t.Execute(w, hasil)
		defer jsonFile.Close()
	})
	http.ListenAndServe(":8080", nil)
}

func randNumber() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	randNum := rand.Intn(max-min+1) + min
	return randNum
}
