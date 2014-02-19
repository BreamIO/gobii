package gaze

import (
	"fmt"
	"unsafe"
	"unicode/utf16"
)

type Error int

// Converts a C string (*char) to a Go string.
func convertCStringToString(charPtr uintptr) string {
	// If given a null pointer, return the empty string.
	if charPtr == 0 {
		return ""
	}

	// Make a buffer
	buf := make([]uint16, 0, 256)

	// For each 16-bit character
	for ptr := charPtr; ; ptr += 2 {
		// Dereference the current pointer position to an uint16
		char := *(*uint16)(unsafe.Pointer(ptr))

		// If current char is '\0', the string is terminated.
		if char == 0 {
			return string(utf16.Decode(buf))
		}

		// Append the character to the string.
		buf = append(buf, char)
	}
}

func (e Error) Error() string {
	res, _, _ := tobiigaze[create].Call(uintptr(e))

	return convertCStringToString(res)
}
