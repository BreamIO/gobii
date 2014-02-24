package gaze

/*
#include <stdlib.h>
#include "tobiigaze.h"
#include "tobiigaze_ext.h"
#include "tobiigaze_discovery.h"
*/
import "C"

import (
	"log"
	"unsafe"
)

// An abstraction for tobiigaze_eye_tracker
//
// An eye tracker is not some children's toy.
// It is a real piece of hardware.
// We (and Tobii) do our best here to make sure stuff works,
// but you as a programmer need to understand that exceptions
// (or the nearest equivalent) is not really an exception.
// Stuff will happen. Deal with it.
//
// With that being said.
// Good Luck!

type EyeTracker struct {
	handle *C.tobiigaze_eye_tracker
}

// This is what you are really looking for.
// The basic use of it is to first create one.
// Then you connect to it using Connect.
// At this point, defer a Close().
// Then you need to StartTracking.
// Give us a function of the GazeFunc type and we handle the rest.
func ExampleEyeTracker() {
	et, err := AnyEyeTracker()
	if err != nil {
		log.Fatalln("AHHHH!!! ERROR!!!")
		log.Fatalln(err)
	}
	et.Connect()
	defer et.Close()
	//et.StartTracking(func(data ETData) {
	//	log.Println(data)
	//})
}

// Get a C-style pointer
func (e EyeTracker) cPtr() *C.tobiigaze_eye_tracker {
	return e.handle
}

// Creates a new EyeTracker instance from the given url.
//
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
//
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

// Attempt to connect to the physical eye tracker.
//
// Returns nil if everything went fine.
// otherwise an Error.
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
//
// Implements io.Closer interface
func (e EyeTracker) Close() error {
	if e.IsConnected() {
		C.tobiigaze_disconnect(e.cPtr())
	}

	C.tobiigaze_destroy(e.cPtr())

	return nil
}

// Checks if the eye tracker has been connected.
//
// Returns true if it is connected.
// False otherwise.
func (e EyeTracker) IsConnected() bool {
	return C.tobiigaze_is_connected(e.cPtr()) == 1
}

// A settable option for the eye tracker
//
// Can be set using the "SetOption" method of EyeTrackers.
type EyeTrackerOption C.tobiigaze_option

const (
	OptionTimeout EyeTrackerOption = 0
)

func (e EyeTrackerOption) cTyp() C.tobiigaze_option {
	return (C.tobiigaze_option)(e)
}

// Allows you to set custom settings for the tracker.
// This includes, but is not limited to the possibility to set
// the timeout time before the synchronous operations timesout
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

func (e EyeTracker) StartTracking(callback GazeFunc) {
/*	var err Error

	callback := (C.tobiigaze_gaze_listener)(func (d *C.tobiigaze_gaze_data,
		ext *C.tobiigaze_gaze_data_extensions,
		userData uintptr) {

		cb(gazeDataFromC(d))
	})

	C.tobiigaze_start_tracking(e.cPtr(), &callback, err.cPtr(), nil)
	
	if err.ok() {
		return nil
	}
	
	return err*/
}

// Go level abstraction for the device_info struct.
//
// The EyeTrackerInfo can be used to query some meta-data about the tracker.
type EyeTrackerInfo C.struct_tobiigaze_device_info

// Get a C pointer.
func (e EyeTrackerInfo) cPtr() *C.struct_tobiigaze_device_info {
	return (*C.struct_tobiigaze_device_info)(&e)
}

// Returns the serial number associated with this EyeTrackerInfo object.
func (e EyeTrackerInfo) SerialNumber() string {
	return C.GoString((*C.char)(&e.cPtr().serial_number[0]))
}

// Returns the eye trackers model
func (e EyeTrackerInfo) Model() string {
	return C.GoString((*C.char)(&e.cPtr().model[0]))
}

// Returns the generation of the eye tracker.
//
// Tobii's eye trackers have been through several "generations"
// So this is basically a measure of how ancient the tracker is.
func (e EyeTrackerInfo) Generation() string {
	return C.GoString((*C.char)(&e.cPtr().generation[0]))
}

// Returns a string representing the version number of the firmware
// running on the tracker.
func (e EyeTrackerInfo) FirmwareVersion() string {
	return C.GoString((*C.char)(&e.cPtr().firmware_version[0]))
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

// This type is used for callbacks inserted into
// the EyeTracker for handeling incoming GazeData points.
type GazeFunc func(data GazeData)
