package middleware

import (
	"api-enigma-laundry/config"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {

	username, password, ok := c.Request.BasicAuth()

	key, value, err := GetKey(c)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"Get username authentication on": "https://get-auth-api.yusharwz.my.id", "Message": "Authentication failed",
		})
		return
	}

	if !ok || username != key || password != value {
		c.AbortWithStatusJSON(400, gin.H{
			"Get username authentication on": "https://get-auth-api.yusharwz.my.id", "Message": "Authentication failed",
		})
		return
	}
	c.Next()
}

func GetKey(c *gin.Context) (key, value string, err error) {

	db, err := config.ConnectDb()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"Message": "Server error",
		})
		return
	}

	username, _, _ := c.Request.BasicAuth()
	query := "SELECT username, password FROM auth WHERE username = $1"
	err = db.QueryRow(query, username).Scan(&key, &value)
	if err != nil {
		return
	}
	return key, value, nil
}
