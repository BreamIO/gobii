package gaze_test

import (
	"github.com/zephyyrr/gobii/gaze"
	"testing"
)

func TestVersion(t *testing.T) {
	if gaze.Version() != "..." {
		t.Error(gaze.Version())
	}
}

/*
func TestCreateEyetracker(t *testing.T) {
	const etURL = "tet-tcp://172.68.195.1"

	et, err := EyeTrackerFromURL(etURL)

	if err != nil {
		t.Error(err)
	}

	if et.IsConnected() {
		t.Error("Shouldn't be connected yet.")
	}

	err = et.SetOption(OptionTimeout, 1)

	if err != nil {
		t.Error(err)
	}

	if et.URL() != etURL {
		t.Error("The url which created the Eyetracker doesn't match")
	}

	info, err := et.Info()

	// err = et.Connect()

	if err != nil {
		t.Error(err)
	} else {
		t.Log(info)
	}

}
*/

func TestDiscovery(t *testing.T) {
	_, err := gaze.USBTrackers()

	if err != nil {
		t.Error(err)
	}
}
