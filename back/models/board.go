package models

import (
	"fmt"
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

// GetBoardData : board 테이블에서 데이터 검색
func GetBoardData(db *gorm.DB)([]Board, error) {
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
func GetPostContent(db *gorm.DB, idx int) () {
	var boards []Board
	content := db.Select("Content").Where("Idx = ?", idx).Find(&boards)
	
	fmt.Println(content)
}