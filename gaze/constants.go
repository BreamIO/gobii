package gaze

type txUserParam uintptr
type txHandle uintptr
type txTicket int

type txResult int

// tobiigaze_log_level
// Values for setting the log level in the Tobii Gaze SDK
const (
	tobiigaze_log_level_off     = iota // 0
	tobiigaze_log_level_debug          // 1
	tobiigaze_log_level_info           // 2
	tobiigaze_log_level_warning        // 3
	tobiigaze_log_level_error          // 4
)

// Various tobiigaze constants
// Taken from tobiigaze_data_types.h
const (
	tobiigaze_device_info_max_serial_number_length = 128
	tobiigaze_device_info_max_model_length         = 64
	tobiigaze_device_info_max_generation_length    = 64
	tobiigaze_device_info_max_firmware_length      = 128
	tobiigaze_calibration_data_capacity            = 4 * 1024 * 1024
	tobiigaze_key_size                             = 32
	tobiigaze_max_calibration_point_data_items     = 512
	tobiigaze_device_info_max_size                 = 128
	tobiigaze_device_address_info_max_size         = 138
	tobiigaze_max_devices                          = 9 // Fo Realz?
	tobiigaze_framerates_max_size                  = 32
	tobiigaze_illumination_mode_string_max_size    = 64
	tobiigaze_illumination_modes_max_size          = 16
	tobiigaze_unit_name_max_size                   = 64
	tobiigaze_extension_name_max_size              = 16
	tobiigaze_extensions_max_size                  = 16
	tobiigaze_max_wake_on_gaze_regions             = 4
	tobiigaze_authorize_challenge_max_len          = 512
	tobiigaze_max_gaze_data_extensions             = 32
	tobiigaze_max_gaze_data_extension_length       = 256
	tobiigaze_max_config_key_length                = 128
)

// The possible gaze tracking statuses.
// Taken from tobiigaze_data_types.h
const (
	tobiigaze_tracking_status_no_eyes_tracked = iota
	tobiigaze_tracking_status_both_eyes_tracked
	tobiigaze_tracking_status_only_left_eye_tracked
	tobiigaze_tracking_status_one_eye_tracked_probably_left
	tobiigaze_tracking_status_one_eye_tracked_unknown_which
	tobiigaze_tracking_status_one_eye_tracked_probably_right
	tobiigaze_tracking_status_only_right_eye_tracked
)

// The possible calibration point statuses
// Taken from tobiigaze_data_types.h
const (
	tobiigaze_calibration_point_status_failed_or_invalid = iota - 1
	tobiigaze_calibration_point_status_valid_but_not_used_in_calibration
	tobiigaze_calibration_point_status_valid_and_used_in_calibration
)

const (
	tobiigaze_option_timeout = iota //Timeout for synchronous operations. Value is of type uint32_t
)

type tobiigaze_device_info struct {
	serial_number    [tobiigaze_device_info_max_serial_number_length]byte
	model            [tobiigaze_device_info_max_model_length]byte
	generation       [tobiigaze_device_info_max_generation_length]byte
	firmware_version [tobiigaze_device_info_max_firmware_length]byte
}

type tobiigaze_point_2d struct {
	x, y float64
}

type usb_device_info struct {
	serialNumber    [tobiigaze_device_info_max_size]byte
	productName     [tobiigaze_device_info_max_size]byte
	platformType    [tobiigaze_device_info_max_size]byte
	firmwareVersion [tobiigaze_device_info_max_size]byte
}
