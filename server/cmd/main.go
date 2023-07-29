package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mrefawaladam/server/config"
	"github.com/mrefawaladam/server/routes"
)

func main() {
	db, err := config.InitDB() // Membuat koneksi database sesungguhnya
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	config.MigrateDB() // Menjalankan migration

	r := gin.Default()
	routes.SetupRoutes(r, db)
	r.Run(":8080")
}
