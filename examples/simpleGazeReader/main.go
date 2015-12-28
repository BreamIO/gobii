// Copyright 2014 Bream IO AB, All rights reserved.

// Example showing how to access a tracker and get data from it.
package main

import (
	"fmt"
	"github.com/BreamIO/gobii/gaze"
	"log"
	"time"
	"flag"
)

var auto = flag.Bool("auto", false, "Automagically find a tracker to use.")

func main() {
	flag.Parse()
	log.Println("Creating tracker...")
	var et *gaze.EyeTracker
	var err error
	if *auto {
		et, err = gaze.AnyEyeTracker()
	} else {
		url := flag.Arg(0)
		et, err = gaze.EyeTrackerFromURL(url);
	}

	if err != nil {
		log.Fatalln("Error:", err)
	}

	defer et.Close()
	log.Println("Tracker created.")
	log.Println("Connecting to tracker.")

	checked(et.Connect())

	log.Println("Connected!")

	info, err := et.Info()

	checked(err)

	log.Println(info)

	et.StartTracking(func(data *gaze.GazeData) {
		ts := data.TrackingStatus()
		if ts >= gaze.BothEyesTracked && ts != gaze.OneEyeTrackedUnknownWhich {
			fmt.Println(data)
		}
	})

	time.Sleep(time.Second * 30)
}

func checked(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
