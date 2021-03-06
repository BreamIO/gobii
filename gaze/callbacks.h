#pragma once

#include "tobiigaze_data_types.h"
#include "tobiigaze_ext.h"

void
breamio_listener(const struct tobiigaze_gaze_data *gaze_data, const struct tobiigaze_gaze_data_extension *gaze_data_extension, void *user_data);

void
breamio_calibration_listener(tobiigaze_error_code error_code, void *user_data);

tobiigaze_gaze_listener 
breamio_get_listener();

tobiigaze_async_callback 
breamio_get_callback();

