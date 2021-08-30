package main

import (
	// "time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	DATABASE   = "test"
	COLLECTION = "restaurants"
	USER       = "user"
	secretkey  = "secret"
)

type resturant struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	Cuisine string             `bson:"cuisine"`
	Borough string             `bson:"borough"`
}

type user struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}

func main() {
	r := gin.Default()
	// r.Use(checkAuth)
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "*"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))
	r.Use(cors.Default())
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "*")
		c.Next()
	})

	auth := r.Group("/auth")
	auth.POST("/register", handleRegister)
	auth.POST("/login", handleLogin)
	// checkAuth
	res := r.Group("/res", checkAuth)
	res.GET("", HandleGetResturant)
	res.GET("/:id", HandleGetIdResturant)
	res.DELETE("/:id", handleDeleteResturantbyId)
	res.POST("", HandleCreateResturant)
	res.PUT("", HandleUpdateResturant)
	r.Run(":8000")

}
