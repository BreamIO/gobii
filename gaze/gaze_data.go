package gaze

/*
#include <stdlib.h>
#include "tobiigaze_data_types.h"

size_t sizeofGazedata() {
	return sizeof(struct tobiigaze_gaze_data);
}
*/
import "C"

import (
	"fmt"
	"time"
	"unsafe"
)

/*
A golang representation of a tobiigaze_gaze_data
*/
type GazeData struct {
	timestamp      time.Time
	trackingstatus TrackingStatus
	left           EyeData
	right          EyeData
}

func GazeDataFromC(c_data *C.struct_tobiigaze_gaze_data) (data *GazeData) {
	data = new(GazeData)
	data.timestamp = time.Unix(int64((C.int64_t)(c_data.timestamp)), 0)
	data.trackingstatus = TrackingStatus((C.uint64_t)(c_data.tracking_status))
	data.left = eyeDataFromC((C.struct_tobiigaze_gaze_data_eye)(c_data.left))
	data.right = eyeDataFromC((C.struct_tobiigaze_gaze_data_eye)(c_data.right))
	return
}

func (data GazeData) String() string {
	return fmt.Sprintf("Left: %s | Right: %s | Timestamp: %d", data.left, data.right, data.timestamp)
}

func (data GazeData) Timestamp() time.Time {
	return data.timestamp
}

func (data GazeData) TrackingStatus() TrackingStatus {
	return data.trackingstatus
}

func (data GazeData) Left() EyeData {
	return data.left
}

func (data GazeData) Right() EyeData {
	return data.right
}

type EyeData struct {
	eyePositionFromEyeTracker,
	eyePositionInTrackBox,
	gazePointFromEyeTracker Point3D
	gazePointOnDisplay Point2D
}

func eyeDataFromC(c_data C.struct_tobiigaze_gaze_data_eye) (ed EyeData) {
	ed.eyePositionFromEyeTracker = point3DFromC((c_data.eye_position_from_eye_tracker_mm))
	ed.eyePositionInTrackBox = point3DFromC(c_data.eye_position_in_track_box_normalized)
	ed.gazePointFromEyeTracker = point3DFromC(c_data.gaze_point_from_eye_tracker_mm)
	ed.gazePointOnDisplay = point2DFromC(c_data.gaze_point_on_display_normalized)
	return
}

func (ed EyeData) String() string {
	return fmt.Sprintf("(%f, %f)", ed.gazePointOnDisplay.x, ed.gazePointOnDisplay.y)
}

// Gets the position of the users eyes in millimetres.
//
// Point is relative to the tracker.
func (ed EyeData) EyePositionFromEyeTracker() Point3D {
	return ed.eyePositionFromEyeTracker
}

// Gets the position of the users eyes in normalized coordinates.
//
// Point is relative to the track box.
func (ed EyeData) EyePositionInTrackBox() Point3D {
	return ed.eyePositionInTrackBox
}

// Gets the position of the gaze point in millimetres.
//
// Point is relative to the tracker.
func (ed EyeData) GazePointFromEyeTracker() Point3D {
	return ed.gazePointFromEyeTracker
}

// Gets the position of the gaze point in normalized coordinates.
//
// The point is relative to the upper left corner of the screen.
func (ed EyeData) GazePointOnDisplay() Point2D {
	return ed.gazePointOnDisplay
}

// Golang name of the Tobii Gaze SDK point_3d struct
//
// It contains three float64s (double),
// one for each axis in a three dimensional
// carthesian coordinate system (x, y, z).
type Point3D struct {
	Point2D
	z float64
}

func point3DFromC(c_data C.struct_tobiigaze_point_3d) (p Point3D) {
	p.Point2D.x = float64(c_data.x)
	p.Point2D.y = float64(c_data.y)
	p.z = float64(c_data.z)
	return
}

func (p Point3D) Z() float64 {
	return p.z
}

// Golang name of the Tobii Gaze SDK point_3d struct
//
// It contains two float64s (double),
// one for each axis in a two dimensional
// carthesian coordinate system (x, y).
type Point2D struct {
	x, y float64
}

func point2DFromC(c_data C.struct_tobiigaze_point_2d) (p Point2D) {
	p.x = float64(c_data.x)
	p.y = float64(c_data.y)
	return
}

func (p Point2D) X() float64 {
	return p.x
}

func (p Point2D) Y() float64 {
	return p.y
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


// This function can be used to unit test this file.
// Go test tool does not allow Cgo in the tests, but since a c struct is needed, we will simply create it here, 
// and return it for direct use in the GazeDataFromC function.
// Do not forget to free it after use by calling 
// 		FreeTestStruct(c_data)
func GenTestStruct() *C.struct_tobiigaze_gaze_data {
	c_data := (*C.struct_tobiigaze_gaze_data)(C.malloc(C.sizeofGazedata()))
	
	c_data.timestamp = (C.uint64_t)(1337)
	c_data.tracking_status = (C.tobiigaze_tracking_status)(BothEyesTracked)
	
	//Left
	c_data.left.eye_position_from_eye_tracker_mm.x = 3.2552
	c_data.left.eye_position_from_eye_tracker_mm.y = 4.5342
	c_data.left.eye_position_from_eye_tracker_mm.z = 5.75342
	
	c_data.left.eye_position_in_track_box_normalized.x = 0.25342
	c_data.left.eye_position_in_track_box_normalized.y = 0.5324
	c_data.left.eye_position_in_track_box_normalized.z  = 0.46546
	
	c_data.left.gaze_point_from_eye_tracker_mm.x = 0.25145
	c_data.left.gaze_point_from_eye_tracker_mm.y = 32.54654
	c_data.left.gaze_point_from_eye_tracker_mm.z = 2.75876
	
	c_data.left.gaze_point_on_display_normalized.x = 0.12123
	c_data.left.gaze_point_on_display_normalized.y = 0.45745
	
	//Right
	c_data.right.eye_position_from_eye_tracker_mm.x = 4.2552
	c_data.right.eye_position_from_eye_tracker_mm.y = 5.5342
	c_data.right.eye_position_from_eye_tracker_mm.z = 6.75342
	
	c_data.right.eye_position_in_track_box_normalized.x = 0.35342
	c_data.right.eye_position_in_track_box_normalized.y = 0.4324
	c_data.right.eye_position_in_track_box_normalized.z  = 0.16546
	
	c_data.right.gaze_point_from_eye_tracker_mm.x = 0.45145
	c_data.right.gaze_point_from_eye_tracker_mm.y = 35.54654
	c_data.right.gaze_point_from_eye_tracker_mm.z = 2.85876
	
	c_data.right.gaze_point_on_display_normalized.x = 0.92123
	c_data.right.gaze_point_on_display_normalized.y = 0.65745
	
	return c_data
}

//Function to free a struct given by 
//		GenTestStruct()
func FreeTestStruct(c_data *C.struct_tobiigaze_gaze_data) {
	C.free(unsafe.Pointer(c_data))
}