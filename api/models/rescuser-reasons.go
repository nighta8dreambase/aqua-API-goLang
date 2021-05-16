package models

import (
	"errors"
	"time"

	"github.com/TanutN/Aqua/api/services"
	"github.com/jinzhu/gorm"
)

type RescuerResons struct {
	ID         uint32    `gorm:"primary_key;auto_increment;not null" json:"id"`
	UUID       string    `gorm:"index;type:char(36);default:null" json:"uuid"`
	TITLE      string    `gorm:"type:varchar(255);default:null" json:"title"`
	STATUS     uint      `gorm:"type:tinyint(4);default:1" json:"status"`
	PROJECT_ID uint32    `gorm:"type:bigint(20) UNSIGNED;default:null;" json:"project_id"`
	BRANCH_ID  uint32    `gorm:"type:bigint(20) UNSIGNED;default:null;" json:"branch_id"`
	CREATEDBY  uint32    `gorm:"type:bigint(20) UNSIGNED;default:null;" json:"created_by"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type ReasonInput struct {
	TITLE  string `gorm:"type:varchar(255);default:null" json:"title"`
	STATUS uint   `gorm:"type:tinyint(4);default:1" json:"status"`
}

func (u *RescuerResons) CreateReason(db *gorm.DB, projectKey string) (*RescuerResons, error) {
	var err error
	createReason := RescuerResons{
		UUID:      services.GenUUIDv4().String(),
		TITLE:     u.TITLE,
		STATUS:    u.STATUS,
		CREATEDBY: 1,
	}
	err = db.Debug().Model(&RescuerResons{}).Create(&createReason).Error
	//err = db.Debug().Model(&RescuerResons{}).Joins("left join projects on projects.id = rescuer_resons.project_id").Where("projects.project_key=?", projectKey).Create(&createReason).Error
	if err != nil {
		return &RescuerResons{}, err
	}
	return &createReason, nil
}

func (u *RescuerResons) FindAllReason(db *gorm.DB, projectKey string) (*[]RescuerResons, error) {
	var err error
	model := []RescuerResons{}
	//err = db.Debug().Model(&RescuerResons{}).Joins("left join projects on projects.id = rescuer_resons.project_id").Where("projects.project_key=?", projectKey).Find(&model).Error
	err = db.Debug().Model(&RescuerResons{}).Limit(100).Find(&model).Error
	if err != nil {
		return &[]RescuerResons{}, err
	}
	return &model, err
}

func (u *RescuerResons) FindReasonByID(db *gorm.DB, id string, projectKey string) (*RescuerResons, error) {
	var err error
	err = db.Debug().Model(RescuerResons{}).Where("uuid = ?", id).Take(&u).Error
	//err = db.Debug().Model(RescuerResons{}).Joins("left join projects on projects.id = rescuer_resons.project_id").Where("projects.project_key=?", projectKey).Where("rescuer_resons.uuid = ?", id).Take(&u).Error
	if err != nil {
		return &RescuerResons{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &RescuerResons{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *ReasonInput) UpdateAReason(db *gorm.DB, id string, user *RescuerResons) (*RescuerResons, error) {
	var err error
	err = db.Debug().Model(&user).Updates(&u).Error
	if err != nil {
		return user, err
	}
	return user, err
}

func (u *RescuerResons) DeleteAReason(db *gorm.DB, uid string) (int64, error) {
	db = db.Debug().Model(&RescuerResons{}).Where("uuid = ?", uid).Take(&RescuerResons{}).Delete(&RescuerResons{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
