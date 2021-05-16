package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Tracker_sos struct {
	ID        uint32    `gorm:"type:UNSIGNED;primary_key;auto_increment;not null" json:"id"`
	UUID      string    `gorm:"type:char(36);default:null" json:"uuid"`
	WEARERID  uint      `gorm:"type:bigint(20);size:100;unique;default:null" json:"wearer_id"`
	DEVICEID  string    `gorm:"type:varchar(255);" json:"device_id"`
	BTUTCTIME time.Time `gorm:"type:datetime;" json:"BTUtcTime"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *Tracker_sos) FindAllSOS(db *gorm.DB) (*[]Tracker_sos, error) {
	var err error
	sos := []Tracker_sos{}
	err = db.Debug().Model(&Tracker_sos{}).Limit(100).Find(&sos).Error
	if err != nil {
		return &[]Tracker_sos{}, err
	}
	return &sos, err
}

func (u *Wearing_devices) Create(db *gorm.DB) (*Tracker_sos, error) {
	var err error
	sos := Tracker_sos{
		WEARERID:  u.WEARERID,
		DEVICEID:  u.IMEI,
		BTUTCTIME: time.Now(),
	}
	err = db.Debug().Model(&Tracker_sos{}).Create(&sos).Error
	if err != nil {
		return &Tracker_sos{}, err
	}
	return &sos, nil
}