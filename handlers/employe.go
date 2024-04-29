package handlers

import (
	"api-enigma-laundry/asset/entity"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ViewDataPegawai(c *gin.Context, db *sql.DB) {

	query := "SELECT * FROM mst_pegawai"

	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "pegawai dengan id tersebut tidak ditemukan"})
		return
	}
	defer rows.Close()

	var employees []entity.Pegawai
	for rows.Next() {
		var employe entity.Pegawai
		err = rows.Scan(&employe.ID, &employe.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mendapatkan daftar pegawai"})
			return
		}
		employees = append(employees, employe)
	}
	if len(employees) > 0 {
		c.JSON(http.StatusOK, employees)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pegawai dengan id tersebut tidak ditemukan"})
	}
}

func ViewDataPegawaiById(c *gin.Context, db *sql.DB) {
	searchId := c.Param("id")

	query := "SELECT * FROM mst_pegawai WHERE id = $1;"
	rows, err := db.Query(query, searchId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "pegawai dengan id tersebut tidak ditemukan"})
		return
	}
	defer rows.Close()

	var employees []entity.Pegawai
	for rows.Next() {
		var employe entity.Pegawai
		err = rows.Scan(&employe.ID, &employe.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mendapatkan daftar pegawai"})
			return
		}
		employees = append(employees, employe)
	}
	if len(employees) > 0 {
		c.JSON(http.StatusOK, employees)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pegawai dengan id tersebut tidak ditemukan"})
	}
}

func AddPegawai(c *gin.Context, db *sql.DB) {
	var employe entity.Pegawai
	if err := c.ShouldBindJSON(&employe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gagal membaca body json"})
		return
	}

	query := "INSERT INTO mst_pegawai (nama_pegawai) VALUES ($1) RETURNING id;"

	var employeId int

	err := db.QueryRow(query, employe.Name).Scan(&employeId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal insert ke database"})
		return
	}
	employe.ID = employeId
	c.JSON(http.StatusCreated, employe)
}

func UpdatePegawai(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	var employe entity.Pegawai
	if err := c.ShouldBindJSON(&employe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gagal membaca body json"})
		return
	}

	var existingEmploye entity.Pegawai
	err := db.QueryRow("SELECT id, nama_pegawai FROM mst_pegawai WHERE id = $1", id).Scan(&existingEmploye.ID, &existingEmploye.Name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pegawai dengan id tersebut tidak ditemukan"})
		return
	}

	if employe.Name == "" {
		employe.Name = existingEmploye.Name
	}

	query := "UPDATE mst_pegawai SET nama_pegawai = $1 WHERE id = $2;"
	_, err = db.Exec(query, employe.Name, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal update data ke database"})
		return
	}
	intId, _ := strconv.Atoi(id)
	employe.ID = intId
	c.JSON(http.StatusOK, employe)
}

func DeletePegawai(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	var existingEmploye entity.Pegawai
	err := db.QueryRow("SELECT id FROM mst_pegawai WHERE id = $1", id).Scan(&existingEmploye.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pegawai dengan id tersebut tidak ditemukan"})
		return
	}

	query := "DELETE FROM mst_pegawai WHERE id = $1;"
	_, err = db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menghapus pegawai"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pegawai berhasil dihapus"})
}
