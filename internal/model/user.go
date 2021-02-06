package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID            uuid.UUID `gorm:"primaryKey;uniqueIndex;type:uuid;default:uuid_generate_v4()"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	Username      string `gorm:"type:varchar(20);uniqueIndex"`
	Password      string
	IsSuperAdmin  bool
	TokenExpireAt time.Time
}

type UserExternal struct {
	ID            uuid.UUID `json:"id"`
	IsSuperAdmin  bool      `json:"is_super_admin"`
	TokenExpireAt time.Time `json:"token_expire_at"`
}

func (u *User) ToExternal() UserExternal {
	externalUser := UserExternal{
		ID:            u.ID,
		IsSuperAdmin:  u.IsSuperAdmin,
		TokenExpireAt: u.TokenExpireAt,
	}

	return externalUser
}
