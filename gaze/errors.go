package gaze

/*
#cgo LDFLAGS: -L/home/victorystick/go/bin -ltobiigazecore -Wl,-rpath=/home/victorystick/go/bin
#include "tobiigaze.h"
#include "tobiigaze_error_codes.h"
*/
import "C"

type Error C.tobiigaze_error_code

func (e Error) Error() string {
	return C.GoString(C.tobiigaze_get_error_message(
		C.tobiigaze_error_code(e)))
}

func (e *Error) cPtr() *C.tobiigaze_error_code {
	return (*C.tobiigaze_error_code)(e)
}

func (e Error) ok() bool {
	return e == C.TOBIIGAZE_ERROR_SUCCESS
}
