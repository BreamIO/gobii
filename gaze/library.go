package gaze

/*
#include "tobiigaze.h"
*/
import "C"

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
