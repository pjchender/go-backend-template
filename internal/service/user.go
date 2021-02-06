package service

import (
	"github.com/google/uuid"
	"github.com/pjchender/go-backend-template/internal/model"
)

type GetUserRequest struct {
	ID uuid.UUID
}

func (svc *Service) GetUserByID(param GetUserRequest) (*model.User, error) {
	user, err := svc.db.GetUserByID(param.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
