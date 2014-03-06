package gaze_test

import (
	"github.com/zephyyrr/gobii/gaze"
	"testing"
)

var errors = []int{
	0, 1, 2, 3, 4,
	100, 101,
	200, 201, 202, 203, 204, 205,
	300, 301,

	0x20000500,
	0x20000501,
	0x20000502,
	0x20000503,
	0x20000504,
	0x20000505,
	0x20000506,
	0x20000507,
	0x20000508,
	0x20000509,
	0x2000050A,
}

var errorStrings = []string{
	"Operation successful.",
	"Unknown error.",
	"Out of memory.",
	"Buffer allocated by client too small.",
	"Input parameter is invalid",
	"Operation timed out.",
	"Operation aborted.",
	"Invalid or badly formatted url.",
	"Host name lookup failed for eye tracker url.",
	"Connection to eye tracker failed.",
	"Communication with eye tracker failed.",
	"Already connected",
	"Not connected",
	"Protocol error.",
	"Protocol version mismatch",
	"Firmware error: Unknown operation.",
	"Firmware error: Operation not supported by eye tracker.",
	"Firmware error: Operation failed.",
	"Firmware error: Protocol payload error.",
	"Firmware error: Protocol unknown id error.",
	"Firmware error: Authorization failed.",
	"Firmware error: Extension required.",
	"Firmware error: Internal error.",
	"Firmware error: State error.",
	"Firmware error: Invalid parameter.",
	"Firmware error: Operation aborted.",
}

func TestError(t *testing.T) {
	if !gaze.Error(0).Ok() {
		t.Fatal("Error code 0 should be an error.")
	}

	for i, e := range errors {
		if gaze.Error(e).Error() != errorStrings[i] {
			t.Fatalf("Expected '%s' to be '%s'\n",
				gaze.Error(e), errorStrings[i])
		}
	}
}
