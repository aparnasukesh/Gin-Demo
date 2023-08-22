package main

import (
	"github.com/gin-gonic/gin"
)

type singupDataData struct {
	Username string `json:"username"`
	Password int    `json:"password"`
}

func main() {
	r := gin.Default()

	r.POST("/signup", SignUP)
	r.POST("/login", Login)
	r.GET("/homepage", HomePage)
	r.Run(":2000")
}

var SingupData singupDataData

func SignUP(ctx *gin.Context) {
	SingupData = singupDataData{}

	err := ctx.BindJSON(&SingupData)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "singup sucussful",
	})

}

func Login(ctx *gin.Context) {
	loginData := singupDataData{}

	err := ctx.BindJSON(&loginData)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err,
		})
		return
	}
	if loginData.Username == SingupData.Username && loginData.Password == SingupData.Password {
		ctx.JSON(200, gin.H{
			"message": "Login successful",
		})
	} else {
		ctx.JSON(400, gin.H{
			"message":  "Invalid username and password",
			"redirect": "http://localhost:2000/homepage",
		})
	}
}

func HomePage(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Welcome to homepage",
	})
}
