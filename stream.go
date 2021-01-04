package tesla

import (
	"bufio"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// StreamEvent of a vehicle returned by the Tesla API
type StreamEvent struct {
	Elevation  int       `json:"elevation"`
	EstHeading int       `json:"est_heading"`
	EstLat     float64   `json:"est_lat"`
	EstLng     float64   `json:"est_lng"`
	EstRange   int       `json:"est_range"`
	Heading    int       `json:"heading"`
	Odometer   float64   `json:"odometer"`
	Power      int       `json:"power"`
	Range      int       `json:"range"`
	ShiftState string    `json:"shift_state"`
	Soc        int       `json:"soc"`
	Speed      int       `json:"speed"`
	Timestamp  time.Time `json:"timestamp"`
}

// Stream starts a stream from the vehicle in the form of a go channel
func (v Vehicle) Stream() (chan *StreamEvent, chan error, error) {
	url := StreamURL + "/stream/" + strconv.Itoa(v.VehicleID) + "/?values=" + StreamParams
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(ActiveClient.Auth.Email, v.Tokens[0])
	resp, err := ActiveClient.HTTP.Do(req)

	if err != nil {
		return nil, nil, err
	}

	eventChan := make(chan *StreamEvent)
	errChan := make(chan error)
	go readStream(resp, eventChan, errChan)

	return eventChan, errChan, nil
}

func readStream(resp *http.Response, eventChan chan *StreamEvent, errChan chan error) {
	reader := bufio.NewReader(resp.Body)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	defer resp.Body.Close()

	for scanner.Scan() {
		streamEvent, err := parseStreamEvent(scanner.Text())
		if err == nil {
			eventChan <- streamEvent
		} else {
			errChan <- err
		}
	}
	errChan <- errors.New("http stream closed")
}

func parseStreamEvent(event string) (*StreamEvent, error) {
	data := strings.Split(event, ",")
	if len(data) != 13 {
		return nil, errors.New("invalid message from tesla api stream")
	}

	streamEvent := &StreamEvent{}
	timestamp, _ := strconv.ParseInt(data[0], 10, 64)
	streamEvent.Timestamp = time.Unix(0, timestamp*int64(time.Millisecond))
	streamEvent.Speed, _ = strconv.Atoi(data[1])
	streamEvent.Odometer, _ = strconv.ParseFloat(data[2], 64)
	streamEvent.Soc, _ = strconv.Atoi(data[3])
	streamEvent.Elevation, _ = strconv.Atoi(data[4])
	streamEvent.EstHeading, _ = strconv.Atoi(data[5])
	streamEvent.EstLat, _ = strconv.ParseFloat(data[6], 64)
	streamEvent.EstLng, _ = strconv.ParseFloat(data[7], 64)
	streamEvent.Power, _ = strconv.Atoi(data[8])
	streamEvent.ShiftState = data[9]
	streamEvent.Range, _ = strconv.Atoi(data[10])
	streamEvent.EstRange, _ = strconv.Atoi(data[11])
	streamEvent.Heading, _ = strconv.Atoi(data[12])
	return streamEvent, nil
}
