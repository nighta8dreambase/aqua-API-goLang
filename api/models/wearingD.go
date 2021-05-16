package models

import (
	"time"

	"encoding/json"

	"github.com/jinzhu/gorm"
)

type Wearing_devices struct {
	ID        uint32    `gorm:"type:bigint(20);type:UNSIGNED;primary_key;auto_increment;not null" json:"id"`
	UUID      string    `gorm:"type:char(36);not null" json:"uuid"`
	WEARERID  uint      `gorm:"type:bigint(20);size:100;unique;default:null" json:"wearer_id"`
	DEVICEID  string    `gorm:"type:bigint(20);default:null" json:"device_id"`
	IMEI      string    `gorm:"type:varchar(255);default:null;" json:"imei"`
	STATUS    uint      `gorm:"type:tinyint(4);primary_key;default:1" json:"status"`
	CREATEDBY uint      `gorm:"type:bigint(20);type:UNSIGNED;default:null" json:"created_by"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Input_Wearing struct {
	IMEI  string          `json:"imei"`
	MODEL string          `json:"model"`
	Data  json.RawMessage `json:"data"`
}

func (u *Wearing_devices) FindAllDevices(db *gorm.DB) (*[]Wearing_devices, error) {
	var err error
	sos := []Wearing_devices{}
	err = db.Debug().Model(&Tracker_sos{}).Limit(100).Find(&sos).Error
	if err != nil {
		return &[]Wearing_devices{}, err
	}
	return &sos, err
}

func (u *Input_Wearing) FindSosDevicesByImei(db *gorm.DB) (*Wearing_devices, error) {
	var err error
	sos := Wearing_devices{}
	err = db.Debug().Model(&Wearing_devices{}).Where(&Wearing_devices{IMEI: u.IMEI}).Order("id desc").First(&sos).Error
	if err != nil {
		return &Wearing_devices{}, err
	}
	return &sos, err
}
