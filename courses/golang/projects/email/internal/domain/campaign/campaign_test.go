package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign x"
	content := "value"
	contacts := []string{"email@email.com", "email2@email.com"}

	campaign := NewCampaign(name, content, contacts)

	// if campaign.ID != "1" {
	// 	t.Errorf("id expected 1: %s", campaign.ID)
	// } else if campaign.Name != name {
	// 	t.Errorf("expected name %s: %s", name, campaign.Name)
	// } else if campaign.Content != content {
	// 	t.Errorf("expected content %s: %s", content, campaign.Content)
	// } else if len(campaign.Contacts) != len(contacts) {
	// 	t.Errorf("expected %d emails: %d", len(contacts), len(campaign.Content))
	// }

	assert.Equal(campaign.ID, "1")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}
