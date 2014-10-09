package gaze

//#include "tobiigaze_data_types.h"
import "C"

type CalibrationPoint struct {
}

func CalibrationPointFromC(data C.struct_tobiigaze_calibration_point_data) CalibrationPoint {
	return CalibrationPoint{}
}
