package main

import (
	"github.com/VaguelySerious/scontrino-api/models"
	"github.com/VaguelySerious/scontrino-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/api/v1/expenses", controllers.ListExpenses)
	// r.GET("/expenses/:id", controllers.ShowExpense)
	r.POST("/api/v1/expenses", controllers.CreateExpense)
	// r.PUT("/expenses/:id", controllers.UpdateExpense)
	// r.PATCH("/expenses/:id", controllers.UpdateExpense)
	r.DELETE("/api/v1/expenses/:id", controllers.RemoveExpense)

	// Run the server
	r.Run()
}
