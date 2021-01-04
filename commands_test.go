package tesla

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommands(t *testing.T) {
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
	assert.Nil(t, err)

	vehicle := vehicles[0]
	err = vehicle.AutoparkAbort()
	assert.Nil(t, err)

	err = vehicle.AutoparkForward()
	assert.Nil(t, err)

	err = vehicle.AutoparkReverse()
	assert.Nil(t, err)

	err = vehicle.ToggleHomelink()
	assert.Nil(t, err)

	_, err = vehicle.Wakeup()
	assert.Nil(t, err)

	err = vehicle.FlashLights()
	assert.Nil(t, err)

	err = vehicle.HonkHorn()
	assert.Nil(t, err)

	err = vehicle.OpenChargePort()
	assert.Nil(t, err)

	err = vehicle.ResetValetPIN()
	assert.Nil(t, err)

	err = vehicle.SetChargeLimit(50)
	assert.Nil(t, err)

	err = vehicle.SetChargeLimitStandard()
	assert.Equal(t, "already_standard", err.Error())

	err = vehicle.StartCharging()
	assert.Equal(t, "complete", err.Error())

	err = vehicle.StopCharging()
	assert.Nil(t, err)

	err = vehicle.SetChargeLimitMax()
	assert.Nil(t, err)

	err = vehicle.StartAirConditioning()
	assert.Nil(t, err)

	err = vehicle.StopAirConditioning()
	assert.Nil(t, err)

	err = vehicle.UnlockDoors()
	assert.Nil(t, err)

	err = vehicle.LockDoors()
	assert.Nil(t, err)

	err = vehicle.SetTemperature(68.1, 73.4)
	assert.Nil(t, err)

	err = vehicle.Start("pass")
	assert.Nil(t, err)

	err = vehicle.MoveRoof("vent", 0)
	assert.Nil(t, err)

	err = vehicle.MoveRoof("open", 0)
	assert.Nil(t, err)

	err = vehicle.MoveRoof("move", 50)
	assert.Nil(t, err)

	err = vehicle.MoveRoof("close", 0)
	assert.Nil(t, err)

	BaseURL = previousURL
}
