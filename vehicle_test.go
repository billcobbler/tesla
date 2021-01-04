package tesla

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVehicles(t *testing.T) {
	ts := serveHTTP(t)
	defer ts.Close()
	previousURL := BaseURL
	BaseURL = ts.URL + "/api/1"

	auth := &Auth{
		GrantType:    "password",
		ClientID:     "someclient123",
		ClientSecret: "somesecret456",
		Email:        "nobody@example.com",
		Password:     "pass",
	}
	client, _ := NewClient(auth)

	vehicles, err := client.Vehicles()
	v := vehicles[0]
	assert.Nil(t, err)
	assert.Equal(t, "Otto", v.DisplayName)
	assert.Equal(t, "online", v.State)
	assert.True(t, v.CalendarEnabled)
	assert.True(t, v.NotificationsEnabled)
	assert.True(t, v.RemoteStartEnabled)

	BaseURL = previousURL
}
