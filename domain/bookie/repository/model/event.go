package model

import "time"

type Event struct {
	EventID       int64     `gorm:"column:event_id"`
	EventTypeID   int64     `gorm:"column:event_type_id"`
	SportID       int64     `gorm:"column:sport_id"`
	CompetitionID int64     `gorm:"column:competition_id"`
	StartTime     time.Time `gorm:"column:start_time"`
	EndTime       time.Time `gorm:"column:end_time"`
	EventStatusID int64     `gorm:"column:event_status_id"`
	IsNeutral     bool      `gorm:"column:is_neutral"`
	CreatedAt     time.Time `gorm:"column:created_at"`
}

// 資料表名稱
func (Event) TableName() string {
	return "event"
}
