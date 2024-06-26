package handlers

import (
	"api-enigma-laundry/asset/entity"
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddTransaksi(c *gin.Context, db *sql.DB) {

	tx, err := db.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer tx.Rollback()

	var transaksi entity.TransaksiAndDetail
	if err := c.ShouldBindJSON(&transaksi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gagal sinkron data transaksi"})
		return
	}

	query := "INSERT INTO transaksi (id_pelanggan, id_pegawai, tanggal_keluar, status_pembayaran) VALUES ($1, $2, $3, $4) RETURNING id, tanggal_masuk"

	err = tx.QueryRow(query, transaksi.IDPelanggan, transaksi.IDPegawai, transaksi.TanggalKeluar, transaksi.StatusPembayaran).Scan(&transaksi.ID, &transaksi.TanggalMasuk)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal insert ke database"})
		return
	}

	for i, detail := range transaksi.DetailTransaksi {
		err = tx.QueryRow("INSERT INTO detail_transaksi (id_transaksi, id_layanan, quantity) VALUES ($1, $2, $3) RETURNING id, id_transaksi", transaksi.ID, detail.IDLayanan, detail.Quantity).Scan(&transaksi.DetailTransaksi[i].ID, &transaksi.DetailTransaksi[i].IDTransaksi)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal insert ke database"})
			return
		}
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal commit, perubahan telah di rollback"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"transaksi": transaksi})
}

func ViewTransaction(c *gin.Context, db *sql.DB) {

	query := `
	SELECT 
		mst_pelanggan.nama_pelanggan, 
		layanan.nama_layanan, 
		detail_transaksi.quantity, 
		mst_pegawai.nama_pegawai, 
		transaksi.tanggal_masuk, 
		layanan.harga * detail_transaksi.quantity AS total_harga 
	FROM 
		transaksi 
	JOIN 
		mst_pelanggan ON transaksi.id_pelanggan = mst_pelanggan.id 
	JOIN 
		detail_transaksi ON transaksi.id = detail_transaksi.id_transaksi 
	JOIN 
		layanan ON detail_transaksi.id_layanan = layanan.id 
	JOIN 
		mst_pegawai ON transaksi.id_pegawai = mst_pegawai.id`

	rows, err := db.Query(query)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "transaksi dengan id tersebut tidak ditemukan"})
		return
	}
	defer rows.Close()

	var transactionDetails []entity.TransactionDetail
	for rows.Next() {
		var detail entity.TransactionDetail
		err = rows.Scan(&detail.NamaPelanggan, &detail.NamaLayanan, &detail.Quantity, &detail.NamaPegawai, &detail.TanggalMasuk, &detail.Harga)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mendapatkan daftar transaksi"})
			return
		}
		transactionDetails = append(transactionDetails, detail)
	}

	if len(transactionDetails) > 0 {
		c.JSON(http.StatusOK, gin.H{"All Transaksi": transactionDetails})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaksi dengan id tersebut tidak ditemukan"})
	}
}

func ViewTransactionByCustomerName(c *gin.Context, db *sql.DB) {
	customerName := strings.ToLower(c.Param("name"))

	query := `
		SELECT 
			mst_pelanggan.nama_pelanggan, 
			layanan.nama_layanan, 
			detail_transaksi.quantity, 
			mst_pegawai.nama_pegawai, 
			transaksi.tanggal_masuk, 
			layanan.harga * detail_transaksi.quantity AS total_harga 
		FROM 
			transaksi 
		JOIN 
			mst_pelanggan ON transaksi.id_pelanggan = mst_pelanggan.id 
		JOIN 
			detail_transaksi ON transaksi.id = detail_transaksi.id_transaksi 
		JOIN 
			layanan ON detail_transaksi.id_layanan = layanan.id 
		JOIN 
			mst_pegawai ON transaksi.id_pegawai = mst_pegawai.id 
		WHERE 
			LOWER(mst_pelanggan.nama_pelanggan) LIKE '%' || $1 || '%';
	`

	rows, err := db.Query(query, customerName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mendapatkan daftar transaksi"})
		return
	}
	defer rows.Close()

	var totalPembayaran int
	var transactionDetails []entity.TransactionDetail
	for rows.Next() {
		var detail entity.TransactionDetail
		err = rows.Scan(&detail.NamaPelanggan, &detail.NamaLayanan, &detail.Quantity, &detail.NamaPegawai, &detail.TanggalMasuk, &detail.Harga)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mendapatkan daftar transaksi"})
			return
		}
		transactionDetails = append(transactionDetails, detail)
		totalPembayaran += detail.Harga
	}

	if len(transactionDetails) > 0 {
		c.JSON(http.StatusOK, gin.H{"Detail Transaksi": transactionDetails, "Total Pembayaran": totalPembayaran})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaksi untuk pelanggan dengan nama tersebut tidak ditemukan"})
	}
}

func ViewTransactionByCustomerID(c *gin.Context, db *sql.DB) {
	customerID := c.Param("id")

	query := `
		SELECT 
			mst_pelanggan.nama_pelanggan, 
			layanan.nama_layanan, 
			detail_transaksi.quantity, 
			mst_pegawai.nama_pegawai, 
			transaksi.tanggal_masuk, 
			layanan.harga * detail_transaksi.quantity AS total_harga 
		FROM 
			transaksi 
		JOIN 
			mst_pelanggan ON transaksi.id_pelanggan = mst_pelanggan.id 
		JOIN 
			detail_transaksi ON transaksi.id = detail_transaksi.id_transaksi 
		JOIN 
			layanan ON detail_transaksi.id_layanan = layanan.id 
		JOIN 
			mst_pegawai ON transaksi.id_pegawai = mst_pegawai.id 
		WHERE 
			mst_pelanggan.id = $1
	`

	rows, err := db.Query(query, customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mendapatkan daftar transaksi"})
		return
	}
	defer rows.Close()

	var totalPembayaran int
	var transactionDetails []entity.TransactionDetail
	for rows.Next() {
		var detail entity.TransactionDetail
		err = rows.Scan(&detail.NamaPelanggan, &detail.NamaLayanan, &detail.Quantity, &detail.NamaPegawai, &detail.TanggalMasuk, &detail.Harga)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mendapatkan daftar transaksi"})
			return
		}
		transactionDetails = append(transactionDetails, detail)
		totalPembayaran += detail.Harga
	}

	if len(transactionDetails) > 0 {
		c.JSON(http.StatusOK, gin.H{"Detail Transaksi": transactionDetails, "Total Pembayaran": totalPembayaran})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaksi untuk pelanggan dengan ID tersebut tidak ditemukan"})
	}
}

func ViewTransactionByTransactionID(c *gin.Context, db *sql.DB) {
	transactionID := c.Param("id")

	query := `
      SELECT 
            mst_pelanggan.nama_pelanggan, 
            layanan.nama_layanan, 
            detail_transaksi.quantity, 
            mst_pegawai.nama_pegawai, 
            transaksi.tanggal_masuk, 
            layanan.harga * detail_transaksi.quantity AS total_harga 
      FROM 
            transaksi 
      JOIN 
            mst_pelanggan ON transaksi.id_pelanggan = mst_pelanggan.id 
      JOIN 
            detail_transaksi ON transaksi.id = detail_transaksi.id_transaksi 
      JOIN 
            layanan ON detail_transaksi.id_layanan = layanan.id 
      JOIN 
            mst_pegawai ON transaksi.id_pegawai = mst_pegawai.id 
      WHERE 
            transaksi.id = $1
   `

	rows, err := db.Query(query, transactionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mendapatkan daftar transaksi"})
		return
	}
	defer rows.Close()

	var totalPembayaran int
	var transactionDetails []entity.TransactionDetail
	for rows.Next() {
		var detail entity.TransactionDetail
		err = rows.Scan(&detail.NamaPelanggan, &detail.NamaLayanan, &detail.Quantity, &detail.NamaPegawai, &detail.TanggalMasuk, &detail.Harga)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mendapatkan daftar transaksi"})
			return
		}
		transactionDetails = append(transactionDetails, detail)
		totalPembayaran += detail.Harga
	}

	if len(transactionDetails) > 0 {
		c.JSON(http.StatusOK, gin.H{"Detail Transaksi": transactionDetails, "Total Pembayaran": totalPembayaran})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaksi dengan ID tersebut tidak ditemukan"})
	}
}
