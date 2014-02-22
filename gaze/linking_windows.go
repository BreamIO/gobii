package gaze

/*
#cgo LDFLAGS: -Llib -lTobiiGazeCore64 -Wl,-rpath=XORIGIN
#include <Windows.h>
*/
import "C"

func init() {
	C.SetErrorMode(0)
}
