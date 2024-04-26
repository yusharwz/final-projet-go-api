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
			c.JSON(http.StatusUnauthorized, gin.H{"Messege": "Need authentication. Get your credentials on https://get-credential-api.yusharwz.my.id/"})
			return
		}

		if !isValidCredentials(strings.ToLower(username), password, db) {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{"Messege": "Invalid credential. Get a valid credentials on https://get-credential-api.yusharwz.my.id/"})
			return
		}
		c.Next()
	}
}

func isValidCredentials(username, password string, db *sql.DB) bool {

	key, hashPassword, err := databaseValidator(username, db)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return false
	}

	if key == "" || hashPassword == "" {
		return false
	}
	return true
}

func databaseValidator(username string, db *sql.DB) (key, hashPassword string, err error) {

	var chance int
	var status string

	query := "SELECT username, password, hit_chance, status FROM auth WHERE username = $1"
	err = db.QueryRow(query, username).Scan(&key, &hashPassword, &chance, &status)
	if err != nil {
		return "", "", err
	}

	if status == "free" {
		if chance <= 0 {
			return "", "", nil
		}
	}

	chance--
	sqlStatement := "UPDATE auth SET hit_chance = $1 WHERE username = $2 AND status = $3"
	db.Exec(sqlStatement, chance, username, "free")

	return key, hashPassword, nil
}
