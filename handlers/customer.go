package handlers

import (
	"api-enigma-laundry/asset/entity"
	"api-enigma-laundry/config"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var db = config.ConnectDb()

func ViewDataPelanggan(c *gin.Context) {
	searchId := c.Query("id")

	var rows *sql.Rows
	var err error

	query := "SELECT * FROM mst_pelanggan"

	if searchId != "" {
		query += " WHERE id = $1;"
		rows, err = db.Query(query, searchId)
	} else {
		rows, err = db.Query(query)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "pelanggan dengan id tersebut tidak ditemukan"})
		return
	}
	defer rows.Close()

	var customers []entity.Customers
	for rows.Next() {
		var customer entity.Customers
		err = rows.Scan(&customer.Id, &customer.Name, &customer.NoHp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mendapatkan daftar pelanggan"})
			return
		}
		customers = append(customers, customer)
	}
	if len(customers) > 0 {
		c.JSON(http.StatusOK, customers)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pelanggan dengan id tersebut tidak ditemukan"})
	}
}

func ViewDataPelangganById(c *gin.Context) {
	searchId := c.Param("id")

	var rows *sql.Rows
	var err error

	query := "SELECT * FROM mst_pelanggan WHERE id = $1;"
	rows, err = db.Query(query, searchId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "pelanggan dengan id tersebut tidak ditemukan"})
		return
	}
	defer rows.Close()

	var customers []entity.Customers
	for rows.Next() {
		var customer entity.Customers
		err = rows.Scan(&customer.Id, &customer.Name, &customer.NoHp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mendapatkan daftar pelanggan"})
			return
		}
		customers = append(customers, customer)
	}
	if len(customers) > 0 {
		c.JSON(http.StatusOK, customers)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pelanggan dengan id tersebut tidak ditemukan"})
	}
}

func AddPelanggan(c *gin.Context) {
	var customer entity.Customers
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO mst_pelanggan (nama_pelanggan, nomor_hp) VALUES ($1, $2) RETURNING id;"

	var customerId int

	err := db.QueryRow(query, customer.Name, customer.NoHp).Scan(&customerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	customer.Id = customerId
	c.JSON(http.StatusCreated, customer)
}

func UpdatePelanggan(c *gin.Context) {
	id := c.Param("id")

	var customer entity.Customers
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingCustomer entity.Customers
	err := db.QueryRow("SELECT id, nama_pelanggan, nomor_hp FROM mst_pelanggan WHERE id = $1", id).Scan(&existingCustomer.Id, &existingCustomer.Name, &existingCustomer.NoHp)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pelanggan dengan id tersebut tidak ditemukan"})
		return
	}

	if customer.Name == "" {
		customer.Name = existingCustomer.Name
	}

	if customer.NoHp == "" {
		customer.NoHp = existingCustomer.NoHp
	}

	query := "UPDATE mst_pelanggan SET nama_pelanggan = $1, nomor_hp = $2 WHERE id = $3;"
	_, err = db.Exec(query, customer.Name, customer.NoHp, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	intId, _ := strconv.Atoi(id)
	customer.Id = intId
	c.JSON(http.StatusOK, customer)
}

func DeletePelanggan(c *gin.Context) {
	id := c.Param("id")

	var existingCustomer entity.Customers
	err := db.QueryRow("SELECT id FROM mst_pelanggan WHERE id = $1", id).Scan(&existingCustomer.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pelanggan dengan id tersebut tidak ditemukan"})
		return
	}

	query := "DELETE FROM mst_pelanggan WHERE id = $1;"
	_, err = db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pelanggan berhasil dihapus"})
}
