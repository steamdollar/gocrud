package main

import (
	"back/board"
	"back/common"
	"back/models"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// TODO : psql > mssql server
	// db setup
	db := models.SetupDatabase()
	app := common.NewApp(db)
	
	// 각 컨트롤러들을 가져온다
	boardController := board.NewBoardController(app)
	
	controllers := []common.Controller{
		boardController,
	}
	
	// router 등록
	router := SetupRouter(controllers)
	
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
	
	
	// routers
	router.GET("/", func(c *gin.Context) {
		c.String(200, "\n Hello World\n Can you hear me? \n")
	})
	
	router.Run(":4000")
}