package main

import "github.com/gin-gonic/gin"

type User struct {
	Name string
	Age  int
}

func (u *User) init() {
	u.Name = "xd"
	u.Age = 12
}

func main() {
	r := gin.Default()

	u := User{}
	u.init()

	r.GET("/", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "GET")

		c.JSON(200, gin.H{
			"data": gin.H{
				"name": u.Name,
				"age":  u.Age,
			},
		})
	})

	r.Run("127.0.0.1:5000")
}
