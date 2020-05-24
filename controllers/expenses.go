package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"

	"scontrino-api/models"
)

const user = "peter"

func ListExpenses(c *gin.Context) {

	query := models.DB.Order("date DESC")

	// TODO Sort order

	// TODO Proper pagination

	if c.Query("name") != "" {
		query = query.Where("name LIKE ?", c.Query("name"))
	}
	if c.Query("mincost") != "" {
		query = query.Where("cost > ?", c.Query("mincost"))
	}
	if c.Query("maxcost") != "" {
		query = query.Where("cost < ?", c.Query("maxcost"))
	}
	if c.Query("category") != "" {
		query = query.Where("category == ?", c.Query("category"))
	}
	if c.Query("start") != "" {
		query = query.Where("date >= ?", c.Query("start"))
	}
	if c.Query("end") != "" {
		query = query.Where("date < ?", c.Query("end"))
	}

	if c.Query("limit") != "" {
		query = query.Limit(c.Query("limit"))
	} else {
		query = query.Limit(100)
	}

	if c.Query("offset") != "" {
		query = query.Offset(c.Query("offset"))
	} else {
		query = query.Offset(0)
	}

	var expenses []models.Expense
	query.Find(&expenses)

	c.JSON(http.StatusOK, gin.H{"data": expenses})
}

func CreateExpense(c *gin.Context) {
	// Validate input
	var input models.CreateExpenseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Date == "" {
		input.Date = time.Now().Format(time.RFC3339[0:10])
	}

	// Create expense
	expense := models.Expense{
		Name: input.Name,
		Category: input.Category,
		Cost: input.Cost,
		Sharing: input.Sharing,
		Date: input.Date,
		Notes: input.Notes,
		GroupID: input.GroupID,
	}

	if errs := validator.Validate(expense); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs})
		return
	}

	models.DB.Create(&expense)
	c.JSON(http.StatusOK, gin.H{"data": expense})
}

// func UpdateExpense(c *gin.Context) {
// 	// Get model if exist
// 	var expense models.Expense
// 	if err := models.DB.Where("id = ?", c.Param("id")).First(&expense).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	// Validate input
// 	var input models.UpdateExpenseInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	models.DB.Model(&expense).Updates(input)

// 	c.JSON(http.StatusOK, gin.H{"data": expense})
// }

func RemoveExpense(c *gin.Context) {
	// Get model if exist
	var expense models.Expense
	if err := models.DB.Where("id = ?", c.Param("id")).First(&expense).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&expense)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// func ShowExpense(c *gin.Context) {
// 	// Get model if exist
// 	var expense models.Expense
// 	if err := models.DB.Where("id = ?", c.Param("id")).First(&expense).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": expense})
// }
