package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateOrder() {
	db, err := ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	_, err = db.Exec("Insert into orders (ordered_at,customer_name) values(?,?)", "2019-11-09T21:21:46+00:00", "Tom Jerry")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Insert Success")

	var orderID = Orders{}

	err = db.QueryRow("select order_id from orders").Scan(&orderID.Order_id)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = db.Exec("Insert into items (item_code,description,quantity,order_id) values(?,?,?,?)", "123", "Iphone 10X", 1, orderID.Order_id)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Insert Success")

	var result = Orders{}
	var itemData = Items{}

	err = db.QueryRow("select Orders.ordered_at, orders.customer_name, items.item_code , items.description, items.quantity FROM Orders INNER JOIN items ON Orders.order_id=items.order_id").Scan(&result.Ordered_at, &result.Customer_name, &itemData.Item_code, &itemData.Description, &itemData.Quantity)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var data = []Orders{
		{Ordered_at: result.Ordered_at, Customer_name: result.Customer_name, Item: []Items{{Item_code: itemData.Item_code, Description: itemData.Description, Quantity: itemData.Quantity}}},
	}

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		if r.Method == "GET" {
			fmt.Println(data)
			res, err := json.Marshal(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			fmt.Println(res)
			w.Write(res)
			return
		}
		http.Error(w, "Bad Request", http.StatusBadRequest)
	})
}
