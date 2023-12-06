package repository

import (
	"go-clean-arch/domain/bookie/entity"
	"go-clean-arch/domain/bookie/repository/model"

	"gorm.io/gorm"
)

type bookieRepository struct {
	masterDB *gorm.DB
	slaveDB  *gorm.DB
}

func NewBookieRepository(masterDB *gorm.DB, slaveDB *gorm.DB) entity.BookieRepository {
	return &bookieRepository{
		masterDB: masterDB,
		slaveDB:  slaveDB,
	}
}

func (br *bookieRepository) GetConfigSportList() (configSportList []entity.Sport, err error) {
	var tempSportList []model.ConfigSport

	// 查詢
	err = br.slaveDB.Find(&tempSportList).Error
	if err != nil {
		return
	}

	for _, v := range tempSportList {
		// 格式化回傳
		configSportList = append(configSportList, entity.Sport{
			SportID: v.SportID,
			SportName: struct {
				En string "json:\"en\""
				Cn string "json:\"cn\""
			}{
				En: v.SportNameEn,
				Cn: v.SportNameCn,
			},
		})
	}

	return
}

func (br *bookieRepository) GetEventByEventID(eventID int64) (event entity.Event, err error) {
	var tempEvent model.Event

	// 查詢，這邊沒有找到 event 資料時會視為錯誤
	err = br.slaveDB.Where("event_id = ?", eventID).First(&tempEvent).Error
	if err != nil {
		return
	}

	// 格式化回傳
	event = entity.Event{
		EventID:       tempEvent.EventID,
		EventTypeID:   tempEvent.EventTypeID,
		SportID:       tempEvent.SportID,
		CompetitionID: tempEvent.CompetitionID,
		StartTime:     tempEvent.StartTime,
		EndTime:       tempEvent.EndTime,
		EventStatusID: tempEvent.EventStatusID,
		IsNeutral:     tempEvent.IsNeutral,
		CreatedAt:     tempEvent.CreatedAt,
	}

	return
}
