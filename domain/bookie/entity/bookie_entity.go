package entity

import "time"

//go:generate mockgen -destination=mock/bookie_entity.go -source=bookie_entity.go
type Sport struct {
	SportID   int64 `json:"sport_id"`
	SportName struct {
		En string `json:"en"`
		Cn string `json:"cn"`
	} `json:"sport_name"`
}

type Event struct {
	EventID       int64     `json:"event_id"`
	EventTypeID   int64     `json:"event_type_id"`
	SportID       int64     `json:"sport_id"`
	CompetitionID int64     `json:"competition_id"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	EventStatusID int64     `json:"event_status_id"`
	IsNeutral     bool      `json:"is_neutral"`
	CreatedAt     time.Time `json:"created_at"`
}

type BookieUsecase interface {
	GetSportList() (sportList []Sport, err error)
	GetEventByEventID(eventID int64) (event Event, err error)
}

type BookieRepository interface {
	GetConfigSportList() (configSportList []Sport, err error)
	GetEventByEventID(eventID int64) (event Event, err error)
}
