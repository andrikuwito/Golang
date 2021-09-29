package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Items struct {
	Item_id     int    `json:"lineItemId"`
	Item_code   string `json :"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type Orders struct {
	Order_id      int     `json:"orderId"`
	Customer_name string  `json:"customerName"`
	Ordered_at    string  `json:"orderedAt"`
	Item          []Items `json: "items"`
}

var (
	username   = "root"
	password   = "root"
	hostname   = "127.0.0.1"
	dbname     = "orders_by"
	connection = fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	CreateOrder()
	UpdateOrder()
	Delete()
	http.ListenAndServe(":8080", nil)
}
