package main

import (
	"api-enigma-laundry/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	groupCustomers := router.Group("/api/customers")
	{
		// Menampilkan semua pelanggan
		groupCustomers.GET("/", handlers.ViewDataPelanggan)

		// Menampilkan pelanggan bedasarkan id
		groupCustomers.GET("/:id", handlers.ViewDataPelangganById)

		// Mendaftarkan pelanggan baru
		groupCustomers.POST("/add", handlers.AddPelanggan)

		// Update data pelanggan
		groupCustomers.PUT("/update/:id", handlers.UpdatePelanggan)

		// Hapus data pelanggan
		groupCustomers.DELETE("/delete/:id", handlers.DeletePelanggan)
	}

	groupEmployess := router.Group("/api/employees")
	{
		// Menampilkan semua pelanggan
		groupEmployess.GET("/", handlers.ViewDataPegawai)

		// Menampilkan pelanggan bedasarkan id
		groupEmployess.GET("/:id", handlers.ViewDataPegawaiById)

		// Mendaftarkan pelanggan baru
		groupEmployess.POST("/add", handlers.AddPegawai)

		// Update data pelanggan
		groupEmployess.PUT("/update/:id", handlers.UpdatePegawai)

		// Hapus data pelanggan
		groupEmployess.DELETE("/delete/:id", handlers.DeletePegawai)
	}

	groupTransactions := router.Group("/api/transactions")
	{
		// Menampilkan semua transaksi
		groupTransactions.GET("/", handlers.ViewTransaction)

		// Menampilkan transaksi bedasarkan id transaksi
		groupTransactions.GET("/search/id/:id", handlers.ViewTransactionByTransactionID)

		// Menampilkan transaksi bedasarkan id pelanggan
		groupTransactions.GET("/search/users/id/:id", handlers.ViewTransactionByCustomerID)

		// Menampilkan transaksi bedasarkan nama pelanggan
		groupTransactions.GET("/search/users/name/:name", handlers.ViewTransactionByCustomerName)

		// Mendaftarkan transaksi baru
		groupTransactions.POST("/add", handlers.AddTransaksi)
	}

	router.Run(":8080")
}
