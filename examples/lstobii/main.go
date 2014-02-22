package main

import (
	"fmt"
	"github.com/zephyyrr/gobii/gaze"
)

func main() {
	fmt.Println("Gaze version:", gaze.Version())
	fmt.Println("USB trackers:")
	//gaze.ListUSBTrackers();
	list, _ := gaze.USBTrackers()
	for i, tracker := range list {
		fmt.Println(i, ": ", tracker)
	}
}