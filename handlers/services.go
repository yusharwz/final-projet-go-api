package handlers

import (
	"api-enigma-laundry/asset/entity"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ViewListService(c *gin.Context) {

	var rows *sql.Rows
	var err error

	query := "SELECT * FROM layanan"

	rows, err = db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil daftar layanan"})
		return
	}
	defer rows.Close()

	var services []entity.Layanan
	for rows.Next() {
		var service entity.Layanan
		err = rows.Scan(&service.Id, &service.NamaLayanan, &service.Satuan, &service.Harga)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mendapatkan daftar layanan"})
			return
		}
		services = append(services, service)
	}
	if len(services) > 0 {
		c.JSON(http.StatusOK, services)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Layanan dengan id tersebut tidak ditemukan"})
	}
}

func ViewListServiceById(c *gin.Context) {
	searchId := c.Param("id")

	var rows *sql.Rows
	var err error

	query := "SELECT * FROM layanan WHERE id = $1"

	rows, err = db.Query(query, searchId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil daftar layanan"})
		return
	}
	defer rows.Close()

	var services []entity.Layanan
	for rows.Next() {
		var service entity.Layanan
		err = rows.Scan(&service.Id, &service.NamaLayanan, &service.Satuan, &service.Harga)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mendapatkan daftar layanan"})
			return
		}
		services = append(services, service)
	}
	if len(services) > 0 {
		c.JSON(http.StatusOK, services)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Layanan dengan id tersebut tidak ditemukan"})
	}
}

func AddNewService(c *gin.Context) {
	var service entity.Layanan
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO layanan (nama_layanan, satuan, harga) VALUES ($1, $2, $3) RETURNING id;"

	var serviceId int

	err := db.QueryRow(query, service.NamaLayanan, service.Satuan, service.Harga).Scan(&serviceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	service.Id = serviceId
	c.JSON(http.StatusCreated, service)
}

func UpdateService(c *gin.Context) {
	id := c.Param("id")

	var service entity.Layanan
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingService entity.Layanan
	err := db.QueryRow("SELECT id, nama_layanan, satuan, harga FROM layanan WHERE id = $1", id).Scan(&existingService.Id, &existingService.NamaLayanan, &existingService.Satuan, &existingService.Harga)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Layanan dengan id tersebut tidak ditemukan"})
		return
	}

	if service.NamaLayanan == "" {
		service.NamaLayanan = existingService.NamaLayanan
	}

	if service.Satuan == "" {
		service.Satuan = existingService.Satuan
	}

	if service.Harga == 0 {
		service.Harga = existingService.Harga
	}

	query := "UPDATE layanan SET nama_layanan = $1, satuan = $2, harga = $3 WHERE id = $4;"
	_, err = db.Exec(query, service.NamaLayanan, service.Satuan, service.Harga, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	intId, _ := strconv.Atoi(id)
	service.Id = intId
	c.JSON(http.StatusOK, service)
}

func DeleteService(c *gin.Context) {
	id := c.Param("id")

	var existingService entity.Layanan
	err := db.QueryRow("SELECT id FROM layanan WHERE id = $1", id).Scan(&existingService.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Layanan dengan id tersebut tidak ditemukan"})
		return
	}

	query := "DELETE FROM layanan WHERE id = $1;"
	_, err = db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Layanan berhasil dihapus"})
}
