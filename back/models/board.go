package models

import (
	"time"

	"gorm.io/gorm"
)

// db의 칼럼으로 사용하는 struct는 전부 그 속성들이 대문자로 시작해야 한다.
// 그래야 공개가 되서 다른 곳에서 사용 가능
type Board struct {
	Idx int `gorm:"PrimaryKey"`
	Author string
	Date time.Time
	Content string
	Views int // `gorm:"column:view_Count"`
}

func GetBoardData(db *gorm.DB)([]Board, error) {
	var boards []Board
	
	if result := db.Find(&boards); result.Error != nil {
		
		return nil, result.Error
	}
	return boards, nil
}