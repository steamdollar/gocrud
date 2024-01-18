package board

import (
	"back/common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BoardController struct {
	App *common.App
}

// 미들웨어 역할을 할 메소드들을 가진 BoardContoller 반환
func NewBoardController(app *common.App) *BoardController {
	return &BoardController{App: app}
}

// RegisterRoutes : 미들웨어, url을 묶은 라우터들을 모은다.
// 이렇게 모아서 main 함수에 넣어주면 됨
func (bc *BoardController) RegisterRoutes (router *gin.Engine) {
	boardRoutes := router.Group("/boards") 
	{
		boardRoutes.GET("/sample", bc.GetSample)
        	boardRoutes.GET("/view/:idx", bc.GetPostContent)
	}
}

func (bc *BoardController) GetSample(c *gin.Context) {
	boards, err := GetBoardData(bc.App.DB)
		
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
		
	c.JSON(http.StatusOK, boards)
}

func (bc *BoardController) GetPostContent(c *gin.Context) {
	idx, err := strconv.Atoi(c.Param("idx"))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index format"})
		return
	}

	result, err := GetPostContent(bc.App.DB, idx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "failed to get content"})
		return
	}
    
	c.JSON(http.StatusOK, gin.H{"content": result})
}