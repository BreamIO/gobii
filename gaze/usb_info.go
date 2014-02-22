package gaze

/*
#include <stdlib.h>
#include "tobiigaze.h"
#include "tobiigaze_discovery.h"

size_t getInfoSize() {
	return sizeof(struct usb_device_info);
}
*/
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

type USBInfo C.struct_usb_device_info

func (info USBInfo) cPtr() *C.struct_usb_device_info {
	return (*C.struct_usb_device_info)(&info)
}

func (info USBInfo) SerialNumber() string {
	return C.GoString((*C.char)(&info.cPtr().serialNumber[0]))
}

func (info USBInfo) ProductName() string {
	return C.GoString((*C.char)(&info.cPtr().productName[0]))
}

func (info USBInfo) PlatformType() string {
	return C.GoString((*C.char)(&info.cPtr().platformType[0]))
}

func (info USBInfo) FirmwareVersion() string {
	return C.GoString((*C.char)(&info.cPtr().firmwareVersion[0]))
}

func (info USBInfo) String() string {
	return fmt.Sprintf("\tName: %s\n\tSerialnumber: %s\n\tPlatform: %s\n\tFirmware: %s", 
		info.ProductName(), 
		info.SerialNumber(),
		info.PlatformType(),
		info.FirmwareVersion())
}

func USBTrackers() ([]USBInfo, error) {
	var err Error
	capacity := 10
	var length uint32
	
	infos_void := C.malloc(C.getInfoSize() * C.size_t(capacity))
	defer C.free(infos_void)
	infos := (*C.struct_usb_device_info)(infos_void)
	
	C.tobiigaze_list_usb_eye_trackers(infos, 
		C.uint32_t(capacity), 
		(*C.uint32_t)(&length),
		err.cPtr())
		
	var goInfos []USBInfo
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&goInfos)))
	sliceHeader.Cap = int(length)
	sliceHeader.Len = int(length)
	sliceHeader.Data = uintptr(unsafe.Pointer(infos))

	if length == 0 {
		return nil, err
	}
	
	//Transfer data into Go runtime handled memory.
	result := make([]USBInfo, length, length)
	copy(result, goInfos)
	return result, nil;
}

func ListUSBTrackers() error {
	list, err := USBTrackers()
	if err != nil {
		return err;
	}
	for i, info := range list {
		fmt.Println(i, ": ", info)
	}
	return err
}
