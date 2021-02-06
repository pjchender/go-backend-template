package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Organization struct {
	ID        uuid.UUID `gorm:"primaryKey;uniqueIndex;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	NISOrgID  string `gorm:"uniqueIndex"`
	Name      string
	NickName  string
}

type OrganizationExternal struct {
	ID       uuid.UUID `json:"id"`
	NISOrgID string    `json:"nisOrgId"`
	Name     string    `json:"name"`
	NickName string    `json:"nickName"`
}

func (o *Organization) ToExternal() OrganizationExternal {
	return OrganizationExternal{
		ID:       o.ID,
		NISOrgID: o.NISOrgID,
		Name:     o.Name,
		NickName: o.NickName,
	}
}
