package gaze

/*
#include <stdlib.h>
#include <stdio.h>
#include "tobiigaze.h"
#include "tobiigaze_ext.h"
#include "tobiigaze_data_types.h"
#include "tobiigaze_discovery.h"
#include "tobiigaze_calibration.h"
#include "tobiigaze_display_area.h"
#include "callbacks.h"
*/
import "C"

import (
	"fmt"
	"sync"
	"unsafe"
)

type EyeTracker interface {
	Connect() error
	Close() error
	IsConnected() bool
	SetOption(EyeTrackerOption, int) error
	URL() string
	StartTracking(GazeFunc) error

	StartCalibration(Callback)
	StopCalibration(Callback)
	AddPointToCalibration(*Point2D, Callback)
	RemovePointFromCalibration(*Point2D, Callback)
	ComputeAndSetCalibration(Callback)
	CalibrationPoints() ([]CalibrationPoint, error)
	SetDisplayArea(width, height, angle float64) error
}

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

type GazeTracker struct {
	handle              *C.tobiigaze_eye_tracker
	gazeCallback        GazeFunc
	calibrationLock     sync.Mutex
	calibrationCallback Callback
}

// Get a C-style pointer
func (e GazeTracker) cPtr() *C.tobiigaze_eye_tracker {
	return e.handle
}

// Get a C-style pointer
func (e GazeTracker) CPtr() *C.tobiigaze_eye_tracker {
	return e.cPtr()
}

// Creates a new GazeTracker instance from the given url.
//
// An error will be returned if there was an error.
// The URL should have a format like:
//   "tet-tcp://XXX.XXX.XXX.XXX"
//   "tet-usb://XXX.XXX.XXX.XXX"
// ("tet" likely stands for Tobii Eye Tracker)
func EyeTrackerFromURL(url string) (EyeTracker, error) {
	var err Error

	cUrl := C.CString(url)
	defer C.free(unsafe.Pointer(cUrl))

	et := C.tobiigaze_create(cUrl, err.cPtr())

	if !err.Ok() {
		return nil, err
	}
	return &GazeTracker{et, nil, sync.Mutex{}, nil}, nil
}

// Gets the URL of any connected GazeTracker.
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

	if !err.Ok() {
		return "", err
	}

	return C.GoString(url), nil
}

// Attempts to return any connected GazeTracker.
// Otherwise returns an error.
func AnyEyeTracker() (EyeTracker, error) {
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
func (e *GazeTracker) Connect() error {
	var err Error

	if e.IsConnected() {
		return Error(C.TOBIIGAZE_ERROR_ALREADY_CONNECTED)
	}

	go func() {
		var err Error
		C.tobiigaze_run_event_loop(e.cPtr(), err.cPtr())
		//log.Println("Gobii/Tobii Event loop has terminated:", err)
	}()

	C.tobiigaze_connect(e.cPtr(), err.cPtr())

	if !err.Ok() {
		return err
	}

	return nil
}

// Closes the connection to the GazeTracker
//
// Implements io.Closer interface
func (e *GazeTracker) Close() error {
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
func (e *GazeTracker) IsConnected() bool {
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
func (e *GazeTracker) SetOption(o EyeTrackerOption, value int) error {
	var err Error

	C.tobiigaze_set_option(e.cPtr(), o.cTyp(),
		unsafe.Pointer(&value), err.cPtr())

	if err.Ok() {
		return nil
	}

	return err
}

// Returns the URL of the GazeTracker, or if
// an error occurs, the empty string ("").
func (e *GazeTracker) URL() string {
	var err Error

	str := C.GoString(C.tobiigaze_get_url(e.cPtr(), err.cPtr()))

	if err.Ok() {
		return str
	}

	return ""
}

//export exportedTrackingCallback
// Callback reused by all trackers. Warning! Dark magic!
//
// data has to be copied to be persisted.
// ext is currently not used.
// userData will always be a pointer to an GazeTracker instance.
func exportedTrackingCallback(data *C.struct_tobiigaze_gaze_data,
	ext *C.struct_tobiigaze_gaze_data_extension,
	userData unsafe.Pointer) {

	et := (*GazeTracker)(unsafe.Pointer(userData))
	//fmt.Println(GazeDataFromC(data))
	if et.gazeCallback != nil {
		et.gazeCallback(GazeDataFromC(data))
	}
}

// The callback parameter is now silently ignored.
func (e *GazeTracker) StartTracking(callback GazeFunc) error {
	var err Error
	e.gazeCallback = callback
	C.tobiigaze_start_tracking(e.cPtr(),
		//(C.tobiigaze_gaze_listener)(unsafe.Pointer(&trackingCallback)),
		C.breamio_get_listener(),
		err.cPtr(),
		unsafe.Pointer(e)) //unsafe reference to the tracker. Needed to get the real callback later

	if err.Ok() {
		return nil
	}

	return err
}

//export exportedCalibrationCallback
// Callback reused by all trackers. Warning! Dark magic!
//
// data has to be copied to be persisted.
// ext is currently not used.
// userData will always be a pointer to an GazeTracker instance.
func exportedCalibrationCallback(error_code C.tobiigaze_error_code,
	userData unsafe.Pointer) {

	et := (*GazeTracker)(unsafe.Pointer(userData))
	//fmt.Println(GazeDataFromC(data))
	callback := et.calibrationCallback
	et.calibrationLock.Unlock()
	if callback != nil {
		err := (Error)(error_code)
		if err.Ok() {
			go callback(nil)
		} else {
			go callback(err)
		}

	}
	et.calibrationCallback = nil
}

// Starts calibration of the eye tracker.
//
// If another calibration call is in the pipeline, this call will block until it is done.
// This is to protect user data from corruption and garbage collection.
// Please understand.
func (e *GazeTracker) StartCalibration(callback Callback) {
	e.calibrationLock.Lock()
	e.calibrationCallback = callback //Store it to avoid GC sweep.
	C.tobiigaze_calibration_start_async(e.cPtr(),
		C.breamio_get_callback(),
		unsafe.Pointer(e)) //unsafe reference to the tracker. Needed to get the real callback later
}

// Stops calibration of the eye tracker.
//
// If another calibration call is in the pipeline, this call will block until it is done.
// This is to protect user data from corruption and garbage collection.
// Please understand.
func (e *GazeTracker) StopCalibration(callback Callback) {
	e.calibrationLock.Lock()
	e.calibrationCallback = callback //Store it to avoid GC sweep.
	C.tobiigaze_calibration_stop_async(e.cPtr(),
		C.breamio_get_callback(),
		unsafe.Pointer(e)) //unsafe reference to the tracker. Needed to get the real callback later
}

// Adds a point to the current calibration.
//
// The callback is called with a Error if called without a prior successful StartCalibration call.
//
// If another calibration call is in the pipeline, this call will block until it is done.
// This is to protect user data from corruption and garbage collection.
// Please understand.
func (e *GazeTracker) AddPointToCalibration(point *Point2D, callback Callback) {
	e.calibrationLock.Lock()
	e.calibrationCallback = callback //Store it to avoid GC sweep.
	C.tobiigaze_calibration_add_point_async(e.cPtr(),
		&C.struct_tobiigaze_point_2d{(C.double)(point.x), (C.double)(point.y)},
		C.breamio_get_callback(),
		unsafe.Pointer(e)) //unsafe reference to the tracker. Needed to get the real callback later
}

// Removes a point from the current calibration.
//
// The callback is called with a Error if called without a prior successful StartCalibration call.
//
// If another calibration call is in the pipeline, this call will block until it is done.
// This is to protect user data from corruption and garbage collection.
// Please understand.
func (e *GazeTracker) RemovePointFromCalibration(point *Point2D, callback Callback) {
	e.calibrationLock.Lock()
	e.calibrationCallback = callback //Store it to avoid GC sweep.
	C.tobiigaze_calibration_remove_point_async(e.cPtr(),
		&C.struct_tobiigaze_point_2d{(C.double)(point.x), (C.double)(point.y)},
		C.breamio_get_callback(),
		unsafe.Pointer(e)) //unsafe reference to the tracker. Needed to get the real callback later
}

// Calculates the result of the calibration and stores it in the tracker.
//
// The callback is called with a Error if called without a prior successful StartCalibration call.
//
// If another calibration call is in the pipeline, this call will block until it is done.
// This is to protect user data from corruption and garbage collection.
// Please understand.
func (e *GazeTracker) ComputeAndSetCalibration(callback Callback) {
	e.calibrationLock.Lock()
	e.calibrationCallback = callback //Store it to avoid GC sweep.
	C.tobiigaze_calibration_compute_and_set_async(e.cPtr(),
		C.breamio_get_callback(),
		unsafe.Pointer(e)) //unsafe reference to the tracker. Needed to get the real callback later
}

func (e GazeTracker) String() string {
	return fmt.Sprintf("<gaze.GazeTracker %x>", e.handle)
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

// Returns a print-friendly string representation of the info.
func (e EyeTrackerInfo) String() string {
	return fmt.Sprintf("\tModel: %s\n\tSerialnumber: %s\n\tGeneration: %s\n\tFirmware: %s",
		e.Model(),
		e.SerialNumber(),
		e.Generation(),
		e.FirmwareVersion())
}

// Gets a Go style EyeTrackerInfo object or an error
func (e GazeTracker) Info() (EyeTrackerInfo, error) {
	var err Error
	var info EyeTrackerInfo

	C.tobiigaze_get_device_info(e.cPtr(), info.cPtr(), err.cPtr())

	if err.Ok() {
		return info, nil
	}

	return info, err
}

func (e GazeTracker) CalibrationPoints() ([]CalibrationPoint, error) {
	calibration := new(C.struct_tobiigaze_calibration)
	err := new(Error)
	C.tobiigaze_get_calibration(e.cPtr(), calibration, err.cPtr())
	if !err.Ok() {
		return nil, err
	}
	var points [tobiigaze_max_calibration_point_data_items]C.struct_tobiigaze_calibration_point_data
	var resultLength uint32
	c_reslength := (C.uint32_t)(resultLength)
	C.tobiigaze_get_calibration_point_data_items(calibration, &points[0], (C.uint32_t)(len(points)), &c_reslength, err.cPtr())
	if !err.Ok() {
		return nil, err
	}

	result := make([]CalibrationPoint, resultLength, resultLength)

	for i := uint32(0); i < resultLength; i++ {
		result[i] = CalibrationPointFromC(points[i])
	}

	return result, nil
}

func (e GazeTracker) SetDisplayArea(width, height, angle float64) error {
	var disp_area C.struct_tobiigaze_display_area
	disp_area = getDisplayArea(width, height, angle)

	err := new(Error)
	C.tobiigaze_set_display_area(e.cPtr(), &disp_area, err.cPtr())
	if err.Ok() {
		return nil
	}
	return err
}

func getDisplayArea(width, height, angle float64) C.struct_tobiigaze_display_area {
	var disp_area C.struct_tobiigaze_display_area

	/*disp_area.top_left.x = (C.double)(-width / 2)
	disp_area.top_left.y = (C.double)(math.Cos(angle) * (height + 30))
	disp_area.top_left.z = -(C.double)(math.Sin(angle)*(height+10) + 10)

	disp_area.top_right.x = (C.double)(width / 2)
	disp_area.top_right.y = (C.double)(math.Cos(angle) * (height + 30))
	disp_area.top_right.z = -(C.double)(math.Sin(angle)*(height+10) + 10)

	disp_area.bottom_left.x = (C.double)(-width / 2)
	disp_area.bottom_left.y = (C.double)(30)
	disp_area.bottom_left.z = (C.double)(-10)*/

	// Dreamhack values of Tracker #1
	disp_area.top_left.x = -263
	disp_area.top_left.y = 323
	disp_area.top_left.z = -25

	disp_area.top_right.x = 260
	disp_area.top_right.y = 323
	disp_area.top_right.z = -25

	disp_area.bottom_left.x = -263
	disp_area.bottom_left.y = 25
	disp_area.bottom_left.z = -10

	return disp_area
}

// This type is used for callbacks inserted into
// the GazeTracker for handeling incoming GazeData points.
type GazeFunc func(data *GazeData)

//Common tobii callback
//Used for most calibration functions.
type Callback func(err error)
