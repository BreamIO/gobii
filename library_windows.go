// +build windows

package tobii

import (
	"fmt"
	"log"
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
	txLastIndex
)

const eyexName = `Tobii.EyeX.Client.dll`

var (
	txFunc = make([]uintptr, txLastIndex, txLastIndex)

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
	const nargs uintptr = 3

	ret, _, callErr := syscall.Syscall(txFunc[txInitializeSystem],
		nargs,
		txSystemComponentOverrideFlagNone,
		0, // null
		0) // null

	if callErr != 0 {
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

	ret, _, callErr := syscall.Syscall(txFunc[txCreateContext],
		nargs,
		uintptr(unsafe.Pointer(&handle)),
		0, //false
		0)

	if callErr != 0 {
		abort(txName[txCreateContext], callErr)
	}

	result := txResult(ret)

	if result != txResultOk {
		return 0, result
	}

	return handle, nil
}

func init() {
	eyex, err := syscall.LoadLibrary(eyexName)

	if err != nil {
		fmt.Println("There was an error loading " + eyexName)
		log.Fatal(err)
	}

	for i, name := range txName {
		txFunc[i], err = syscall.GetProcAddress(eyex, name)

		if err != nil {
			abort("Initialization of Tobii EyeX function " + name, err)
		}
	}
}
