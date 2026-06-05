package model

import (
	internalerror "email/internal/domain/campaign/internal-error"
	"time"

	"github.com/google/uuid"
)

type Contact struct {
	Email string `validate:"required,email"`
}

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"required,min=5,max=24"`
	CreatedAt time.Time `validate:"required"`
	Content   string    `validate:"required,min=5,max=500"`
	Contacts  []Contact `validate:"required,min=1,dive"` // dive is used to validate each element in the slice
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		contacts[i].Email = email
	}

	campaign := &Campaign{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
	}

	validationErr := internalerror.ValidateStruct(campaign)

	if validationErr != nil {
		return nil, validationErr
	}

	return campaign, nil
}
