// +build windows

package gaze

//#include <Windows.h>
import "C"

import (
	"fmt"
	"syscall"
	//"unsafe"
)

const (

	//tobiigaze.h
	create = iota
	destroy
	connect
	disconnect
	run_event_loop
	break_event_loop
	start_tracking
	stop_tracking
	get_device_info
	get_track_box
	get_url
	is_connected
	get_error_message

	//tobiigaze_discovery.h
	list_usb_eye_trackers
	get_connected_eye_tracker

	// not a function
	lastIndex
)

const dllName = `TobiiGazeCore64.dll`

var (
	tobiigaze = make([]callable, lastIndex, lastIndex)

	tobiigazeNames = []string{
		"tobiigaze_create",
		"tobiigaze_destroy",

		"tobiigaze_connect",
		"tobiigaze_disconnect",

		"tobiigaze_run_event_loop",
		"tobiigaze_break_event_loop",

		"tobiigaze_start_tracking",
		"tobiigaze_stop_tracking",

		"tobiigaze_get_device_info",
		"tobiigaze_get_track_box",
		"tobiigaze_get_url",
		"tobiigaze_is_connected",
		"tobiigaze_get_error_message",

		"tobiigaze_list_usb_eye_trackers",
		"tobiigaze_get_connected_eye_tracker",
	}
)

type callable interface {
	Call(...uintptr) (uintptr, uintptr, error)
}

func abort(funcname string, err error) {
	panic(fmt.Sprintf("%s failed: %v", funcname, err))
}

func wInitializeSystem() {

}

func wCreateContext(something bool) (uintptr, error) {
	return 0, nil
}

func wReleaseContext(handle uintptr) error {
	return nil;
}

func init() {
	var err error

	// Hack to make Windows report
	// what caused a LoadDLL failure
	C.SetErrorMode(0)

	gaze, err := syscall.LoadDLL(dllName)

	if err != nil {
		abort("Failed to load "+dllName, err)
	}

	for i, name := range tobiigazeNames {
		tobiigaze[i], err = gaze.FindProc(name)

		if err != nil {
			abort("Loading Tobii Gaze function "+name, err)
		}
	}
}
