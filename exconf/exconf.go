package main

/*
#include "tobiigaze_display_area.h"
#include "tobiigaze_data_types.h"
*/
import "C"

import (
	"github.com/zephyyrr/gobii/gaze"
	"log"
	"math"
)

func main() {
	trackerinfos, err := gaze.USBTrackers()
	if err != nil {
		log.Println("Error @ main#getTrackers :", err)
		return
	}

	for i, trackerinfo := range trackerinfos {
		tracker, err := gaze.EyeTrackerFromURL("tet-usb://" + trackerinfo.SerialNumber())
		if err != nil {
			log.Println("Error @ main#createTracker", i, ":", err)
			continue
		}
		tracker.Connect()
		err = tracker.SetDisplayArea(520, 225, math.Pi/16)
		if err != nil {
			log.Println("Error @ main#SetDisplayArea", i, ":", err)
			return
		}
		log.Println("Set DisplayArea of", "tet-usb://"+trackerinfo.SerialNumber())
	}
}
