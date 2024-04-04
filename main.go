package main

import (
	"api-enigma-laundry/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	groupApi := router.Group("/api")
	{
		groupCustomers := groupApi.Group("/customers") // /api/customers
		{
			// Menampilkan semua pelanggan
			groupCustomers.GET("/", handlers.ViewDataPelanggan) // /api/customers

			// Menampilkan pelanggan bedasarkan id
			groupCustomers.GET("/:id", handlers.ViewDataPelangganById) // /api/customers/search/:id

			// Mendaftarkan pelanggan baru
			groupCustomers.POST("/", handlers.AddPelanggan) // /api/customers/add

			// Update data pelanggan
			groupCustomers.PUT("/:id", handlers.UpdatePelanggan) // /api/customers/update/:id

			// Hapus data pelanggan
			groupCustomers.DELETE("/:id", handlers.DeletePelanggan) // /api/customers/delete/:id
		}

		groupServices := groupApi.Group("/services") // /api/services
		{
			// Menampilkan semua pelanggan
			groupServices.GET("/", handlers.ViewListService) // /api/services

			// Menampilkan pelanggan bedasarkan id
			groupServices.GET("/:id", handlers.ViewListServiceById) // /api/services/search/:id

			// Mendaftarkan pelanggan baru
			groupServices.POST("/", handlers.AddNewService) // /api/services/add

			// Update data pelanggan
			groupServices.PUT("/:id", handlers.UpdateService) // /api/services/update/:id

			// Hapus data pelanggan
			groupServices.DELETE("/:id", handlers.DeleteService) // /api/services/delete/:id
		}

		groupEmployess := groupApi.Group("/employees") // /api/employees
		{
			// Menampilkan semua pegawai
			groupEmployess.GET("/", handlers.ViewDataPegawai) // /api/employees

			// Menampilkan pegawai bedasarkan id
			groupEmployess.GET("/:id", handlers.ViewDataPegawaiById) // /api/employees/search/:id

			// Mendaftarkan pegawai baru
			groupEmployess.POST("/", handlers.AddPegawai) // /api/employees/add

			// Update data pegawai
			groupEmployess.PUT("/:id", handlers.UpdatePegawai) // /api/employees/update/:id

			// Hapus data pegawai
			groupEmployess.DELETE("/:id", handlers.DeletePegawai) // /api/employees/delete/:id
		}

		groupTransactions := groupApi.Group("/transactions") // /api/transactions
		{
			// Menampilkan semua transaksi
			groupTransactions.GET("/", handlers.ViewTransaction) // /api/transactions

			// Menampilkan transaksi bedasarkan id transaksi
			groupTransactions.GET("/:id", handlers.ViewTransactionByTransactionID) // /api/transactions/search/:id

			// Menampilkan transaksi bedasarkan id pelanggan
			groupTransactions.GET("/customers/id/:id", handlers.ViewTransactionByCustomerID) // /api/transactions/search/users/id/:id

			// Menampilkan transaksi bedasarkan nama pelanggan
			groupTransactions.GET("/customers/name/:name", handlers.ViewTransactionByCustomerName) // /api/transactions/search/users/name/:name

			// Mendaftarkan transaksi baru
			groupTransactions.POST("/", handlers.AddTransaksi) // /api/transactions/add
		}
	}

	router.Run(":8080")
}
