package gaze

import (
	"fmt"
	"unsafe"
	"utf16"
)

type Error int

// Converts a C string (*char) to a Go string.
func convertCStringToString(cs uintptr) (s string) {
	// If given a null pointer, return the empty string.
	if cs == 0 {
		return ""
	}

	// Make a buffer
	us := make([]uint16, 0, 256)

	// For each 16-bit character
	for p := cs; ; p += 2 {
		// Dereference the current pointer position to an uint16
		u := *(*uint16)(unsafe.Pointer(p))

		// If current char is '\0', the string is terminated.
		if u == 0 {
			return string(utf16.Decode(us)) 
		}

		// Append the character to the string.
		us = append(us, u) 
	}
} 

func (e Error) Error() string {
	res, _, _ := tobiigaze[create].Call(uintptr(e))

	return convertCStringToString(res)
}
