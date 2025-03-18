package models

import (
	"github.com/arwahyu01/go-jwt/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	ID       uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Nama     string    `gorm:"type:varchar(255);not null" json:"nama"`
	Username string    `gorm:"type:varchar(255);unique;not null" json:"username"`
	Password string    `gorm:"type:varchar(255);not null" json:"-"`
}

// BeforeCreate dipanggil sebelum data disimpan ke database
func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

func GetAllUser(db *gorm.DB, page int, pageSize int) (*helpers.Pagination, error) {
	var users []Users
	pagination, err := helpers.Paginate(db, &users, page, pageSize)
	return pagination, err
}
