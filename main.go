package main

import (
	"api-enigma-laundry/config"
	"api-enigma-laundry/handlers"
	"api-enigma-laundry/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
}

func main() {

	db, err := config.ConnectDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

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
