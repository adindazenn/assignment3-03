package api

import (
    "github.com/gin-gonic/gin"
    "fmt"
    "github.com/adindazenn/assignment3-03/assignment3/model"
    "github.com/adindazenn/assignment3-03/assignment3/database"
    "net/http"
)

func UpdateData(c *gin.Context) {
    db, err := database.InitDB()
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}
	var dataToUpdate model.Data
	if err := c.ShouldBindJSON(&dataToUpdate); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("Error:", err.Error())
		return
	}

	// pembaruan data ke dalam database
	if err := db.Model(&model.Data{}).Where("id = ?", 1).Updates(&dataToUpdate).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update data in database"})
		fmt.Println("Error updating data in database:", err)
		return
	}

    c.JSON(http.StatusOK, gin.H{"message": "Data updated successfully"})
}
