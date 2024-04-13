package middleware

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

		if !isValidCredentials(strings.ToLower(username), password, db) {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{"Messege": "Invalid credential. Get a valid credentials on https://get-auth-api.yusharwz.my.id/"})
			return
		}
		c.Next()
	}
}

func isValidCredentials(username, password string, db *sql.DB) bool {

	key, hashPassword, chance, err := databaseValidator(username, db)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return false
	}

	if key == "" || hashPassword == "" || chance == 0 {
		return false
	}

	chance--
	sqlStatement := "UPDATE auth SET hit_chance = $1 WHERE username = $2"
	db.Exec(sqlStatement, chance, username)

	return true
}

func databaseValidator(username string, db *sql.DB) (key, hashPassword string, chance int, err error) {

	query := "SELECT username, password, hit_chance FROM auth WHERE username = $1"
	err = db.QueryRow(query, username).Scan(&key, &hashPassword, &chance)
	if err != nil {
		return "", "", 0, err
	}

	if chance <= 0 {
		return "", "", 0, nil
	}

	return key, hashPassword, chance, nil
}
