// +build windows

package gobii/gaze

//#include <Windows.h>
import "C"

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (

	//tobiigaze.h
	tobiigaze_create = iota
	tobiigaze_destroy
	tobiigaze_connect
	tobiigaze_disconnect
	tobiigaze_run_event_loop
	tobiigaze_break_event_loop
	tobiigaze_start_tracking
	tobiigaze_stop_tracking
	tobiigaze_get_device_info
	tobiigaze_get_track_box
	tobiigaze_get_url
	tobiigaze_is_connected
	tobiigaze_get_error_message

	//tobiigaze_discovery.h
	tobiigaze_list_usb_eye_trackers
	tobiigaze_get_connected_eye_tracker
	
	// not a function
	lastIndex
)

const eyexName = `Tobii.EyeX.Client.dll`
const OK = "Åtgärden har slutförts."

var (
	txFunc = make([]*syscall.Proc, lastIndex, lastIndex)

	txName = []string{
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

func abort(funcname string, err error) {
	panic(fmt.Sprintf("%s failed: %v", funcname, err))
}

func wInitializeSystem() {
	//txInitializeSystem(TX_SYSTEMCOMPONENTOVERRIDEFLAG_NONE, NULL, NULL)
	ret, _, callErr := txFunc[txInitializeSystem].Call(
		txSystemComponentOverrideFlagNone,
		0, // null
		0) // null

	result := txResult(ret)

	if result != txResultOk {
		abort(txName[txInitializeSystem], callErr)
	}
	
}

func wCreateContext(something bool) (uintptr, error) {
	var handle uintptr = TxEmptyHandle;
	pointer := uintptr(unsafe.Pointer(&handle)) //**void

	ret, _, _ := txFunc[txCreateContext].Call(
		pointer,
		txTrue)

	result := txResult(ret)
	if result == txResultOk {
		return handle, nil
	}

	return 0, result
}

func wReleaseContext(handle uintptr) error {

}

func init() {

}
