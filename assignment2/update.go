package main

import "fmt"

func UpdateOrder() {
	db, err := ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()
	var orderID = Orders{}

	err = db.QueryRow("select order_id from orders").Scan(&orderID.Order_id)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = db.Exec("update orders set customer_name=? ,ordered_at=? where order_id=?", "Spike Tyke", "2019-11-09T21:21:46", orderID.Order_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("update Success")

	_, err = db.Exec("update items set item_id=? ,quantity=? where order_id=?", 1, 10, orderID.Order_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("update Success")
}
