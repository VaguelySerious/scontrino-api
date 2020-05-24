package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/VaguelySerious/scontrino-api/models"
	"github.com/VaguelySerious/scontrino-api/controllers"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://wielander.me","http://localhost:8080","https://scontrino.wielander.me"},
		AllowMethods:     []string{"GET","OPTION","POST","PUT","PATCH","DELETE"},
		// AllowHeaders:     []string{"Origin"},
		// ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	// Run the server
	r.Run()
}
