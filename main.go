package main

import (
	"api-enigma-laundry/config"
	"api-enigma-laundry/handlers"
	"api-enigma-laundry/middleware"
	"api-enigma-laundry/pkg"
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func init() {
	config.LoadEnvVariables()
}

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Setup CORS
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: false,
		AllowOrigins:    []string{"*"},
		AllowMethods:    []string{"GET", "POST", "PUT", "OPTIONS", "DELETE"},
		AllowHeaders: []string{
			"Origin", "Content-Type",
			"Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           120 * time.Second,
	}))

	// Get current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Error().Msg("Failed to get current working directory")
	}

	// Ensure log directory exists
	logDir := currentDir + "/log"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Error().Msg("Failed to create log directory")
	}

	// Konfigurasi Zerolog
	zerolog.TimeFieldFormat = "02-01-2006 15:04:05"
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	// Konfigurasi logger untuk menulis ke file
	logFile := &lumberjack.Logger{
		Filename:   logDir + "/app.log",
		MaxSize:    10, // dalam MB
		MaxBackups: 30,
		MaxAge:     7, // dalam hari
		Compress:   true,
	}
	defer logFile.Close()

	gin.DefaultWriter = logFile

	r.Use(gin.Logger())

	// Set output logger ke stdout dan Loki
	log.Logger = log.Output(zerolog.MultiLevelWriter(os.Stdout, logFile))

	r.Use(logger.SetLogger(
		logger.WithLogger(func(_ *gin.Context, l zerolog.Logger) zerolog.Logger {
			return l.Output(os.Stdout).With().Caller().Logger()
		}),
	))

	r.Use(gin.Recovery())

	db, err := config.ConnectDb()
	if err != nil {
		log.Error().Msg("Database connection failed")
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
