// Copyright 2014 Bream IO AB. All rights reserved.

// The gaze package contains bindings for the Tobii Technologies AB Gaze SDK.
// It is used to communicate with a Tobii branded eye tracker.
// It makes extensive use of Cgo.
// No trackers was harmed in the making of this package.
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
