package tesla

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

// Auth contains authentication details required to authenticate with the Tesla API
type Auth struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Email        string `json:"email"`
	GrantType    string `json:"grant_type"`
	Password     string `json:"password"`
}

// Token is returned by the Tesla API after a successful auth request
type Token struct {
	AccessToken  string `json:"access_token"`
	Expires      int64
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

// Client provides an API to communicate with the Tesla API
type Client struct {
	Auth           *Auth
	Endpoint       *url.URL
	HTTP           *http.Client
	Token          *Token
	StreamEndpoint *url.URL
}

var (
	BaseURL      = "https://owner-api.teslamotors.com/api/1"
	ActiveClient *Client
	StreamParams = "speed,odometer,soc,elevation,est_heading,est_lat,est_lng,power,shift_state,range,est_range,heading"
	StreamURL    = "https://streaming.vn.teslamotors.com"
)

// NewClient uses the given Auth to create a client for the Tesla API
func NewClient(auth *Auth) (*Client, error) {
	client := &Client{
		Auth: auth,
		HTTP: &http.Client{},
	}

	endpoint, err := url.Parse(BaseURL)
	if err != nil {
		return nil, err
	}
	client.Endpoint = endpoint

	stream, err := url.Parse(StreamURL)
	if err != nil {
		return nil, err
	}
	client.StreamEndpoint = stream

	token, err := client.authorize(auth)
	if err != nil {
		return nil, err
	}

	client.Token = token
	ActiveClient = client
	return client, nil
}

// TokenExpired returns true if the client's token expires within 30 minutes
func (c Client) TokenExpired() bool {
	exp := time.Unix(c.Token.Expires, 0)
	return time.Until(exp) < time.Duration(30*time.Minute)
}

// authorize uses the given Auth credentials to authenticate with the the Tesla API
func (c Client) authorize(auth *Auth) (*Token, error) {
	// clear the current token
	c.Token = nil
	u, _ := url.Parse(c.Endpoint.String())
	u.Path = path.Join("oauth/token")
	auth.GrantType = "password"
	data, _ := json.Marshal(auth)
	body, err := c.post(u.String(), data)
	if err != nil {
		return nil, err
	}

	token := &Token{}
	err = json.Unmarshal(body, token)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	token.Expires = now.Add(time.Second * time.Duration(token.ExpiresIn)).Unix()
	return token, nil
}

// // delete makes and HTTP DELETE request to the given url
func (c Client) delete(url string) error {
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	_, err := c.processRequest(req)
	return err
}

// get makes an HTTP GET request to the given url
func (c Client) get(url string) ([]byte, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	return c.processRequest(req)
}

// post makes an HTTP POST request to the given url with a provided body
func (c Client) post(url string, body []byte) ([]byte, error) {
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	return c.processRequest(req)
}

// put makes an HTTP PUT request to the given url with the provided body
func (c Client) put(url string, body []byte) ([]byte, error) {
	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	return c.processRequest(req)
}

// processRequest processes a provided http.Request
func (c Client) processRequest(req *http.Request) ([]byte, error) {
	if c.Token != nil && c.TokenExpired() {
		token, err := c.authorize(c.Auth)
		if err != nil {
			return nil, errors.New("unable to refresh token")
		}
		c.Token = token
	}

	c.setHeaders(req)
	res, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Sets the required headers for calls to the Tesla API
func (c Client) setHeaders(req *http.Request) {
	if c.Token != nil {
		req.Header.Set("Authorization", "Bearer "+c.Token.AccessToken)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
}
