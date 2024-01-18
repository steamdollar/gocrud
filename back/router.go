package main

import (
	"back/common"

	"github.com/gin-gonic/gin"
)

func SetupRouter(controllers []common.Controller) *gin.Engine {
	router := gin.Default()
	
	for _, controller := range controllers {
		controller.RegisterRoutes(router)
	}
	
	return router
}

// func (app *App) GetSample(c *gin.Context) {
// 	boards, err := models.GetBoardData(app.DB)
		
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
		
// 	c.JSON(http.StatusOK, boards)
// }
    
// func (app *App) GetPostContent(c *gin.Context) {
// 	var post models.Board
// 	// idx := 
// 	intIdx, err := strconv.Atoi(c.Param("idx"))
	
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index format"})
// 		return
// 	}
     
// 	result := app.DB.Where("idx = ?", intIdx).First(&post)
// 	if result.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error" : "invalid idx"})
// 		return
// 	}
    
// 	c.JSON(http.StatusOK, gin.H{"content": post.Content})
// }