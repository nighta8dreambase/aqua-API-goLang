package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Projects struct {
	ID             uint32    `gorm:"type:bigint(20);primary_key;auto_increment;not null" json:"id"`
	UUID           string    `gorm:"index;type:char(36);not null" json:"uuid"`
	NAME_TH        string    `gorm:"type:varchar(255);default:null" json:"name_th"`
	NAME_EN        string    `gorm:"type:varchar(255);default:null" json:"name_en"`
	PROJECT_KEY    string    `gorm:"type:varchar(255);default:null" json:"project_key"`
	PROJECT_DOMAIN string    `gorm:"type:varchar(255);default:null" json:"project_domain"`
	IS_QUARANTINE  uint      `gorm:"type:tinyint(1);default:null" json:"is_quarantine"`
	STATUS         uint      `gorm:"type:tinyint(4);default:1" json:"status"`
	CREATEDBY      uint32    `gorm:"type:bigint(20) UNSIGNED;default:null;" json:"created_by"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

/*func (u *RescuerResons) FindAllReason(db *gorm.DB, projectKey string) (*[]RescuerResons, error) {
	var err error
	model := []RescuerResons{}
	err = db.Debug().Model(&RescuerResons{}).Limit(100).Find(&model).Error
	if err != nil {
		return &[]RescuerResons{}, err
	}
	return &model, err
}*/

func (u *RescuerResons) FindProjectIdByKey(db *gorm.DB, key string) (*RescuerResons, error) {
	var err error
	err = db.Debug().Model(RescuerResons{}).Where("project_key = ?", key).Take(&u).Error
	if err != nil {
		return &RescuerResons{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &RescuerResons{}, errors.New("User Not Found")
	}
	return u, err
}

/*func (u *ReasonInput) UpdateAReason(db *gorm.DB, id string, user *RescuerResons) (*RescuerResons, error) {
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
}*/
