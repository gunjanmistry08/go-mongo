package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func checkAuth(c *gin.Context) {
	client, context, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(context)
	// cookie, err := c.Cookie("jwt")
	// if err != nil {
	// 	log.Printf("Cookie does not exist:%v\n", err)
	// 	// c.Redirect(http.StatusTemporaryRedirect, "/login")
	// 	c.JSON(http.StatusNotFound, gin.H{"err": "no cookie plj rejister UwU"})
	// 	c.Abort()
	// 	return
	// }
	auth := c.GetHeader("Auth")
	log.Printf("AuthH:%v", auth)
	cookie := strings.SplitAfter(auth, "Bearer ")[1]
	log.Printf("Cookie:%v", cookie)
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) { return []byte(secretkey), nil })
	if err != nil {
		log.Printf("token error : %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"err": "token error plj rejister UwU"})
		c.Abort()
		return
	}
	claims := token.Claims.(*jwt.StandardClaims)
	log.Printf("Issuer : %v", claims.Issuer)
	u, err := GetUserbyEmail(claims.Issuer)
	if err != nil {
		log.Printf("unauth")
		// c.Redirect(http.StatusTemporaryRedirect, "/login")
		c.JSON(http.StatusNotFound, gin.H{"err": "no cookie plj rejister UwU"})
		c.Abort()
		return

	}
	log.Println(u)
	c.Next()
}
