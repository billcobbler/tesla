package tesla

import (
	"encoding/json"
	"net/url"
	"path"
)

// Vehicle returned from the Tesla API
type Vehicle struct {
	BackseatToken          interface{} `json:"backseat_token"`
	BackseatTokenUpdatedAt interface{} `json:"backseat_token_updated_at"`
	CalendarEnabled        bool        `json:"calendar_enabled"`
	Color                  interface{} `json:"color"`
	DisplayName            string      `json:"display_name"`
	ID                     int64       `json:"id"`
	IDS                    string      `json:"id_s"`
	NotificationsEnabled   bool        `json:"notifications_enabled"`
	OptionCodes            string      `json:"option_codes"`
	RemoteStartEnabled     bool        `json:"remote_start_enabled"`
	State                  string      `json:"state"`
	Tokens                 []string    `json:"tokens"`
	VehicleID              int         `json:"vehicle_id"`
	Vin                    string      `json:"vin"`
}

// VehicleResponse represents vehicle details from the Tesla API
type VehicleResponse struct {
	Count    int      `json:"count"`
	Response *Vehicle `json:"response"`
}

// Vehicles returned from the Tesla API (some API responses contain multiple vehicles)
type Vehicles []struct {
	*Vehicle
}

// VehiclesResponse represents multiple vehicle details from the Tesla API
type VehiclesResponse struct {
	Response Vehicles `json:"response"`
	Count    int      `json:"count"`
}

// Vehicles fetches all vehicles associated with a Tesla account
func (c *Client) Vehicles() (Vehicles, error) {
	u, _ := url.Parse(c.Endpoint.String())
	u.Path = path.Join(u.Path, "vehicles")
	vehiclesResponse := &VehiclesResponse{}
	body, err := c.get(u.String())
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, vehiclesResponse)
	if err != nil {
		return nil, err
	}
	return vehiclesResponse.Response, nil
}
