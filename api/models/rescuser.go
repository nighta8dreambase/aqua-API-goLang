package models

import (
	"errors"
	"time"

	"github.com/TanutN/Aqua/api/services"
	"github.com/jinzhu/gorm"
)

type Sos_Rescuser struct {
	ID        uint32    `gorm:"type:UNSIGNED;primary_key;auto_increment;not null" json:"id"`
	UUID      string    `gorm:"type:char(36);default:null" json:"uuid"`
	WEARERID  uint      `gorm:"type:bigint(20);size:100;unique;default:null" json:"wearer_id"`
	DEVICEID  string    `gorm:"type:varchar(255);" json:"device_id"`
	BTUTCTIME time.Time `gorm:"type:datetime;" json:"BTUtcTime"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Rescuser struct {
	ID         uint32    `gorm:"primary_key;auto_increment;not null" json:"id"`
	UUID       string    `gorm:"index;type:char(36);default:null" json:"uuid"`
	NAME_TH    string    `gorm:"type:varchar(255);default:null" json:"name_th"`
	NAME_EN    string    `gorm:"type:varchar(255);default:null" json:"name_en"`
	PHONE_NO   string    `gorm:"type:varchar(255);default:null" json:"phone_no"`
	ADDRESS    string    `gorm:"type:varchar(255);default:null" json:"address"`
	RADIUS_JOB float64   `gorm:"type:DECIMAL(8,6);default:null" json:"radius_job"`
	LATITUDE   float64   `gorm:"type:DOUBLE(11,7);default:null" json:"latitude"`
	LONGTITUDE float64   `gorm:"type:DOUBLE(11,7);default:null" json:"longtitude"`
	PROJECT_ID uint32    `gorm:"type:bigint(20) UNSIGNED;default:null;" json:"project_id"`
	BRANCH_ID  uint32    `gorm:"type:bigint(20) UNSIGNED;default:null;" json:"branch_id"`
	CREATEDBY  uint32    `gorm:"type:bigint(20) UNSIGNED;default:null;" json:"created_by"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
type RescuserResponse struct {
	UUID       string    `json:"id"`
	NAME_TH    string    `json:"name_th"`
	NAME_EN    string    `json:"name_en"`
	PHONE_NO   string    `json:"phone_no"`
	ADDRESS    string    `json:"address"`
	RADIUS_JOB float64   `json:"radius_job"`
	LATITUDE   float64   `json:"latitude"`
	LONGTITUDE float64   `json:"longtitude"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type RescuserUpdate struct {
	NAME_TH    string  `json:"name_th"`
	NAME_EN    string  `json:"name_en"`
	PHONE_NO   string  `json:"phone_no"`
	ADDRESS    string  `json:"address"`
	RADIUS_JOB float64 `json:"radius_job"`
	LATITUDE   float64 `json:"latitude"`
	LONGTITUDE float64 `json:"longtitude"`
	CREATEDBY  uint32  `json:"created_by"`
}

type Rescuser_Input struct {
	SOS_ID  string `json:"sos_id"`
	COMMENT string `json:"comment"`
	REASON  string `json:"reason"`
}

func (u *Rescuser) CreateRescuser(db *gorm.DB) (*Rescuser, error) {
	var err error
	createRescues := Rescuser{
		UUID:       services.GenUUIDv4().String(),
		NAME_TH:    u.NAME_TH,
		NAME_EN:    u.NAME_EN,
		PHONE_NO:   u.PHONE_NO,
		ADDRESS:    u.ADDRESS,
		RADIUS_JOB: u.RADIUS_JOB,
		LATITUDE:   u.LONGTITUDE,
		LONGTITUDE: u.LATITUDE,
		CREATEDBY:  1,
	}
	err = db.Debug().Model(&Rescuser{}).Create(&createRescues).Error
	if err != nil {
		return &Rescuser{}, err
	}
	return &createRescues, nil
}

func (u *Rescuser) FindAllRescusers(db *gorm.DB) (*[]RescuserResponse, error) {
	var err error
	model := Rescuser{}
	rescusers := []RescuserResponse{}
	err = db.Debug().Model(&Rescuser{}).Select(RescuserResponse{}).Limit(100).Find(&model).Error
	if err != nil {
		return &[]RescuserResponse{}, err
	}
	return &rescusers, err
}

func (u *Rescuser) FindRescueByID(db *gorm.DB, id string) (*Rescuser, error) {
	var err error
	err = db.Debug().Model(Rescuser{}).Where("uuid = ?", id).Take(&u).Error
	if err != nil {
		return &Rescuser{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Rescuser{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *RescuserUpdate) UpdateARescuser(db *gorm.DB, id string, user *Rescuser) (*Rescuser, error) {
	var err error
	err = db.Debug().Model(&user).Updates(&u).Error
	if err != nil {
		return user, err
	}
	return user, err
}

func (u *Rescuser) DeleteARescuser(db *gorm.DB, uid string) (int64, error) {

	db = db.Debug().Model(&Rescuser{}).Where("uuid = ?", uid).Take(&Rescuser{}).Delete(&Rescuser{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (u *Rescuser_Input) FindAll(db *gorm.DB) (*[]Tracker_sos, error) {
	var err error
	sos := []Tracker_sos{}
	err = db.Debug().Model(&Tracker_sos{}).Limit(100).Find(&sos).Error
	if err != nil {
		return &[]Tracker_sos{}, err
	}
	return &sos, err
}
