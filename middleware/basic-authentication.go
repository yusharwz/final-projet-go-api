package middleware

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{"Messege": "Need authentication. Get your credentials on https://get-auth-api.yusharwz.my.id/"})
			return
		}

		if !isValidCredentials(username, password, db) {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{"Messege": "Invalid credential. Get a valid credentials on https://get-auth-api.yusharwz.my.id/"})
			return
		}
		c.Next()
	}
}

func isValidCredentials(username, password string, db *sql.DB) bool {

	key, value, err := databaseValidator(username, password, db)
	if err != nil {
		return false
	}

	if key == "" || value == "" {
		return false
	}
	return true
}

func databaseValidator(username, password string, db *sql.DB) (key, value string, err error) {
	var chance int

	query := "SELECT username, password, hit_chance FROM auth WHERE username = $1 AND password = $2"
	err = db.QueryRow(query, username, password).Scan(&key, &value, &chance)
	if err != nil {
		return "", "", err
	}

	if chance <= 0 {
		return "", "", nil
	}

	chance--
	sqlStatement := "UPDATE auth SET hit_chance = $1 WHERE username = $2"
	_, err = db.Exec(sqlStatement, chance, username)
	if err != nil {
		return "", "", err
	}

	return key, value, nil
}
