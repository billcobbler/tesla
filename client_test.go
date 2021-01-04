package tesla

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ChargeAlreadySetJSON = `{"response":{"reason":"already_standard","result":false}}`
	ChargeStateJSON      = `{"response":{"charging_state":"Complete","charge_limit_soc":90,"charge_limit_soc_std":90,"charge_limit_soc_min":50,"charge_limit_soc_max":100,"charge_to_max_range":false,"battery_heater_on":null,"not_enough_power_to_heat":null,"max_range_charge_counter":0,"fast_charger_present":null,"fast_charger_type":"\u003Cinvalid\u003E","battery_range":235.92,"est_battery_range":200.46,"ideal_battery_range":304.73,"battery_level":90,"usable_battery_level":90,"battery_current":null,"charge_energy_added":19.94,"charge_miles_added_rated":64.5,"charge_miles_added_ideal":83.0,"charger_voltage":null,"charger_pilot_current":null,"charger_actual_current":null,"charger_power":null,"time_to_full_charge":0.0,"trip_charging":null,"charge_rate":0.0,"charge_port_door_open":null,"motorized_charge_port":true,"scheduled_charging_start_time":null,"scheduled_charging_pending":false,"user_charge_enable_request":null,"charge_enable_request":true,"eu_vehicle":false,"charger_phases":null,"charge_port_latch":"\u003Cinvalid\u003E","charge_current_request":40,"charge_current_request_max":40,"managed_charging_active":false,"managed_charging_user_canceled":false,"managed_charging_start_time":null}}`
	ChargedJSON          = `{"response":{"reason":"complete","result":false}}`
	ClimateStateJSON     = `{"response":{"inside_temp":null,"outside_temp":null,"driver_temp_setting":22.0,"passenger_temp_setting":22.0,"left_temp_direction":17,"right_temp_direction":17,"is_auto_conditioning_on":null,"is_front_defroster_on":null,"is_rear_defroster_on":false,"fan_status":null,"is_climate_on":false,"min_avail_temp":15,"max_avail_temp":28,"seat_heater_left":0,"seat_heater_right":0,"seat_heater_rear_left":0,"seat_heater_rear_right":0,"seat_heater_rear_center":0,"seat_heater_rear_right_back":0,"seat_heater_rear_left_back":0,"smart_preconditioning":false}}`
	CommandResponseJSON  = `{"response":{"reason":"","result":true}}`
	DriveStateJSON       = `{"response":{"shift_state":null,"speed":null,"latitude":3.6,"longitude":-149.1,"heading":57,"gps_as_of":1452491619}}`
	GuiSettingsJSON      = `{"response":{"gui_distance_units":"mi/hr","gui_temperature_units":"F","gui_charge_rate_units":"mi/hr","gui_24_hour_time":true,"gui_range_display":"Rated"}}`
	TrueJSON             = `{"response":true}`
	VehiclesJSON         = `{"response":[{"color":null,"display_name":"Otto","id":123,"option_codes":"MDL3,RENA,AU01,BC3B,BS00,CDM0,CH07,PBCW,DA02,DCF0,DRLH,DV4W,FG31,HP00,IN3PB,LP01,ME02,MT310,PA00,PPSQ,PI01,PK00,PS01,PX00B,RFG3,SC01,SP00,SR01,SU00,TM00,TP03,W39B,X003,X007,X013,X027,X028,X031,X037,X040,YF00,","user_id":123,"vehicle_id":456,"vin":"abc123","tokens":["1","2"],"state":"online","id_s":"123","remote_start_enabled":true,"calendar_enabled":true,"notifications_enabled":true,"backseat_token":null,"backseat_token_updated_at":null}],"count":1}`
	VehicleStateJSON     = `{"response":{"api_version":3,"calendar_supported":true,"car_type":"s","car_version":"2.9.12","center_display_state":0,"dark_rims":false,"df":0,"dr":0,"exterior_color":"Black","ft":0,"has_spoiler":true,"locked":true,"notifications_supported":true,"odometer":3738.84633,"parsed_calendar_supported":true,"perf_config":"P2","pf":0,"pr":0,"rear_seat_heaters":1,"remote_start":false,"remote_start_supported":true,"rhd":false,"roof_color":"None","rt":0,"seat_type":1,"sun_roof_installed":2,"sun_roof_percent_open":0,"sun_roof_state":"unknown","third_row_seats":"None","valet_mode":false,"vehicle_name":"Macak","wheel_type":"Super21Gray"}}`
	WakeupResponseJSON   = `{"response":{"color":null,"display_name":"Otto","id":123,"option_codes":"MDL3,RENA,AU01,BC3B,BS00,CDM0,CH07,PBCW,DA02,DCF0,DRLH,DV4W,FG31,HP00,IN3PB,LP01,ME02,MT310,PA00,PPSQ,PI01,PK00,PS01,PX00B,RFG3,SC01,SP00,SR01,SU00,TM00,TP03,W39B,X003,X007,X013,X027,X028,X031,X037,X040,YF00,","user_id":123,"vehicle_id":456,"vin":"abc123","tokens":["1","2"],"state":"online","id_s":"123","remote_start_enabled":true,"calendar_enabled":true,"notifications_enabled":true,"backseat_token":null,"backseat_token_updated_at":null}}`
)

func TestClient(t *testing.T) {
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
	client, err := NewClient(auth)

	req, _ := http.NewRequest("GET", "http://foo.com", nil)
	client.setHeaders(req)
	assert.Equal(t, "Bearer sometoken123", req.Header.Get("Authorization"))
	assert.Equal(t, "application/json", req.Header.Get("Accept"))
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"))

	assert.Nil(t, err)
	assert.Equal(t, "sometoken123", client.Token.AccessToken)

	BaseURL = previousURL
}

func TestTokenExpired(t *testing.T) {
	expiredToken := &Token{
		AccessToken: "foo",
		TokenType:   "bar",
		ExpiresIn:   1,
		Expires:     0,
	}

	validToken := &Token{
		AccessToken: "foo",
		TokenType:   "bar",
		ExpiresIn:   1,
		Expires:     9999999999999,
	}

	client := &Client{
		Token: expiredToken,
	}

	assert.True(t, client.TokenExpired())

	client.Token = validToken
	assert.False(t, client.TokenExpired())

}

func serveHTTP(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		body, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		switch req.URL.String() {
		case "/oauth/token":
			checkHeaders(t, req)
			auth := &Auth{}
			json.Unmarshal(body, auth)
			assert.Equal(t, "nobody@example.com", auth.Email)
			assert.Equal(t, "pass", auth.Password)
			assert.Equal(t, "someclient123", auth.ClientID)
			assert.Equal(t, "somesecret456", auth.ClientSecret)
			w.WriteHeader(200)
			w.Write([]byte("{\"access_token\": \"sometoken123\"}"))
		case "/api/1/vehicles":
			checkHeaders(t, req)
			w.WriteHeader(200)
			w.Write([]byte(VehiclesJSON))
		case "/api/1/vehicles/123/mobile_enabled":
			checkHeaders(t, req)
			w.WriteHeader(200)
			w.Write([]byte(TrueJSON))
		case "/api/1/vehicles/123/data_request/charge_state":
			checkHeaders(t, req)
			w.WriteHeader(200)
			w.Write([]byte(ChargeStateJSON))
		case "/api/1/vehicles/123/data_request/climate_state":
			w.WriteHeader(200)
			w.Write([]byte(ClimateStateJSON))
		case "/api/1/vehicles/123/data_request/drive_state":
			checkHeaders(t, req)
			w.WriteHeader(200)
			w.Write([]byte(DriveStateJSON))
		case "/api/1/vehicles/123/data_request/gui_settings":
			checkHeaders(t, req)
			w.WriteHeader(200)
			w.Write([]byte(GuiSettingsJSON))
		case "/api/1/vehicles/123/data_request/vehicle_state":
			checkHeaders(t, req)
			w.WriteHeader(200)
			w.Write([]byte(VehicleStateJSON))
		case "/api/1/vehicles/123/wake_up":
			checkHeaders(t, req)
			w.WriteHeader(200)
			w.Write([]byte(WakeupResponseJSON))
		case "/api/1/vehicles/123/command/set_charge_limit":
			w.WriteHeader(200)
			assert.Equal(t, string(body), `{"percent": 50}`)
		case "/api/1/vehicles/123/command/charge_standard":
			checkHeaders(t, req)
			w.WriteHeader(200)
			w.Write([]byte(ChargeAlreadySetJSON))
		case "/api/1/vehicles/123/command/charge_start":
			checkHeaders(t, req)
			w.WriteHeader(200)
			w.Write([]byte(ChargedJSON))
		case "/api/1/vehicles/123/command/charge_stop",
			"/api/1/vehicles/123/command/charge_max_range",
			"/api/1/vehicles/123/command/charge_port_door_open",
			"/api/1/vehicles/123/command/flash_lights",
			"/api/1/vehicles/123/command/honk_horn",
			"/api/1/vehicles/123/command/auto_conditioning_start",
			"/api/1/vehicles/123/command/auto_conditioning_stop",
			"/api/1/vehicles/123/command/door_unlock",
			"/api/1/vehicles/123/command/door_lock",
			"/api/1/vehicles/123/command/reset_valet_pin",
			"/api/1/vehicles/123/command/set_temps?driver_temp=68.1&passenger_temp=73.4",
			"/api/1/vehicles/123/command/remote_start_drive?password=pass":
			checkHeaders(t, req)
			w.WriteHeader(200)
			w.Write([]byte(CommandResponseJSON))
		case "/stream/123/?values=speed,odometer,soc,elevation,est_heading,est_lat,est_lng,power,shift_state,range,est_range,heading":
			w.WriteHeader(200)
			events := StreamEventString + "\n" +
				StreamEventString + "\n" +
				BadStreamEventString + "\n"
			b := bytes.NewBufferString(events)
			b.WriteTo(w)
		case "/api/1/vehicles/123/command/autopark_request":
			w.WriteHeader(200)
			autoParkRequest := &AutoParkRequest{}
			err := json.Unmarshal(body, autoParkRequest)
			assert.Nil(t, err)
			var validAction bool
			switch autoParkRequest.Action {
			case "start_forward":
				validAction = true
			case "start_reverse":
				validAction = true
			case "abort":
				validAction = true
			}
			assert.True(t, validAction)
			assert.Equal(t, 456, autoParkRequest.VehicleID)
			assert.Equal(t, 3.6, autoParkRequest.Lat)
			assert.Equal(t, -149.1, autoParkRequest.Lon)
		case "/api/1/vehicles/123/command/trigger_homelink":
			w.WriteHeader(200)
			autoParkRequest := &AutoParkRequest{}
			err := json.Unmarshal(body, autoParkRequest)
			assert.Nil(t, err)
			assert.Equal(t, 3.6, autoParkRequest.Lat)
			assert.Equal(t, -149.1, autoParkRequest.Lon)
		case "/api/1/vehicles/123/command/sun_roof_control":
			w.WriteHeader(200)
			passed := false
			strBody := string(body)
			if strBody == `{"state": "vent", "percent":0}` {
				passed = true
			}
			if strBody == `{"state": "open", "percent":0}` {
				passed = true
			}
			if strBody == `{"state": "move", "percent":50}` {
				passed = true
			}
			if strBody == `{"state": "close", "percent":0}` {
				passed = true
			}
			assert.True(t, passed)

		}
	}))
}

func checkHeaders(t *testing.T, req *http.Request) {
	assert.Equal(t, "application/json", req.Header["Accept"][0])
	assert.Equal(t, "application/json", req.Header["Content-Type"][0])
}
