package main

import (
	"github.com/gin-gonic/gin"

	"github.com/VaguelySerious/scontrino-api/models"
	"github.com/VaguelySerious/scontrino-api/controllers"
)

const (
	corssites = "https://wielander.me,http://localhost:8080,https://scontrino.wielander.me"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	r.Use(CORSMiddleware())

	// Routes
	r.GET("/api/v1/expenses", controllers.ListExpenses)
	r.GET("/api/v1/expenses/:id", controllers.ShowExpense)
	r.POST("/api/v1/expenses", controllers.CreateExpense)
	// TODO Make PUT overwrite instead of doing the same as PATCH
	r.PUT("/api/v1/expenses/:id", controllers.UpdateExpense)
	r.PATCH("/api/v1/expenses/:id", controllers.UpdateExpense)
	r.DELETE("/api/v1/expenses/:id", controllers.RemoveExpense)


	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", corssites)
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}
