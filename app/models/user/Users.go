package user

import (
	"github.com/arwahyu01/go-jwt/helpers/pagination"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;"`
	GoogleID        string    `json:"google_id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Email           string    `json:"email"`
	EmailVerifiedAt string    `json:"email_verified_at"`
	Password        string    `json:"-"`
	LevelID         string    `json:"level_id"`
	AccessGroupID   string    `json:"access_group_id"`
	Status          string    `json:"status"`
	Pembatalan      string    `json:"pembatalan"`
	Detail          string    `json:"detail"`
	RememberToken   string    `json:"remember_token"`
	CreatedAt       string    `json:"created_at"`
	UpdatedAt       string    `json:"updated_at"`
	DeletedAt       string    `json:"deleted_at"`
}

// BeforeCreate dipanggil sebelum data disimpan ke database
func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

func FetchAll(db *gorm.DB, page int, pageSize int) (*pagination.Page, error) {
	var users []Users
	data, err := pagination.Paginate(db, &users, page, pageSize)
	return data, err
}
