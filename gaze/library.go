package gaze

/*
#include <stdlib.h>
#include "tobiigaze.h"
#include "tobiigaze_ext.h"
#include "tobiigaze_discovery.h"

size_t getInfoSize() {
	return sizeof(struct usb_device_info);
}
*/
import "C"

import (
	"fmt"
	"log"
	"reflect"
	"unsafe"
)

// The version of the library on the form "1.0.2".
// Is currently "..." on Linux
func Version() string {
	return C.GoString(C.tobiigaze_get_version())
}

// Probably won't be used, but demonstrated
// the general idea behind the interface.
type abstractedCType interface {
	cPtr() *interface{}
	cTyp() interface{}
}

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
	log.Println(url)
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

type USBInfo C.struct_usb_device_info

func (info USBInfo) cPtr() *C.struct_usb_device_info {
	return (*C.struct_usb_device_info)(&info)
}

func (info USBInfo) SerialNumber() string {
	return C.GoString((*C.char)(&info.cPtr().serialNumber[0]))
}

func (info USBInfo) ProductName() string {
	return C.GoString((*C.char)(&info.cPtr().productName[0]))
}

func (info USBInfo) PlatformType() string {
	return C.GoString((*C.char)(&info.cPtr().platformType[0]))
}

func (info USBInfo) FirmwareVersion() string {
	return C.GoString((*C.char)(&info.cPtr().firmwareVersion[0]))
}

func (info USBInfo) String() string {
	return fmt.Sprintf("\tName: %s\n\tSerialnumber: %s\n\tPlatform: %s\n\tFirmware: %s", 
		info.ProductName(), 
		info.SerialNumber(),
		info.PlatformType(),
		info.FirmwareVersion())
}

func USBTrackers() ([]USBInfo, error) {
	var err Error
	capacity := 10
	var length uint32
	
	infos_void := C.malloc(C.getInfoSize() * C.size_t(capacity))
	defer C.free(infos_void)
	infos := (*C.struct_usb_device_info)(infos_void)
	
	C.tobiigaze_list_usb_eye_trackers(infos, 
		C.uint32_t(capacity), 
		(*C.uint32_t)(&length),
		err.cPtr())
		
	var goInfos []USBInfo
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&goInfos)))
	sliceHeader.Cap = int(length)
	sliceHeader.Len = int(length)
	sliceHeader.Data = uintptr(unsafe.Pointer(infos))

	if length == 0 {
		return nil, err
	}
	
	//Transfer data into Go runtime handled memory.
	result := make([]USBInfo, length, length)
	copy(result, goInfos)
	return result, nil;
}

func ListUSBTrackers() error {
	list, err := USBTrackers()
	if err != nil {
		return err;
	}
	for i, info := range list {
		fmt.Println(i, ": ", info)
	}
	return err
}
