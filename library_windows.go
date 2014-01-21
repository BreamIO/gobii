// +build windows

package tobii

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	txInitializeSystem = iota
	txUninitializeSystem
	txCreateContext
	txGetIsSystemInitialized
	txReleaseContext
	txGetContext
	txGetTrackedObjects
	txGetObjectType
	txGetObjectTypeName
	txReleaseObject
	txEnableConnection
	txDisableConnection
	txWriteLogMessage
	txFormatObjectAsText
	txSetInvalidArgumentHandler

	// not a function
	lastIndex
)

const eyexName = `Tobii.EyeX.Client.dll`

var (
	txFunc = make([]*syscall.Proc, lastIndex, lastIndex)

	txName = []string{
		"txInitializeSystem",
		"txUninitializeSystem",
		"txGetIsSystemInitialized",
		"txCreateContext",
		"txReleaseContext",
		"txGetContext",
		"txGetTrackedObjects",
		"txGetObjectType",
		"txGetObjectTypeName",
		"txReleaseObject",
		"txEnableConnection",
		"txDisableConnection",
		"txWriteLogMessage",
		"txFormatObjectAsText",
		"txSetInvalidArgumentHandler",
	}
)

func abort(funcname string, err error) {
        panic(fmt.Sprintf("%s failed: %v", funcname, err))
}

func wInitializeSystem() error {
	ret, _, callErr := txFunc[txInitializeSystem].Call(
		txSystemComponentOverrideFlagNone,
		0, // null
		0) // null

	if callErr != nil {
		abort(txName[txInitializeSystem], callErr)
	}

	result := txResult(ret)

	if result != txResultOk {
		return result
	}

	return nil
}

func wCreateContext(smoething bool) (uintptr, error) {
	const nargs uintptr = 3
	var handle uintptr

	ret, _, callErr := txFunc[txCreateContext].Call(
		uintptr(unsafe.Pointer(&handle)),
		0) //false

	if callErr != nil {
		abort(txName[txCreateContext], callErr)
	}

	result := txResult(ret)

	if result != txResultOk {
		return 0, result
	}

	return handle, nil
}

func init() {
	var err error

	eyex := syscall.MustLoadDLL(eyexName)

	for i, name := range txName {
		txFunc[i], err = eyex.FindProc(name)

		if err != nil {
			abort("Loading Tobii EyeX function " + name, err)
		}
	}
}
