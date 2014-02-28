package gaze_test

import (
	"fmt"
	"github.com/zephyyrr/gobii/gaze"
)

// This is what you are really looking for.
//
// The example below shows how to use the a tracker.
// The expected output, if a tracker is connected, is the points on the screen where your eyes are looking.
func ExampleEyeTracker() {
	et, err := gaze.AnyEyeTracker()
	if err != nil {
		fmt.Println("AHHHH!!! ERROR!!!")
		fmt.Println(err)
		return
	}
	et.Connect()
	defer et.Close()
	et.StartTracking(func(data ETData) {
		fmt.Println(data)
	})
}