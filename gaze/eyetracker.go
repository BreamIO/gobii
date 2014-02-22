package gaze

/*
#include <stdlib.h>
#include "tobiigaze.h"
#include "tobiigaze_ext.h"
#include "tobiigaze_discovery.h"
*/
import "C"

import (
	"unsafe"
)

// An abstraction for tobiigaze_eye_tracker
type EyeTracker struct {
	handle *C.tobiigaze_eye_tracker
}

// Get a C-style pointer
func (e EyeTracker) cPtr() *C.tobiigaze_eye_tracker {
	return e.handle
}

// Creates a new EyeTracker instance from the given url.
// An error will be returned if there was an error.
// The URL should have a format like:
//   "tet-tcp://XXX.XXX.XXX.XXX"
//   "tet-usb://XXX.XXX.XXX.XXX"
// ("tet" likely stands for Tobii Eye Tracker)
func EyeTrackerFromURL(url string) (*EyeTracker, error) {
	var err Error

	cUrl := C.CString(url)
	defer C.free(unsafe.Pointer(cUrl))

	et := C.tobiigaze_create(cUrl, err.cPtr())

	if !err.ok() {
		return nil, err
	}

	return &EyeTracker{et}, nil
}

// Gets the URL of any connected EyeTracker.
// Otherwise returns an error.
func AnyEyeTrackerURL() (string, error) {
	const capacity uint32 = 256
	var err Error

	url := (*C.char)(C.malloc(C.size_t(capacity)))
	defer C.free(unsafe.Pointer(url))

	C.tobiigaze_get_connected_eye_tracker((url),
		C.uint32_t(capacity),
		err.cPtr())

	if !err.ok() {
		return "", err
	}

	return C.GoString(url), nil
}

// Attempts to return any connected EyeTracker.
// Otherwise returns an error.
func AnyEyeTracker() (*EyeTracker, error) {
	url, err := AnyEyeTrackerURL()

	if err != nil {
		return nil, err
	}

	return EyeTrackerFromURL(url)
}

// Attempt to connect to the physical eyetracker.
// Blocking function which may return an error.
func (e EyeTracker) Connect() error {
	var err Error

	C.tobiigaze_connect(e.cPtr(), err.cPtr())

	if err.ok() {
		return nil
	}

	return err
}

// Closes the connection to the EyeTracker
// Implements Closer interface
func (e EyeTracker) Close() error {
	if e.IsConnected() {
		C.tobiigaze_disconnect(e.cPtr())
	}

	C.tobiigaze_destroy(e.cPtr())

	return nil
}

// Checks if the eye tracker has been connected.
func (e EyeTracker) IsConnected() bool {
	return C.tobiigaze_is_connected(e.cPtr()) == 1
}

type EyeTrackerOption C.tobiigaze_option

const (
	OptionTimeout EyeTrackerOption = 0
)

func (e EyeTrackerOption) cTyp() C.tobiigaze_option {
	return (C.tobiigaze_option)(e)
}

func (e EyeTracker) SetOption(o EyeTrackerOption, value int) error {
	var err Error

	C.tobiigaze_set_option(e.cPtr(), o.cTyp(),
		unsafe.Pointer(&value), err.cPtr())

	if err.ok() {
		return nil
	}

	return err
}

// Returns the URL of the EyeTracker, or if
// an error occurs, the empty string ("").
func (e EyeTracker) URL() string {
	var err Error

	str := C.GoString(C.tobiigaze_get_url(e.cPtr(), err.cPtr()))

	if err.ok() {
		return str
	}

	return ""
}

// Go level abstraction for the device_info struct.
type EyeTrackerInfo C.struct_tobiigaze_device_info

// Get a C pointer.
func (e EyeTrackerInfo) cPtr() *C.struct_tobiigaze_device_info {
	return (*C.struct_tobiigaze_device_info)(&e)
}

// Gets a Go style EyeTrackerInfo object or an error
func (e EyeTracker) Info() (EyeTrackerInfo, error) {
	var err Error
	var info EyeTrackerInfo

	C.tobiigaze_get_device_info(e.cPtr(), info.cPtr(), err.cPtr())

	if err.ok() {
		return info, nil
	}

	return info, err
}