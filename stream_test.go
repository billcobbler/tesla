package tesla

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStreamSpec(t *testing.T) {
	ts := serveHTTP(t)
	defer ts.Close()
	previousURL := BaseURL
	BaseURL = ts.URL + "/api/1"
	vehicle := &Vehicle{}
	vehicle.VehicleID = 123
	vehicle.Tokens = []string{"456", "789"}

	previousStreamURL := StreamURL
	StreamURL = ts.URL

	eventChan, errChan, err := vehicle.Stream()
	assert.Nil(t, err)

	select {
	case event := <-eventChan:
		assert.Equal(t, 65, event.Speed)
	case err = <-errChan:
		assert.Nil(t, err)
	}
	select {
	case event := <-eventChan:
		assert.Equal(t, 65, event.Speed)
	case err = <-errChan:
		assert.Nil(t, err)
	}
	select {
	case event := <-eventChan:
		assert.Nil(t, event)
	case err = <-errChan:
		assert.Equal(t, "invalid message from tesla api stream", err.Error())
	}
	select {
	case event := <-eventChan:
		assert.Nil(t, event)
	case err = <-errChan:
		assert.Equal(t, "http stream closed", err.Error())
	}

	BaseURL = previousURL
	StreamURL = previousStreamURL
}
