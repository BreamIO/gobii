package gaze

import "testing"

func TestVersion(t *testing.T) {
	if Version() != "..." {
		t.Fatal(Version())
	}
}

func TestCreateEyetracker(t *testing.T) {
	const etURL = "tet-tcp://172.68.195.1"

	et, err := CreateEyeTracker(etURL)

	if err != nil {
		t.Fatal(err)
	}

	if et.IsConnected() {
		t.Fatal("Shouldn't be connected yet.")
	}

	err = et.SetOption(OptionTimeout, 1)

	if err != nil {
		t.Fatal(err)
	}

	if et.URL() != etURL {
		t.Fatal("The url which created the Eyetracker doesn't match")
	}

	info, err := et.Info()

	// err = et.Connect()

	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(info)
	}

}

func TestDiscovery(t *testing.T) {
	err := ListTrackers()

	if err != nil {
		t.Fatal(err)
	}
}
