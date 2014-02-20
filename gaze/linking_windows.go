package gaze

/*
#cgo LDFLAGS: -Llib -lTobiiGazeCore64 -Wl,-rpath=XORIGIN
*/
import "C"

func init() {
	C.SetErrorMode(0)
}