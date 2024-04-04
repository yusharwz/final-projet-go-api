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
		groupCustomers.GET("/search/:id", handlers.ViewDataPelangganById)

		// Mendaftarkan pelanggan baru
		groupCustomers.POST("/add", handlers.AddPelanggan)

		// Update data pelanggan
		groupCustomers.PUT("/update/:id", handlers.UpdatePelanggan)

		// Hapus data pelanggan
		groupCustomers.DELETE("/delete/:id", handlers.DeletePelanggan)
	}

	groupServices := router.Group("/api/services")
	{
		// Menampilkan semua pelanggan
		groupServices.GET("/", handlers.ViewListService)

		// Menampilkan pelanggan bedasarkan id
		groupServices.GET("/search/:id", handlers.ViewListServiceById)

		// Mendaftarkan pelanggan baru
		groupServices.POST("/add", handlers.AddNewService)

		// Update data pelanggan
		groupServices.PUT("/update/:id", handlers.UpdateService)

		// Hapus data pelanggan
		groupServices.DELETE("/delete/:id", handlers.DeleteService)
	}

	groupEmployess := router.Group("/api/employees")
	{
		// Menampilkan semua pegawai
		groupEmployess.GET("/", handlers.ViewDataPegawai)

		// Menampilkan pegawai bedasarkan id
		groupEmployess.GET("/search/:id", handlers.ViewDataPegawaiById)

		// Mendaftarkan pegawai baru
		groupEmployess.POST("/add", handlers.AddPegawai)

		// Update data pegawai
		groupEmployess.PUT("/update/:id", handlers.UpdatePegawai)

		// Hapus data pegawai
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
