package model

import (
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
	Contacts  []Contact `validate:"required,min=1"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		contacts[i].Email = email
	}

	return &Campaign{
		ID:        uuid.NewString(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		CreatedAt: time.Now(),
	}, nil
}
