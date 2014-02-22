//Copyright 2014 Bream IO AB. All rights reserved.

// Simple packaged that demonstrates how to list connected eye trackers
// Can be used as a debugging tool to detect what trackers is actually detected by the system
// Only lists USB trackers and not TCP-over-USB trackers.
package main

import (
	"fmt"
	"github.com/zephyyrr/gobii/gaze"
)

func main() {
	fmt.Println("Gaze version:", gaze.Version())
	fmt.Println("USB trackers:")

	list, err := gaze.USBTrackers()

	if err != nil {
		fmt.Println("None found, due to:", err)
		return
	}

	if len(list) == 0 {
		fmt.Println("None found.")
		return
	}

	for i, tracker := range list {
		fmt.Println(i, ": ", tracker)
	}
}
