package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pjchender/go-backend-template/internal/model"
)

type GetOrganizationRequest struct {
	ID uuid.UUID
}

type FirstOrCreateOrganizationByNISOrgIDRequest struct {
	NISOrgID string `json:"nisOrgId" binding:"required"`
	Name     string `json:"name"`
	NickName string `json:"nickName"`
}

func (svc *Service) GetOrganizationByID(param GetOrganizationRequest) (*model.Organization, error) {
	fmt.Println(param.ID)
	org, err := svc.db.GetOrganizationByID(param.ID)
	if err != nil {
		return nil, err
	}

	return org, err
}

func (svc *Service) FirstOrCreateOrganizationByNISOrgID(
	param FirstOrCreateOrganizationByNISOrgIDRequest,
) (*model.Organization, error) {
	org := model.Organization{
		NISOrgID: param.NISOrgID,
		Name:     param.Name,
		NickName: param.NickName,
	}

	theOrg, err := svc.db.FirstOrCreateOrganizationByNISOrgID(&org)
	if err != nil {
		return nil, err
	}

	return theOrg, nil
}
