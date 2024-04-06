package middleware

import (
	"api-enigma-laundry/config"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {

	username, password, _ := c.Request.BasicAuth()

	if username == "" && password == "" {
		c.AbortWithStatusJSON(400, gin.H{
			"Message": "Authentication required. Get username and password authentication on https://get-auth-api.yusharwz.my.id",
		})
		return
	}

	_, _, err := databaseValidator(c)
	if err != nil {
		return
	}
	c.Next()
}

func databaseValidator(c *gin.Context) (key, value string, err error) {

	var chance int

	db, err := config.ConnectDb()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"Message": "Server error",
		})
		return
	}

	username, password, _ := c.Request.BasicAuth()
	query := "SELECT username, password, hit_chance FROM auth WHERE username = $1 AND password = $2"
	db.QueryRow(query, username, password).Scan(&key, &value, &chance)

	if key == "" || value == "" {
		c.AbortWithStatusJSON(400, gin.H{
			"Message": "Invalid username or password. Get a valid username and password authentication on https://get-auth-api.yusharwz.my.id",
		})
		return
	}

	if chance <= 0 {
		c.AbortWithStatusJSON(400, gin.H{
			"Message": "Your hit chance is up",
		})
		return
	} else {
		chance -= 1
		sqlStatement := "UPDATE auth SET hit_chance = $1  WHERE username = $2"
		_, err := db.Query(sqlStatement, chance, username)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"Message": "Server error"})
		}
	}
	return key, value, nil
}
