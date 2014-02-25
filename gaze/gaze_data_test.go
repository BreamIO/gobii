package gaze

import (
	"testing"
)

func TestGazeFromC(t *testing.T) {
	_struct := genTestStruct()
	defer freeTestStruct(_struct)
	data := GazeDataFromC(_struct)
	
	//Check timestamp
	if data.Timestamp().Unix() != 1337 {
		t.Errorf("Timestamp was %d, expected %d", data.Timestamp().Unix(), 1337)
	}
	
	//Check status
	if data.TrackingStatus() != BothEyesTracked {
		t.Errorf("TrackingStatus was %d, expected %d", data.TrackingStatus(), BothEyesTracked)
	}
}

func TestLeftFromC(t *testing.T) {

	/* Testing
		c_data.left.eyePositionFromEyeTracker.x = 3.2552
		c_data.left.eyePositionFromEyeTracker.y = 4.5342
		c_data.left.eyePositionFromEyeTracker.z = 5.75342
		
		c_data.left.eyePositionInTrackBox.x = 0.25342
		c_data.left.eyePositionInTrackBox.y = 0.5324
		c_data.left.eyePositionInTrackBox.z  = 0.46546
		
		c_data.left.gazePointFromEyeTracker.x = 0.25145
		c_data.left.gazePointFromEyeTracker.y = 32.54654
		c_data.left.gazePointFromEyeTracker.z = 2.75876
		
		c_data.left.gazePointOnDisplay.x = 0.12123
		c_data.left.gazePointOnDisplay.y = 0.45745
	*/

	_struct := genTestStruct()
	defer freeTestStruct(_struct)
	left := GazeDataFromC(_struct).Left()
	
	epfet := left.EyePositionFromEyeTracker()
	assertEqFloat64(t, epfet.X(), 3.2552, "left.EyePositionFromEyeTracker().X()")
	assertEqFloat64(t, epfet.Y(), 4.5342, "left.EyePositionFromEyeTracker().X()")
	assertEqFloat64(t, epfet.Z(), 5.75342, "left.EyePositionFromEyeTracker().X()")
	
	epitb := left.EyePositionInTrackBox()
	assertEqFloat64(t, epitb.X(), 0.25342, "left.EyePositionInTrackBox().X()")
	assertEqFloat64(t, epitb.Y(), 0.5324, "left.EyePositionInTrackBox().X()")
	assertEqFloat64(t, epitb.Z(), 0.46546, "left.EyePositionInTrackBox().X()")
	
	gpfet := left.GazePointFromEyeTracker()
	assertEqFloat64(t, gpfet.X(), 0.25145, "left.GazePointFromEyeTracker().X()")
	assertEqFloat64(t, gpfet.Y(), 32.54654, "left.GazePointFromEyeTracker().X()")
	assertEqFloat64(t, gpfet.Z(), 2.75876, "left.GazePointFromEyeTracker().X()")
	
	gpod := left.GazePointOnDisplay()
	assertEqFloat64(t, gpod.X(), 0.12123, "left.GazePointOnDisplay().X()")
	assertEqFloat64(t, gpod.Y(), 0.45745, "left.GazePointOnDisplay().X()")
}

func TestRightFromC(t *testing.T) {

	/* Testing
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
*/

	_struct := genTestStruct()
	defer freeTestStruct(_struct)
	right := GazeDataFromC(_struct).Right()
	
	epfet := right.EyePositionFromEyeTracker()
	assertEqFloat64(t, epfet.X(), 4.2552, "right.EyePositionFromEyeTracker().X()")
	assertEqFloat64(t, epfet.Y(), 5.5342, "right.EyePositionFromEyeTracker().X()")
	assertEqFloat64(t, epfet.Z(), 6.75342, "right.EyePositionFromEyeTracker().X()")
	
	epitb := right.EyePositionInTrackBox()
	assertEqFloat64(t, epitb.X(), 0.35342, "right.EyePositionInTrackBox().X()")
	assertEqFloat64(t, epitb.Y(), 0.4324, "right.EyePositionInTrackBox().X()")
	assertEqFloat64(t, epitb.Z(), 0.16546, "right.EyePositionInTrackBox().X()")
	
	gpfet := right.GazePointFromEyeTracker()
	assertEqFloat64(t, gpfet.X(), 0.45145, "right.GazePointFromEyeTracker().X()")
	assertEqFloat64(t, gpfet.Y(), 35.54654, "right.GazePointFromEyeTracker().X()")
	assertEqFloat64(t, gpfet.Z(), 2.85876, "right.GazePointFromEyeTracker().X()")
	
	gpod := right.GazePointOnDisplay()
	assertEqFloat64(t, gpod.X(), 0.92123, "right.GazePointOnDisplay().X()")
	assertEqFloat64(t, gpod.Y(), 0.65745, "right.GazePointOnDisplay().X()")
}



func assertEqFloat64(t *testing.T, value, expected float64, name string) {
	if value != expected {
		t.Errorf("%s was %d, expected %d", name, value, expected)
	}
}