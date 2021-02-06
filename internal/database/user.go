package database

import (
	"github.com/google/uuid"
	"github.com/pjchender/go-backend-template/internal/model"
)

func (d *GormDatabase) CreateUser(user *model.User) error {
	return d.DB.Create(user).Error
}

func (d *GormDatabase) GetUserByID(id uuid.UUID) (*model.User, error) {
	user := &model.User{}
	if err := d.DB.Take(user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}
