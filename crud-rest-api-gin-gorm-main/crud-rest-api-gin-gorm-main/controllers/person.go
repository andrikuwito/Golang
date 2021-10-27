package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api-km/structs"
)

func (Conn *DBConn) CreatePerson(c *gin.Context){
	var person structs.Person

	person.FirstName = c.PostForm("first_name")
	person.LastName = c.PostForm("last_name")

	Conn.DB.Create(&person)

	result := gin.H{
		"result": person,
	}

	c.JSON(http.StatusOK, result)
}

func (Conn *DBConn) GetPersonById(c *gin.Context){
	var person structs.Person
	var result gin.H

	id := c.Param("id")
	err := Conn.DB.Where("id = ?", id).First(&person).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": person,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (Conn *DBConn) GetPersons(c *gin.Context){
	var persons []structs.Person
	var result gin.H

	Conn.DB.Find(&persons)

	if len(persons) <= 0 {
		result = gin.H{
			"result": nil,
		}
	} else {
		result = gin.H{
			"result": persons,
		}
	}
	c.JSON(http.StatusOK, result)
}


func (Conn *DBConn) UpdatePerson(c *gin.Context){
	id 			:= c.Query("id")

	firstName 	:= c.PostForm("first_name")
	lastName 	:= c.PostForm("last_name")

	var (
		person structs.Person
		newPerson structs.Person
		result gin.H
	)

	err := Conn.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}

	newPerson.FirstName = firstName
	newPerson.LastName = lastName

	err = Conn.DB.Model(&person).Updates(newPerson).Error
	if err != nil {
		result = gin.H{
			"result": "Update failed",
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	result = gin.H{
		"result": "Update data success",
	}
	c.JSON(http.StatusOK, result)
	return

}

func (Conn *DBConn) DeletePerson (c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)
	id := c.Param("id")
	err := Conn.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}

	err = Conn.DB.Delete(&person).Error
	if err != nil {
		result = gin.H{
			"result": "Delete failed",
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	result = gin.H{
		"result": "Delete data success",
	}
	c.JSON(http.StatusOK, result)
	return
}


