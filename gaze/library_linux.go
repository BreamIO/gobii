package gaze

/*
#cgo LDFLAGS: -L/home/victorystick/go/bin -ltobiigazecore -Wl,-rpath=/home/victorystick/go/bin
#include <stdlib.h>
#include "tobiigaze.h"
*/
import "C"

type EyeTracker struct {
	handle *C.tobiigaze_eye_tracker
}

func createEyeTracker(url string) (*EyeTracker, error) {
	var err Error

	cUrl := C.CString(url)

	et := C.tobiigaze_create(cUrl, err.cPtr())

	if !err.ok() {
		return nil, err
	}

	return &EyeTracker{et}, nil
}