package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"rest-api-km/config"
	"rest-api-km/controllers"
)

func main(){

	db := config.InitDB()
	DBConn := &controllers.DBConn{DB:db}
	router := gin.Default()

	router.POST("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	// CRUD - Create, Read, Update & Delete

	//Create
	router.POST("/person", DBConn.CreatePerson)

	//Read -- localhost:8080/person/1 --> Param
	router.GET("/person/:id", DBConn.GetPersonById)
	router.GET("/persons", DBConn.GetPersons)

	//Update -- localhost:8080/person?id=1 --> Query
	router.PUT("/person", DBConn.UpdatePerson)

	//Delete
	router.DELETE("/person/:id", DBConn.DeletePerson)


	router.Run(":8080")
}