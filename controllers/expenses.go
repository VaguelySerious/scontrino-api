package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"

	"github.com/VaguelySerious/scontrino-api/models"
)

const user = "peter"


func ShowExpense(c *gin.Context) {
	var expense models.Expense
	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing ID param"})
		return
	}
	if err := models.DB.Where("id = ?", c.Param("id")).First(&expense).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": expense})
}

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

func UpdateExpense(c *gin.Context) {
	// TODO Also check if int
	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing ID param"})
		return
	}

	// Get model if exist
	var expense models.Expense
	if err := models.DB.Where("id = ?", c.Param("id")).First(&expense).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	// Validate input
	var input models.UpdateExpenseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	partialExpense := models.Expense{
		Name: expense.Name,
		Category: expense.Category,
		Cost: expense.Cost,
		Sharing: expense.Sharing,
		Date: expense.Date,
		Notes: expense.Notes,
	}
	if input.Name != "" {
		partialExpense.Name = input.Name
	}
	if input.Category != "" {
		partialExpense.Category = input.Category
	}
	if input.Cost != 0.0 {
		partialExpense.Cost = input.Cost
	}
	if input.Sharing != 0.0 {
		partialExpense.Sharing = input.Sharing
	}
	if input.Date != "" {
		partialExpense.Date = input.Date
	}
	if input.Notes != "" {
		partialExpense.Notes = input.Notes
	}
	// if input.GroupID {
	// 	partialExpense.GroupID = input.GroupID
	// }

	if errs := validator.Validate(partialExpense); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs})
		return
	}

	// if input.GroupID == 0 {
	// 	partialExpense.Sharing = 0.0
	// }

	models.DB.Model(&expense).Updates(partialExpense)

	c.JSON(http.StatusOK, gin.H{"data": expense})
}

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
