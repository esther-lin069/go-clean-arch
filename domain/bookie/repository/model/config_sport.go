package model

type ConfigSport struct {
	SportID     int64  `gorm:"column:sport_id"`
	SportNameCn string `gorm:"column:sport_name_cn"`
	SportNameEn string `gorm:"column:sport_name_en"`
}

// 資料表名稱
func (ConfigSport) TableName() string {
	return "config_sport"
}
