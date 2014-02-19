package gaze

/*
#cgo LDFLAGS: -L/home/victorystick/go/bin -ltobiigazecore -Wl,-rpath=/home/victorystick/go/bin
#include "tobiigaze.h"
*/
import "C"

import (
	"fmt"
)

type Error C.tobiigaze_error_code

func (e Error) Error() string {
	return C.GoString(C.tobiigaze_get_error_message(
		C.tobiigaze_error_code(e)))
}
