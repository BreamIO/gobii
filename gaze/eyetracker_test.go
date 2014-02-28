package gaze_test

import (
	"fmt"
	"github.com/zephyyrr/gobii/gaze"
)

// This is what you are really looking for.
// The basic use of it is to first create one.
// Then you connect to it using Connect.
// At this point, defer a Close().
// Then you need to StartTracking.
// Give us a function of the GazeFunc type and we handle the rest.
func ExampleEyeTracker() {
	et, err := gaze.AnyEyeTracker()
	if err != nil {
		fmt.Println("AHHHH!!! ERROR!!!")
		fmt.Println(err)
		return
	}
	et.Connect()
	defer et.Close()
	//et.StartTracking(func(data ETData) {
	//	log.Println(data)
	//})
}