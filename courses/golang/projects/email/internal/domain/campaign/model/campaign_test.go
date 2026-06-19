package model

import (
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign x"
	content  = "values for campaign x"
	contacts = []string{"email@email.com", "email2@email.com"}
	fake     = faker.New()
)

func Test_NewCampaign_Create(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	// if campaign.ID != "1" {
	// 	t.Errorf("id expected 1: %s", campaign.ID)
	// } else if campaign.Name != name {
	// 	t.Errorf("expected name %s: %s", name, campaign.Name)
	// } else if campaign.Content != content {
	// 	t.Errorf("expected content %s: %s", content, campaign.Content)
	// } else if len(campaign.Contacts) != len(contacts) {
	// 	t.Errorf("expected %d emails: %d", len(contacts), len(campaign.Content))
	// }

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedAtIsGreaterThanNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedAt, now)
}

func Test_NewCampaign_NameIsRequired(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("name is required", err.Error())
}

func Test_NewCampaign_NameMustBeGreaterThan5Characters(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("x", content, contacts)

	assert.Equal("name must be at least 5 characters long", err.Error())
}

func Test_NewCampaign_NameMustBeLessThan24Characters(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(124), content, contacts)

	assert.Equal("name must be at most 24 characters long", err.Error())
}

func Test_NewCampaign_ContentIsRequired(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content is required", err.Error())
}

func Test_NewCampaign_ContentMustBeGreaterThan5Characters(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "x", contacts)
	assert.Equal("content must be at least 5 characters long", err.Error())
}

func Test_NewCampaign_ContentMustBeLessThan500Characters(t *testing.T) {
	assert := assert.New(t)
	longContent := fake.Lorem().Text(510)

	_, err := NewCampaign(name, longContent, contacts)

	assert.Equal("content must be at most 500 characters long", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.Equal("contacts must be at least 1 characters long", err.Error())
}

func Test_NewCampaign_MustValidateContactsEmail(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{"invalid-email"})

	assert.Equal("email must be a valid email", err.Error())
}

func Test_NewCampaign_StatusMustBePending(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal(CampaignStatusPending, campaign.Status)
}
