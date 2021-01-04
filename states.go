package tesla

import (
	"encoding/json"
	"strconv"
)

// ChargeState represents the charge state of a vehicle
type ChargeState struct {
	BatteryCurrent              interface{} `json:"battery_current"`
	BatteryHeaterOn             bool        `json:"battery_heater_on"`
	BatteryLevel                int         `json:"battery_level"`
	BatteryRange                float64     `json:"battery_range"`
	ChargeCurrentRequest        int         `json:"charge_current_request"`
	ChargeCurrentRequestMax     int         `json:"charge_current_request_max"`
	ChargeEnableRequest         bool        `json:"charge_enable_request"`
	ChargeEnergyAdded           float64     `json:"charge_energy_added"`
	ChargeLimitSoc              int         `json:"charge_limit_soc"`
	ChargeLimitSocMax           int         `json:"charge_limit_soc_max"`
	ChargeLimitSocMin           int         `json:"charge_limit_soc_min"`
	ChargeLimitSocStd           int         `json:"charge_limit_soc_std"`
	ChargeMilesAddedIdeal       float64     `json:"charge_miles_added_ideal"`
	ChargeMilesAddedRated       float64     `json:"charge_miles_added_rated"`
	ChargePortDoorOpen          bool        `json:"charge_port_door_open"`
	ChargePortLatch             string      `json:"charge_port_latch"`
	ChargeRate                  float64     `json:"charge_rate"`
	ChargeToMaxRange            bool        `json:"charge_to_max_range"`
	ChargerActualCurrent        interface{} `json:"charger_actual_current"`
	ChargerPhases               interface{} `json:"charger_phases"`
	ChargerPilotCurrent         interface{} `json:"charger_pilot_current"`
	ChargerPower                interface{} `json:"charger_power"`
	ChargerVoltage              interface{} `json:"charger_voltage"`
	ChargingState               string      `json:"charging_state"`
	EstBatteryRange             float64     `json:"est_battery_range"`
	EuVehicle                   bool        `json:"eu_vehicle"`
	FastChargerPresent          bool        `json:"fast_charger_present"`
	FastChargerType             string      `json:"fast_charger_type"`
	IdealBatteryRange           float64     `json:"ideal_battery_range"`
	ManagedChargingActive       bool        `json:"managed_charging_active"`
	ManagedChargingStartTime    interface{} `json:"managed_charging_start_time"`
	ManagedChargingUserCanceled bool        `json:"managed_charging_user_canceled"`
	MaxRangeChargeCounter       int         `json:"max_range_charge_counter"`
	MotorizedChargePort         bool        `json:"motorized_charge_port"`
	NotEnoughPowerToHeat        bool        `json:"not_enough_power_to_heat"`
	ScheduledChargingPending    bool        `json:"scheduled_charging_pending"`
	ScheduledChargingStartTime  interface{} `json:"scheduled_charging_start_time"`
	TimeToFullCharge            float64     `json:"time_to_full_charge"`
	TripCharging                interface{} `json:"trip_charging"`
	UsableBatteryLevel          int         `json:"usable_battery_level"`
	UserChargeEnableRequest     interface{} `json:"user_charge_enable_request"`
}

// ClimateState represents the state of climate in a vehicle
type ClimateState struct {
	DriverTempSetting       float64     `json:"driver_temp_setting"`
	FanStatus               interface{} `json:"fan_status"`
	InsideTemp              float64     `json:"inside_temp"`
	IsAutoConditioningOn    bool        `json:"is_auto_conditioning_on"`
	IsClimateOn             bool        `json:"is_climate_on"`
	IsFrontDefrosterOn      int         `json:"is_front_defroster_on"`
	IsRearDefrosterOn       bool        `json:"is_rear_defroster_on"`
	LeftTempDirection       float64     `json:"left_temp_direction"`
	MaxAvailTemp            float64     `json:"max_avail_temp"`
	MinAvailTemp            float64     `json:"min_avail_temp"`
	OutsideTemp             float64     `json:"outside_temp"`
	PassengerTempSetting    float64     `json:"passenger_temp_setting"`
	RightTempDirection      float64     `json:"right_temp_direction"`
	SeatHeaterLeft          int         `json:"seat_heater_left"`
	SeatHeaterRearCenter    int         `json:"seat_heater_rear_center"`
	SeatHeaterRearLeft      int         `json:"seat_heater_rear_left"`
	SeatHeaterRearLeftBack  int         `json:"seat_heater_rear_left_back"`
	SeatHeaterRearRight     int         `json:"seat_heater_rear_right"`
	SeatHeaterRearRightBack int         `json:"seat_heater_rear_right_back"`
	SeatHeaterRight         int         `json:"seat_heater_right"`
	SmartPreconditioning    bool        `json:"smart_preconditioning"`
}

// DriveState represents the drive state of a vehicle
type DriveState struct {
	GpsAsOf    int64       `json:"gps_as_of"`
	Heading    int         `json:"heading"`
	Latitude   float64     `json:"latitude"`
	Longitude  float64     `json:"longitude"`
	ShiftState interface{} `json:"shift_state"`
	Speed      float64     `json:"speed"`
}

// GuiSettings represents the GUI settings of a vehicle
type GuiSettings struct {
	Gui24HourTime       bool   `json:"gui_24_hour_time"`
	GuiChargeRateUnits  string `json:"gui_charge_rate_units"`
	GuiDistanceUnits    string `json:"gui_distance_units"`
	GuiRangeDisplay     string `json:"gui_range_display"`
	GuiTemperatureUnits string `json:"gui_temperature_units"`
}

// VehicleState represents the state of a vehicle
type VehicleState struct {
	APIVersion              int     `json:"api_version"`
	AutoParkState           string  `json:"autopark_state"`
	AutoParkStateV2         string  `json:"autopark_state_v2"`
	CalendarSupported       bool    `json:"calendar_supported"`
	CarType                 string  `json:"car_type"`
	CarVersion              string  `json:"car_version"`
	CenterDisplayState      int     `json:"center_display_state"`
	DarkRims                bool    `json:"dark_rims"`
	Df                      int     `json:"df"`
	Dr                      int     `json:"dr"`
	ExteriorColor           string  `json:"exterior_color"`
	Ft                      int     `json:"ft"`
	HasSpoiler              bool    `json:"has_spoiler"`
	Locked                  bool    `json:"locked"`
	NotificationsSupported  bool    `json:"notifications_supported"`
	Odometer                float64 `json:"odometer"`
	ParsedCalendarSupported bool    `json:"parsed_calendar_supported"`
	PerfConfig              string  `json:"perf_config"`
	Pf                      int     `json:"pf"`
	Pr                      int     `json:"pr"`
	RearSeatHeaters         int     `json:"rear_seat_heaters"`
	RemoteStart             bool    `json:"remote_start"`
	RemoteStartSupported    bool    `json:"remote_start_supported"`
	Rhd                     bool    `json:"rhd"`
	RoofColor               string  `json:"roof_color"`
	Rt                      int     `json:"rt"`
	SeatType                int     `json:"seat_type"`
	SpoilerType             string  `json:"spoiler_type"`
	SunRoofInstalled        int     `json:"sun_roof_installed"`
	SunRoofPercentOpen      int     `json:"sun_roof_percent_open"`
	SunRoofState            string  `json:"sun_roof_state"`
	ThirdRowSeats           string  `json:"third_row_seats"`
	ValetMode               bool    `json:"valet_mode"`
	VehicleName             string  `json:"vehicle_name"`
	WheelType               string  `json:"wheel_type"`
}

// StateResponse is the response received when requesting the states of a vehicle
type StateResponse struct {
	Response struct {
		*ChargeState
		*ClimateState
		*DriveState
		*GuiSettings
		*VehicleState
	} `json:"response"`
}

// BoolStateResponse is the response when a state is requested for a simple boolean state
type BoolStateResponse struct {
	Bool bool `json:"response"`
}

// MobileEnabled returns true if the vehicle is mobile enabled for Tesla API control
func (v *Vehicle) MobileEnabled() (bool, error) {
	body, err := ActiveClient.get(BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/mobile_enabled")
	if err != nil {
		return false, err
	}
	response := &BoolStateResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return false, err
	}
	return response.Bool, nil
}

// ChargeState returns the state of charge for the vehicle
func (v *Vehicle) ChargeState() (*ChargeState, error) {
	state, err := fetchState("/charge_state", v.ID)
	if err != nil {
		return nil, err
	}
	return state.Response.ChargeState, nil
}

// ClimateState returns the climate state of the vehicle
func (v Vehicle) ClimateState() (*ClimateState, error) {
	state, err := fetchState("/climate_state", v.ID)
	if err != nil {
		return nil, err
	}
	return state.Response.ClimateState, nil
}

// DriveState returns the drive state of the vehicle
func (v Vehicle) DriveState() (*DriveState, error) {
	state, err := fetchState("/drive_state", v.ID)
	if err != nil {
		return nil, err
	}
	return state.Response.DriveState, nil
}

// GuiSettings returns the GUI settings of the vehicle
func (v Vehicle) GuiSettings() (*GuiSettings, error) {
	state, err := fetchState("/gui_settings", v.ID)
	if err != nil {
		return nil, err
	}
	return state.Response.GuiSettings, nil
}

// VehicleState returns the state of the vehicle
func (v Vehicle) VehicleState() (*VehicleState, error) {
	state, err := fetchState("/vehicle_state", v.ID)
	if err != nil {
		return nil, err
	}
	return state.Response.VehicleState, nil
}

func fetchState(resource string, id int64) (*StateResponse, error) {
	state := &StateResponse{}
	body, err := ActiveClient.get(BaseURL + "/vehicles/" + strconv.FormatInt(id, 10) + "/data_request" + resource)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, state)
	if err != nil {
		return nil, err
	}
	return state, nil
}
