package tesla

import (
	"encoding/json"
	"errors"
	"strconv"
)

// CommandResponse is the response to a command sent to the Tesla API
type CommandResponse struct {
	Response struct {
		Reason string `json:"reason"`
		Result bool   `json:"result"`
	} `json:"response"`
}

// AutoParkRequest respresnts a request to autopark/summon a vehicle
type AutoParkRequest struct {
	Action    string  `json:"action,omitempty"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	VehicleID int     `json:"vehicle_id,omitempty"`
}

// AutoparkAbort tells the vehicle to abort an autopark/summon request
func (v Vehicle) AutoparkAbort() error {
	return v.autoPark("abort")
}

// AutoparkForward tells the vehicle to move forward
func (v Vehicle) AutoparkForward() error {
	return v.autoPark("start_forward")
}

// AutoparkReverse tells the vehicle to move backwards
func (v Vehicle) AutoparkReverse() error {
	return v.autoPark("start_reverse")
}

func (v Vehicle) autoPark(action string) error {
	url := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/autopark_request"
	driveState, _ := v.DriveState()
	autoParkRequest := &AutoParkRequest{
		VehicleID: v.VehicleID,
		Lat:       driveState.Latitude,
		Lon:       driveState.Longitude,
		Action:    action,
	}
	body, _ := json.Marshal(autoParkRequest)

	_, err := sendCommand(url, body)
	return err
}

// FlashLights flashes the vehicle's lights
func (v Vehicle) FlashLights() error {
	url := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/flash_lights"
	_, err := sendCommand(url, nil)
	return err
}

// HonkHorn honks the vehicle's horn
func (v *Vehicle) HonkHorn() error {
	url := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/honk_horn"
	_, err := sendCommand(url, nil)
	return err
}

// LockDoors locks the vehicle's doors
func (v Vehicle) LockDoors() error {
	url := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/door_lock"
	_, err := sendCommand(url, nil)
	return err
}

// UnlockDoors unlocks the vehicle's doors
func (v Vehicle) UnlockDoors() error {
	url := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/door_unlock"
	_, err := sendCommand(url, nil)
	return err
}

// MoveRoof sets the state of the panoramic roof to 1 of 4 presets or a specific percent.
// Each state and percentage: open = 100%, close = 0%, comfort = 80%, vent = %15
// To set a custom percentage provide a state of "move" along with a custom percentage.
func (v Vehicle) MoveRoof(state string, percent int) error {
	url := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/sun_roof_control"
	payload := `{"state": "` + state + `", "percent":` + strconv.Itoa(percent) + `}`
	_, err := ActiveClient.post(url, []byte(payload))
	return err
}

// OpenChargePort tells the vehicle to open the charge port
func (v Vehicle) OpenChargePort() error {
	url := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/charge_port_door_open"
	_, err := sendCommand(url, nil)
	return err
}

// OpenTrunk opens the specified trunk; possible values: 'front', 'rear'
func (v Vehicle) OpenTrunk(trunk string) error {
	url := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/trunk_open" // ?which_trunk=" + trunk
	payload := `{"which_trunk": "` + trunk + `"}`
	_, err := ActiveClient.post(url, []byte(payload))
	return err
}

// ResetValetPIN resets the valet mode PIN
func (v Vehicle) ResetValetPIN() error {
	url := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/reset_valet_pin"
	_, err := sendCommand(url, nil)
	return err
}

// SetChargeLimit sets the vehicle's charge limit to a specific percentage
func (v Vehicle) SetChargeLimit(percent int) error {
	url := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/set_charge_limit"
	payload := `{"percent": ` + strconv.Itoa(percent) + `}`
	_, err := ActiveClient.post(url, []byte(payload))
	return err
}

// SetChargeLimitMax sets the vehicle's charge limit to the max
func (v Vehicle) SetChargeLimitMax() error {
	url := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/charge_max_range"
	_, err := sendCommand(url, nil)
	return err
}

// SetChargeLimitStandard sets the vehicle's charge limit to the default standard
func (v Vehicle) SetChargeLimitStandard() error {
	url := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/charge_standard"
	_, err := sendCommand(url, nil)
	return err
}

// SetTemperature sets the driver and passenger zone temperatures
func (v Vehicle) SetTemperature(driver float64, passenger float64) error {
	driverTemp := strconv.FormatFloat(driver, 'f', -1, 32)
	passengerTemp := strconv.FormatFloat(passenger, 'f', -1, 32)
	url := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/set_temps?driver_temp=" + driverTemp + "&passenger_temp=" + passengerTemp
	_, err := ActiveClient.post(url, nil)
	return err
}

// Start starts the vehicle
func (v Vehicle) Start(password string) error {
	url := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/remote_start_drive?password=" + password
	_, err := sendCommand(url, nil)
	return err
}

// StartAirConditioning starts the vehicle's AC
func (v Vehicle) StartAirConditioning() error {
	url := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/auto_conditioning_start"
	_, err := sendCommand(url, nil)
	return err
}

// StopAirConditioning stops the vehicle's AC
func (v Vehicle) StopAirConditioning() error {
	url := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/auto_conditioning_stop"
	_, err := sendCommand(url, nil)
	return err
}

// StartCharging tells the vehicle to start charging
func (v Vehicle) StartCharging() error {
	url := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/charge_start"
	_, err := sendCommand(url, nil)
	return err
}

// StopCharging tells the vehicle to stop charging
func (v Vehicle) StopCharging() error {
	url := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/charge_stop"
	_, err := sendCommand(url, nil)
	return err
}

// ToggleHomelink tells the vehicle to toggle Homelink garage door opener
func (v Vehicle) ToggleHomelink() error {
	url := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/trigger_homelink"
	driveState, _ := v.DriveState()
	autoParkRequest := &AutoParkRequest{
		Lat: driveState.Latitude,
		Lon: driveState.Longitude,
	}
	body, _ := json.Marshal(autoParkRequest)

	_, err := sendCommand(url, body)
	return err
}

// Wakeup wakes up a vehicle that is powered off
func (v Vehicle) Wakeup() (*Vehicle, error) {
	url := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/wake_up"
	body, err := sendCommand(url, nil)
	if err != nil {
		return nil, err
	}
	vehicleResponse := &VehicleResponse{}
	err = json.Unmarshal(body, vehicleResponse)
	if err != nil {
		return nil, err
	}
	return vehicleResponse.Response, nil
}

// Sends a command to the vehicle
func sendCommand(url string, reqBody []byte) ([]byte, error) {
	body, err := ActiveClient.post(url, reqBody)
	if err != nil {
		return nil, err
	}
	if len(body) > 0 {
		response := &CommandResponse{}
		err = json.Unmarshal(body, response)
		if err != nil {
			return nil, err
		}
		if response.Response.Result != true && response.Response.Reason != "" {
			return nil, errors.New(response.Response.Reason)
		}
	}
	return body, nil
}
