package tesla

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStates(t *testing.T) {
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
	vehicle := vehicles[0]
	status, err := vehicle.MobileEnabled()
	assert.Nil(t, err)
	assert.True(t, status)

	chargeState, err := vehicle.ChargeState()
	assert.Nil(t, err)
	assert.Equal(t, 90, chargeState.BatteryLevel)
	assert.Equal(t, 0.0, chargeState.ChargeRate)
	assert.Equal(t, "Complete", chargeState.ChargingState)

	climateState, err := vehicle.ClimateState()
	assert.Nil(t, err)
	assert.Equal(t, 22.0, climateState.DriverTempSetting)
	assert.Equal(t, 22.0, climateState.PassengerTempSetting)
	assert.False(t, climateState.IsRearDefrosterOn)

	driveState, err := vehicle.DriveState()
	assert.Nil(t, err)
	assert.Equal(t, 3.6, driveState.Latitude)
	assert.Equal(t, -149.1, driveState.Longitude)

	guiSettings, err := vehicle.GuiSettings()
	assert.Nil(t, err)
	assert.Equal(t, "mi/hr", guiSettings.GuiDistanceUnits)
	assert.Equal(t, "F", guiSettings.GuiTemperatureUnits)

	vehicleState, err := vehicle.VehicleState()
	assert.Nil(t, err)
	assert.Equal(t, 3, vehicleState.APIVersion)
	assert.True(t, vehicleState.CalendarSupported)
	assert.Equal(t, 0, vehicleState.Rt)

	BaseURL = previousURL
}
