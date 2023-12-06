package usecase

import (
	"fmt"
	"go-clean-arch/domain/bookie/entity"
	_logger "go-clean-arch/pkg/logger"
)

type bookieUsecase struct {
	bookieRepo entity.BookieRepository
	logger     _logger.Logger
}

func NewBookieUsecase(bookieRepo entity.BookieRepository, logger _logger.Logger) entity.BookieUsecase {
	return &bookieUsecase{
		bookieRepo: bookieRepo,
		logger:     logger,
	}
}

func (bu *bookieUsecase) GetSportList() (sportList []entity.Sport, err error) {
	// 取得 config sport list
	sportList, err = bu.bookieRepo.GetConfigSportList()
	if err != nil {
		bu.logger.Error("GetConfigSportList", err)
		err = fmt.Errorf("bookie_db_error")

		return
	}

	return
}

func (bu *bookieUsecase) GetEventByEventID(eventID int64) (event entity.Event, err error) {
	// 取得 event by event id
	event, err = bu.bookieRepo.GetEventByEventID(eventID)
	if err != nil {
		bu.logger.Error("GetEventByEventID", err)
		err = fmt.Errorf("bookie_db_error")

		return
	}

	return
}
