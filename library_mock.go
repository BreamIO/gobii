// +build linux

package tobii

import (
	"errors"
)

// DLL functions are prefixed with 'tx'
// Wrapper functions are prefixed with 'w'

var initialized = false

func wInitializeSystem() error {
	if initialized {
		return errors.New(txResultSystemAlreadyInitialized.String())
	}
	initialized = true
	return nil
}

func wCreateContext(somethingElse bool) (uintptr, error) {
	// const nargs uintptr = 4
	var handle uintptr = 1

	/*
		res, _, callErr := syscall.Syscall(txCreateContext,
			nargs,
			&handle,
			wBool(bool))

		if res != txResultOk {
			return 0, errors.New(txResult(res).String())
		}

	*/
	return handle, nil
}

func wBool(isTrue bool) uintptr {
	if isTrue {
		return 1
	}
	return 0
}
