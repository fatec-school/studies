package service

import (
	"email/internal/adapter/dto"
	"email/internal/domain/campaign/internal-error"
	"email/internal/domain/campaign/model"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	name    = "test"
	content = "content"
	emails  = []string{"email1@email.com", "email2@email.com"}
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *model.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func validCampaign(campaign *model.Campaign) bool {
	newCampaign := dto.NewCampaignRequest{
		Name:    name,
		Content: content,
		Emails:  emails,
	}

	if campaign.Name != newCampaign.Name ||
		campaign.Content != newCampaign.Content ||
		len(campaign.Contacts) != len(newCampaign.Emails) {
		return false
	}

	return true
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(validCampaign)).Return(nil)
	service := Service{repositoryMock}
	newCampaign := dto.NewCampaignRequest{
		Name:    name,
		Content: content,
		Emails:  emails,
	}

	id, err := service.Create(&newCampaign)

	assert.NotNil(id)
	assert.NotEmpty(id)
	assert.Nil(err)
	repositoryMock.AssertExpectations(t)
}

func Test_Create_SaveCampaign(t *testing.T) {
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(validCampaign)).Return(nil)
	service := Service{Repository: repositoryMock}
	newCampaign := dto.NewCampaignRequest{
		Name:    name,
		Content: content,
		Emails:  emails,
	}

	service.Create(&newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(validCampaign)).Return(nil)
	service := Service{Repository: repositoryMock}
	newCampaign := dto.NewCampaignRequest{
		Name:    name,
		Content: content,
		Emails:  emails,
	}

	newCampaign.Name = ""
	id, err := service.Create(&newCampaign)

	assert.Empty(id)
	assert.NotNil(err)
	assert.Equal("name is required", err.Error())
}

func Test_Create_ValidateRepositoryError(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))
	service := Service{Repository: repositoryMock}
	newCampaign := dto.NewCampaignRequest{
		Name:    name,
		Content: content,
		Emails:  emails,
	}

	_, err := service.Create(&newCampaign)

	assert.True(errors.Is(err, internalerror.ErrInternalError))
}
