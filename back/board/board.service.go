package board

import (
	"time"

	"gorm.io/gorm"
)

// db의 칼럼으로 사용하는 struct는 전부 그 속성들이 대문자로 시작해야 한다.
// 그래야 공개가 되서 다른 곳에서 사용 가능
type BoardSummary struct {
	Idx     int `gorm:"primaryKey"`
	Author  string
	Date    time.Time
	Views   int
}

type Board struct {
	BoardSummary `gorm:"embedded"`
	Content string
}

type BoardService struct {
	DB *gorm.DB
}

func NewBoardService(db *gorm.DB) *BoardService {
	return &BoardService{DB : db}
}


// GetBoardData : board 테이블에서 데이터 검색
func (bs *BoardService) GetBoardData(db *gorm.DB) ([]Board, error) {
	var boards []Board
	
	// gorm은 어떻게 여기서 상호작용할 테이블을 알까?
	// gorm은 구조체와 테이블(db)간 매핑을 생성하고
	// 디폴트로 구조체의 이름의 복수형인 테이블 이름을 찾는다.
	// e.g. Board 구조체는 boards 테이블과 매핑된다.
	//
	// slice의 타입 []Board를 보고 Gorm은 어떤 테이블과 상호작용할 지 알 수 있다.
	// 그러고 나면 모든 데이터를 가져와 슬라이스에 추가한다.
	result := db.Select("Idx", "Author", "Date", "Views").
	// Where("Idx BETWEEN ? AND ?", 3, 7).
	Find(&boards)
	
	if result.Error != nil {
		return []Board{}, result.Error
	}
	
	return boards, nil
}

// GetPostContent : 해당 idx의 데이터 셋의 내용을 리턴
func (bs *BoardService) GetPostContent(db *gorm.DB, idx int) (Board, error) {
	var board Board
	
	result := db.
	Table("boards"). // specify table
	// Select("*"). // specify column
	Where("Idx = ?", idx). // 
	First(&board) //scan result
	
	// how to see entire obj
	//fmt.Printf("%+v\n", result.Config)
	// spew.Dump(result)
	if result.Error != nil {
		return board, result.Error
	}
	
	newViews := board.Views + 1
	updateResult := db.Table("boards").Where("Idx = ?", idx).Update("views", newViews)
	if updateResult.Error != nil {
		return board, updateResult.Error
	}

	return board, nil
}

func (bs *BoardService) CreatePost(db *gorm.DB, postReq CreatePostDTO) (error) {
	newPost := Board{
		BoardSummary : BoardSummary{
			Author: postReq.Author,
			Views : 0,
		},
		Content : postReq.Content,
	}
	
	result := db.Create(&newPost)
	if result.Error != nil {
		return  result.Error
	}
	
	return nil 
}

func (bs *BoardService) DeletePost(db *gorm.DB, delReq DeleteRequestDTO) (error) {
	result := db.
	Where("idx = ? And author = ?", delReq.Idx, delReq.Author).
	Delete(&Board{})
	
	if result.Error != nil {
		return result.Error
	}
	
	return nil	
}