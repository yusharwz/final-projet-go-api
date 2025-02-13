package main

import (
	"api-enigma-laundry/config"
	"api-enigma-laundry/handlers"
	"api-enigma-laundry/middleware"
	"api-enigma-laundry/pkg"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func init() {
	config.LoadEnvVariables()
}

func main() {

	logFile, err := os.OpenFile("/var/log/myapp/open_api.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Gagal membuat file log:", err)
	}
	defer logFile.Close()

	// Konfigurasi Logrus
	logrus.SetOutput(logFile)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	// MultiWriter untuk menyimpan log ke file dan terminal
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	db, err := config.ConnectDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	c := cron.New(cron.WithLocation(time.FixedZone("Asia/Jakarta", 7*60*60))) // Waktu zona Asia/Jakarta (UTC+7)
	_, err = c.AddFunc("0 0 * * *", func() {
		pkg.Reset(db)
	})
	if err != nil {
		fmt.Println("Error adding cron job:", err)
		return
	}
	c.Start()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(middleware.Auth(db))

	groupApi := r.Group("/api")
	{
		groupCustomers := groupApi.Group("/customers") // /api/customers
		{
			// Menampilkan semua pelanggan
			groupCustomers.GET("/", func(c *gin.Context) {
				handlers.ViewDataPelanggan(c, db)
			}) // /api/customers

			// Menampilkan pelanggan bedasarkan id
			groupCustomers.GET("/:id", func(c *gin.Context) {
				handlers.ViewDataPelangganById(c, db)
			}) // /api/customers/:id

			// Mendaftarkan pelanggan baru
			groupCustomers.POST("/", func(c *gin.Context) {
				handlers.AddPelanggan(c, db)
			}) // /api/customers

			// Update data pelanggan
			groupCustomers.PUT("/:id", func(c *gin.Context) {
				handlers.UpdatePelanggan(c, db)
			}) // /api/customers/:id

			// Hapus data pelanggan
			groupCustomers.DELETE("/:id", func(c *gin.Context) {
				handlers.DeletePelanggan(c, db)
			}) // /api/customers/:id
		}

		groupServices := groupApi.Group("/services") // /api/services
		{
			// Menampilkan semua layanan
			groupServices.GET("/", func(c *gin.Context) {
				handlers.ViewListService(c, db)
			}) // /api/services

			// Menampilkan layanan bedasarkan id
			groupServices.GET("/:id", func(c *gin.Context) {
				handlers.ViewListServiceById(c, db)
			}) // /api/services/:id

			// Mendaftarkan layanan baru
			groupServices.POST("/", func(c *gin.Context) {
				handlers.AddNewService(c, db)
			}) // /api/services

			// Update data layanan
			groupServices.PUT("/:id", func(c *gin.Context) {
				handlers.UpdateService(c, db)
			}) // /api/services/:id

			// Hapus data layanan
			groupServices.DELETE("/:id", func(c *gin.Context) {
				handlers.DeleteService(c, db)
			}) // /api/services/:id
		}

		groupEmployess := groupApi.Group("/employees") // /api/employees
		{
			// Menampilkan semua pegawai
			groupEmployess.GET("/", func(c *gin.Context) {
				handlers.ViewDataPegawai(c, db)
			}) // /api/employees

			// Menampilkan pegawai bedasarkan id
			groupEmployess.GET("/:id", func(c *gin.Context) {
				handlers.ViewDataPegawaiById(c, db)
			}) // /api/employees/:id

			// Mendaftarkan pegawai baru
			groupEmployess.POST("/", func(c *gin.Context) {
				handlers.AddPegawai(c, db)
			}) // /api/employees/add

			// Update data pegawai
			groupEmployess.PUT("/:id", func(c *gin.Context) {
				handlers.UpdatePegawai(c, db)
			}) // /api/employees/:id

			// Hapus data pegawai
			groupEmployess.DELETE("/:id", func(c *gin.Context) {
				handlers.DeletePegawai(c, db)
			}) // /api/employees/:id
		}

		groupTransactions := groupApi.Group("/transactions") // /api/transactions
		{
			// Menampilkan semua transaksi
			groupTransactions.GET("/", func(c *gin.Context) {
				handlers.ViewTransaction(c, db)
			}) // /api/transactions

			// Menampilkan transaksi bedasarkan id transaksi
			groupTransactions.GET("/:id", func(c *gin.Context) {
				handlers.ViewTransactionByTransactionID(c, db)
			}) // /api/transactions/:id

			// Menampilkan transaksi bedasarkan id pelanggan
			groupTransactions.GET("/customers/id/:id", func(c *gin.Context) {
				handlers.ViewTransactionByCustomerID(c, db)
			}) // /api/transactions/customers/id/:id

			// Menampilkan transaksi bedasarkan nama pelanggan
			groupTransactions.GET("/customers/name/:name", func(c *gin.Context) {
				handlers.ViewTransactionByCustomerName(c, db)
			}) // /api/transactions/customers/name/:name

			// Mendaftarkan transaksi baru
			groupTransactions.POST("/", func(c *gin.Context) {
				handlers.AddTransaksi(c, db)
			}) // /api/transactions
		}
	}

	r.Run()
}
