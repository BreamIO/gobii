#include "callbacks.h"

void
breamio_listener(const struct tobiigaze_gaze_data *gaze_data, 
	const struct tobiigaze_gaze_data_extension *gaze_data_extension, 
	void *user_data)
{
	exportedTrackingCallback(gaze_data, gaze_data_extension, user_data);
}

tobiigaze_gaze_listener 
breamio_get_listener()
{
	return (tobiigaze_gaze_listener) &breamio_listener;
}

void
breamio_calibration_listener(tobiigaze_error_code error_code, void *user_data)
{
	exportedCalibrationCallback(error_code, user_data);
}

tobiigaze_async_callback 
breamio_get_callback()
{
	return (tobiigaze_async_callback) &breamio_calibration_listener;
}