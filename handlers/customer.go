package handlers

import (
	"api-enigma-laundry/asset/entity"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ViewDataPelanggan(c *gin.Context, db *sql.DB) {

	query := "SELECT * FROM mst_pelanggan"

	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "pelanggan dengan id tersebut tidak ditemukan"})
		return
	}
	defer rows.Close()

	var customers []entity.Customers
	for rows.Next() {
		var customer entity.Customers
		err = rows.Scan(&customer.ID, &customer.Name, &customer.NoHp)
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

func ViewDataPelangganById(c *gin.Context, db *sql.DB) {
	searchId := c.Param("id")

	query := "SELECT * FROM mst_pelanggan WHERE id = $1;"
	rows, err := db.Query(query, searchId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "pelanggan dengan id tersebut tidak ditemukan"})
		return
	}
	defer rows.Close()

	var customers []entity.Customers
	for rows.Next() {
		var customer entity.Customers
		err = rows.Scan(&customer.ID, &customer.Name, &customer.NoHp)
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

func AddPelanggan(c *gin.Context, db *sql.DB) {
	var customer entity.Customers
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gagal membaca body json"})
		return
	}

	query := "INSERT INTO mst_pelanggan (nama_pelanggan, nomor_hp) VALUES ($1, $2) RETURNING id;"

	var customerId int

	err := db.QueryRow(query, customer.Name, customer.NoHp).Scan(&customerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal insert ke database"})
		return
	}
	customer.ID = customerId
	c.JSON(http.StatusCreated, customer)
}

func UpdatePelanggan(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	var customer entity.Customers
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gagal membaca body json"})
		return
	}

	var existingCustomer entity.Customers
	err := db.QueryRow("SELECT id, nama_pelanggan, nomor_hp FROM mst_pelanggan WHERE id = $1", id).Scan(&existingCustomer.ID, &existingCustomer.Name, &existingCustomer.NoHp)
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal update ke database"})
		return
	}
	intId, _ := strconv.Atoi(id)
	customer.ID = intId
	c.JSON(http.StatusOK, customer)
}

func DeletePelanggan(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	var existingCustomer entity.Customers
	err := db.QueryRow("SELECT id FROM mst_pelanggan WHERE id = $1", id).Scan(&existingCustomer.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pelanggan dengan id tersebut tidak ditemukan"})
		return
	}

	query := "DELETE FROM mst_pelanggan WHERE id = $1;"
	_, err = db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menghapus Pelanggan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pelanggan berhasil dihapus"})
}
