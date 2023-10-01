package main

import (
    "github.com/adindazenn/assignment3-03/assignment3/database"
    "github.com/adindazenn/assignment3-03/assignment3/update"
    "github.com/adindazenn/assignment3-03/assignment3/api"
    "github.com/gin-gonic/gin"
)

func main() {
    // Inisialisasi database
    database.InitDB()

    // Mulai service untuk mengupdate data setiap 15 detik
    go update.UpdateDataPeriodically()

    r := gin.Default()

	// Endpoint untuk menerima permintaan PUT untuk memperbarui data
	r.PUT("/api/update", api.UpdateData)
    r.Run(":8080")
}
