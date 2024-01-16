package main

import (
	"back/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	
	// db setup
	db := models.SetupDatabase()
	
	// create sample data 
	sampleBoardData := models.Board{
		Author : "lsj",
		Date : time.Now(),
		Content: "sample post",
		Views : 1,
	}
	
	result := db.Create(&sampleBoardData)
	if result.Error != nil {
		log.Println("Insertion failed", result.Error)
	}
	
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
	
	
	router.Run(":4000")
}