//Copyright 2014 Bream IO AB, All rights reserved

// Example showing how to access a tracker and getting data from it.
package main

import (
	"fmt"
	"github.com/zephyyrr/gobii/gaze"
	"log"
	"time"
)

func main() {
	log.Println("Creating tracker...")
	et, err := gaze.AnyEyeTracker()

	//err = et.onConnectionStateChanged();
	if err != nil {
		log.Fatalln("Error:", err)
	}

	defer et.Close()
	log.Println("Tracker created.")
	log.Println("Connecting to tracker.")
	fmt.Printf("Tracker: %#v\n", et)

	checked(et.Connect())
	log.Println("Connected!")

	log.Println(et.Info())

	//et.startTracking(func (data GazeData) {
	//	fmt.Println(data)
	//})

	time.Sleep(time.Second * 30)
}

func checked(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
