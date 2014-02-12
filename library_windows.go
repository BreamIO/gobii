// +build windows

package gobii

//#include <Windows.h>
import "C"

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
const OK = "Åtgärden har slutförts."

var (
	txFunc = make([]*syscall.Proc, lastIndex, lastIndex)

	txName = []string{
		"txInitializeSystem",
		"txUninitializeSystem",
		"txCreateContext",
		"txGetIsSystemInitialized",
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

func wInitializeSystem() {
	//txInitializeSystem(TX_SYSTEMCOMPONENTOVERRIDEFLAG_NONE, NULL, NULL)
	ret, _, callErr := txFunc[txInitializeSystem].Call(
		txSystemComponentOverrideFlagNone,
		0, // null
		0) // null

	result := txResult(ret)

	if result != txResultOk {
		abort(txName[txInitializeSystem], callErr)
	}
	
}

func wCreateContext(something bool) (uintptr, error) {
	var handle uintptr = TxEmptyHandle;
	pointer := uintptr(unsafe.Pointer(&handle)) //**void

	ret, _, _ := txFunc[txCreateContext].Call(
		pointer,
		txTrue)

	result := txResult(ret)
	if result == txResultOk {
		return handle, nil
	}

	return 0, result
}

func wReleaseContext(handle uintptr) error {
	ret, _, _ := txFunc[txReleaseContext].Call(
		uintptr(unsafe.Pointer(&handle)),
		txCleanupTimeDefault,
		txTrue) //logLeakingObjectsInfo
	
	result := txResult(ret)
	if result == txResultOk {
		return nil
	}
	
	return result
}

func init() {
	var err error

	// Hack to make Windows report
	// what caused a LoadDLL failure
	C.SetErrorMode(0)

	eyex, err := syscall.LoadDLL(eyexName)

	if err != nil {
		abort("Failed to load "+eyexName, err)
	}

	for i, name := range txName {
		txFunc[i], err = eyex.FindProc(name)

		if err != nil {
			abort("Loading Tobii EyeX function "+name, err)
		}
	}
	wInitializeSystem()
}
