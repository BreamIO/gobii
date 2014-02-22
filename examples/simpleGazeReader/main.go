package main

import (
	"fmt"
	"time"
	tobii "github.com/zephyyrr/gobii/gaze"
)

func main() {
	fmt.Println("Attempting to find a connected EyeTracker...")
	et, err := tobii.AnyConnectedEyeTracker()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer et.Close()
	
	fmt.Printf("Tracker: %#v\n", et)
	// gaze := et.GetGazeChannel();
	
	// go func(){
	// 	for point := range gaze {
	// 		fmt.Println(point);
	// 	}
	// }()

	time.Sleep(time.Second*30);
}
