package main

import (
	"back/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	
	// TODO : psql > mssql server
	// db setup
	db := models.SetupDatabase()
	
	// create sample data 
	sampleBoardData := models.Board{
		Content: "sample post",
		BoardSummary: models.BoardSummary{
			Author: "lsj",
			Date:   time.Now(),
			Views:  1,
		},
	}
	
	result := db.Create(&sampleBoardData)
	if result.Error != nil {
		log.Println("Insertion failed", result.Error)
	}
	
	// cors
	
	
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	
	router.Use(cors.New(config))
	
	
	router.GET("/", func(c *gin.Context) {
		c.String(200, "\n Hello World\n Can you hear me? \n")
	})
	
	router.GET("/sample", func(c *gin.Context) {
		boards, err := models.GetBoardData(db)
		
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		
		c.JSON(http.StatusOK, boards)
	})
	
	router.GET("/getPostContent", func(c *gin.Context) {
		c.String(200, "asd")
	})
	
	router.Run(":4000")
}