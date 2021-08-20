package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetResturant(c *gin.Context) {
	restaurants, err := Get()
	if err != nil {
		log.Printf("Could not get resturants: %v", err)
		c.JSON(http.StatusBadGateway, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"res": restaurants})
	return
}

func HandleCreateResturant(c *gin.Context) {
	var rest resturant
	if err := c.ShouldBindJSON(&rest); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	id, err := Create(&rest)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})

}

func HandleGetIdResturant(c *gin.Context) {
	id := c.Param("id")
	arest, err := GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"rest": arest})
	return
}

func HandleUpdateResturant(c *gin.Context) {
	var rest resturant
	if err := c.ShouldBindJSON(&rest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	updatedrest, err := Update(&rest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"rest": updatedrest})
	return
}

func handleDeleteResturantbyId(c *gin.Context) {
	id := c.Param("id")
	count, err := Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deletecount": count})

}
