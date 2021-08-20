package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func handleRegister(c *gin.Context) {
	var u user
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	uid, err := CreateUser(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
	}
	c.JSON(http.StatusCreated, gin.H{"id": uid})
}

func handleLogin(c *gin.Context) {
	var u user
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    u.Email,
		ExpiresAt: time.Now().Add(time.Minute * 30).Unix()},
	)
	token, err := claims.SignedString([]byte(secretkey))

	if err != nil {
		log.Printf("_%v", err.Error())
		c.String(http.StatusBadRequest, err.Error())
	}
	cookie, err := c.Cookie("jwt")
	if err != nil {
		cookie = "Not_Set"
		c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)
	}
	log.Println(cookie)
}
