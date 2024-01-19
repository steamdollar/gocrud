package board

import (
	"back/common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BoardController struct {
	App *common.App
	Service *BoardService
}

type CreatePostDTO struct {
	Author string `json:"author"`
	Content string `json:"content"`
}

type DeleteRequestDTO struct {
	Author string `json:"author"`
	Idx int `json:"idx"`
}

// 미들웨어 역할을 할 메소드들을 가진 BoardContoller 반환
func NewBoardController(app *common.App) *BoardController {
	return &BoardController{App: app, Service: NewBoardService(app.DB)}
}

// RegisterRoutes : 미들웨어, url을 묶은 라우터들을 모은다.
// 이렇게 모아서 main 함수에 넣어주면 됨
func (bc *BoardController) RegisterRoutes (router *gin.Engine) {
	boardRoutes := router.Group("/boards") 
	{
		boardRoutes.GET("/sample", bc.GetSample)
        	boardRoutes.GET("/view/:idx", bc.GetPostContent)
		boardRoutes.POST("/createPost", bc.CreatePost)
		boardRoutes.DELETE("/deletePost", bc.DeletePost)
	}
}

// Controllers

// GetSample : sample post 반환
func (bc *BoardController) GetSample(c *gin.Context) {
	boards, err := bc.Service.GetBoardData(bc.App.DB)
		
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
		
	c.JSON(http.StatusOK, boards)
}

// GetPostContent : idx (req.query) 에 해당하는 post.content 반환
func (bc *BoardController) GetPostContent(c *gin.Context) {
	idx, err := strconv.Atoi(c.Param("idx"))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index format"})
		return
	}

	result, err := bc.Service.GetPostContent(bc.App.DB, idx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "failed to get content"})
		return
	}
    
	c.JSON(http.StatusOK, gin.H{"post": result})
}

// CreatePost : 새로운 데이터셋 생성, 저장
func (bc *BoardController) CreatePost(c *gin.Context) {
	// POST로 받은 내용을 고려해서
	// db에 새로운 dataset을 생성
	var postReq CreatePostDTO
	if err := c.ShouldBindJSON(&postReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return 
	}
	
	err := bc.Service.CreatePost(bc.App.DB, postReq)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Post created successfully"})
	
}

func (bc *BoardController) DeletePost(c *gin.Context) {
	// author ,idx 필요
	var deleteRequestDTO DeleteRequestDTO
	if err := c.ShouldBindJSON(&deleteRequestDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}
	
	err := bc.Service.DeletePost(bc.App.DB, deleteRequestDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        	return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}