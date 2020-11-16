package controllers

import (
	"net/http"

	"../structs"
	"github.com/gin-gonic/gin"
)

// to get one data with {id}
func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&person).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// get all data in person
func (idb *InDB) GetPersons(c *gin.Context) {
	var (
		persons []structs.Person
		result  gin.H
	)

	idb.DB.Find(&persons)
	if len(persons) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": persons,
			"count":  len(persons),
		}
	}

	c.JSON(http.StatusOK, result)
}

// membuat data baru database
func (idb *InDB) CreatePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)
	nama_pertama := c.PostForm("nama_pertama")
	nama_belakang := c.PostForm("nama_belakang")
	person.First_Name = nama_pertama
	person.Last_Name = nama_belakang
	idb.DB.Create(&person)
	result = gin.H{
		"result": person,
	}
	c.JSON(http.StatusOK, result)
}

// update data dari {id} query
func (idb *InDB) UpdatePerson(c *gin.Context) {
	id := c.Query("id")
	nama_pertama := c.PostForm("nama_pertama")
	nama_belakang := c.PostForm("nama_belakang")
	var (
		person    structs.Person
		newPerson structs.Person
		result    gin.H
	)

	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "data tidak ditemukan",
		}
	}
	newPerson.First_Name = nama_pertama
	newPerson.Last_Name = nama_belakang
	err = idb.DB.Model(&person).Updates(newPerson).Error
	if err != nil {
		result = gin.H{
			"result": "update gagal",
		}
	} else {
		result = gin.H{
			"result": "sukses updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

// delete data dari {id}
func (idb *InDB) DeletePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "data tidak ditemukan",
		}
	}
	err = idb.DB.Delete(&person).Error
	if err != nil {
		result = gin.H{
			"result": "delete gagal",
		}
	} else {
		result = gin.H{
			"result": "Data deleted sukses",
		}
	}

	c.JSON(http.StatusOK, result)
}
