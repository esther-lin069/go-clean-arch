package usecase

import (
	"fmt"
	"go-clean-arch/domain/bookie/entity"
	mock_entity "go-clean-arch/domain/bookie/entity/mock"
	mock_logger "go-clean-arch/pkg/logger/mock"
	"reflect"
	"testing"
	"time"

	"go.uber.org/mock/gomock"
)

func Test_bookieUsecase_GetSportList(t *testing.T) {
	ctl := gomock.NewController(t)

	// mocks
	mockLogger := mock_logger.NewMockLogger(ctl)
	mockBookieRepository := mock_entity.NewMockBookieRepository(ctl)

	// mock return value
	soccer := entity.Sport{
		SportID: 1,
		SportName: struct {
			En string "json:\"en\""
			Cn string "json:\"cn\""
		}{
			En: "Soccer",
			Cn: "足球",
		},
	}

	dbErr := fmt.Errorf("db_error")
	// logger expect
	mockLogger.EXPECT().Error("GetConfigSportList", dbErr).Times(1)

	// repo expect
	gomock.InOrder(

		// 球種列表回傳足球
		mockBookieRepository.
			EXPECT().
			GetConfigSportList().
			Return([]entity.Sport{
				soccer,
			}, nil).
			Times(1),

		// 球種列表回傳錯誤
		mockBookieRepository.
			EXPECT().
			GetConfigSportList().
			Return([]entity.Sport{}, dbErr).
			Times(1),
	)

	tests := []struct {
		name          string
		wantSportList []entity.Sport
		wantErr       bool
	}{
		{
			name: "成功取得球種列表",
			wantSportList: []entity.Sport{
				soccer,
			},
			wantErr: false,
		},
		{
			name:          "失敗取得球種列表",
			wantSportList: []entity.Sport{},
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bu := NewBookieUsecase(mockBookieRepository, mockLogger)

			gotSportList, err := bu.GetSportList()
			if (err != nil) != tt.wantErr {
				t.Errorf("bookieUsecase.GetSportList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSportList, tt.wantSportList) {
				t.Errorf("bookieUsecase.GetSportList() = %v, want %v", gotSportList, tt.wantSportList)
			}
		})
	}
}

func Test_bookieUsecase_GetEventByEventID(t *testing.T) {

	ctl := gomock.NewController(t)

	// mocks
	mockLogger := mock_logger.NewMockLogger(ctl)
	mockBookieRepository := mock_entity.NewMockBookieRepository(ctl)

	// mock return value
	event1 := entity.Event{
		EventID:       1,
		EventTypeID:   1,
		SportID:       1,
		CompetitionID: 1,
		StartTime:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		EndTime:       time.Time{},
		EventStatusID: 1,
	}

	// logger expect
	mockLogger.EXPECT().Error("GetEventByEventID", fmt.Errorf("db_error")).Times(1)

	// repo expect

	// 回傳1號賽事
	mockBookieRepository.EXPECT().GetEventByEventID(int64(1)).Return(event1, nil).Times(1)

	// 找不到賽事或DB錯誤
	mockBookieRepository.EXPECT().GetEventByEventID(int64(2)).Return(entity.Event{}, fmt.Errorf("db_error")).Times(1)

	type args struct {
		eventID int64
	}
	tests := []struct {
		name      string
		args      args
		wantEvent entity.Event
		wantErr   bool
	}{
		{
			name: "成功取得賽事",
			args: args{
				eventID: 1,
			},
			wantEvent: event1,
			wantErr:   false,
		},
		{
			name: "失敗取得賽事",
			args: args{
				eventID: 2,
			},
			wantEvent: entity.Event{},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bu := NewBookieUsecase(mockBookieRepository, mockLogger)

			gotEvent, err := bu.GetEventByEventID(tt.args.eventID)
			if (err != nil) != tt.wantErr {
				t.Errorf("bookieUsecase.GetEventByEventID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEvent, tt.wantEvent) {
				t.Errorf("bookieUsecase.GetEventByEventID() = %v, want %v", gotEvent, tt.wantEvent)
			}
		})
	}
}
