package database

import (
	"github.com/google/uuid"
	"github.com/pjchender/go-backend-template/internal/model"
)

func (d *GormDatabase) GetOrganizationByID(id uuid.UUID) (*model.Organization, error) {
	org := model.Organization{}
	err := d.DB.Take(&org, id).Error
	if err != nil {
		return nil, err
	}

	return &org, nil
}

func (d *GormDatabase) FirstOrCreateOrganizationByNISOrgID(org *model.Organization) (*model.Organization, error) {
	err := d.DB.Where(model.Organization{NISOrgID: org.NISOrgID}).FirstOrCreate(org).Error
	if err != nil {
		return nil, err
	}

	return org, nil
}
