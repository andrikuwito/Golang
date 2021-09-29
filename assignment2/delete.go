package main

import (
	"fmt"
	"net/http"
)

func Delete() {

	http.HandleFunc("/orders/1", func(w http.ResponseWriter, r *http.Request) {
		db, err := ConnectDB()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		_, err = db.Exec("Delete from orders where order_id=?", "1")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		_, err = db.Exec("Delete from items where order_id=?", "1")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Fprintln(w, "Delete Success")
	})
}
