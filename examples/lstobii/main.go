package main

import (
	"fmt"
	"github.com/zephyyrr/gobii/gaze"
)

func main() {
	fmt.Println("USB trackers:")
	gaze.ListTrackers();
	/*
	for _, tracker := range gaze.ListUSBTrackers() {
		fmt.Println(tracker)
	}
	*/
	
}