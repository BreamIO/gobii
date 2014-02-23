package gaze

/*
#include "tobiigaze_data_types.h"
*/
import "C"

/*
A golang representation of a tobiigaze_gaze_data
*/
type GazeData C.struct_tobiigaze_gaze_data

func (data GazeData) cPtr() *C.struct_tobiigaze_gaze_data {
	return (*C.struct_tobiigaze_gaze_data)(&data)
}

func (data GazeData) Timestamp() uint64 {
	return uint64((C.uint64_t)(data.cPtr().timestamp))
}

func (data GazeData) TrackingStatus() TrackingStatus {
	return TrackingStatus((C.uint64_t)(data.cPtr().tracking_status))
}

func (data GazeData) Left() EyeData {
	return EyeData((C.struct_tobiigaze_gaze_data_eye)(data.cPtr().left))
}

func (data GazeData) Right() EyeData {
	return EyeData((C.struct_tobiigaze_gaze_data_eye)(data.cPtr().right))
}

type EyeData C.struct_tobiigaze_gaze_data_eye

func (ed EyeData) cPtr() *C.struct_tobiigaze_gaze_data_eye {
	return (*C.struct_tobiigaze_gaze_data_eye)(&ed)
}

// Gets the position of the users eyes in millimetres.
//
// Point is relative to the tracker.
func (ed EyeData) EyePositionFromEyeTracker() Point3D {
	return Point3D((C.struct_tobiigaze_point_3d)(ed.cPtr().eye_position_from_eye_tracker_mm))
}

// Gets the position of the users eyes in normalized coordinates.
//
// Point is relative to the track box.
func (ed EyeData) EyePositionInTrackBox() Point3D {
	return Point3D((C.struct_tobiigaze_point_3d)(ed.cPtr().eye_position_in_track_box_normalized))
}

// Gets the position of the gaze point in millimetres.
//
// Point is relative to the tracker.
func (ed EyeData) GazePointFromEyeTracker() Point3D {
	return Point3D((C.struct_tobiigaze_point_3d)(ed.cPtr().gaze_point_from_eye_tracker_mm))
}

// Gets the position of the gaze point in normalized coordinates.
//
// The point is relative to the upper left corner of the screen.
func (ed EyeData) GazePointOnDisplay() Point2D {
	return Point2D((C.struct_tobiigaze_point_2d)(ed.cPtr().gaze_point_on_display_normalized))
}

// Golang name of the Tobii Gaze SDK point_3d struct
//
// It contains three float64s (double),
// one for each axis in a three dimensional
// carthesian coordinate system (x, y, z).
type Point3D C.struct_tobiigaze_point_3d

func (p Point3D) cPtr() *C.struct_tobiigaze_point_3d {
	return (*C.struct_tobiigaze_point_3d)(&p)
}

func (p Point3D) X() float64 {
	return float64((C.double)(p.cPtr().x))
}

func (p Point3D) Y() float64 {
	return float64((C.double)(p.cPtr().y))
}

func (p Point3D) Z() float64 {
	return float64((C.double)(p.cPtr().z))
}

// Golang name of the Tobii Gaze SDK point_3d struct
//
// It contains two float64s (double),
// one for each axis in a two dimensional
// carthesian coordinate system (x, y).
type Point2D C.struct_tobiigaze_point_2d

func (p Point2D) cPtr() *C.struct_tobiigaze_point_2d {
	return (*C.struct_tobiigaze_point_2d)(&p)
}

func (p Point2D) X() float64 {
	return float64((C.double)(p.cPtr().x))
}

func (p Point2D) Y() float64 {
	return float64((C.double)(p.cPtr().y))
}

type TrackingStatus int

// The possible gaze tracking statuses.
// Taken from tobiigaze_data_types.h
const (
	NoEyesTracked = TrackingStatus(iota)
	BothEyesTracked
	OnlyLeftEyeTracked
	OneEyeTrackedProbablyLeft
	OneEyeTrackedUnknownWhich
	OneEyeTrackedProbablyRight
	OnlyRightEye_Tracked
)
