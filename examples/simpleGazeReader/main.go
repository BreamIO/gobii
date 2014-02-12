package main

import (
	"fmt"
	"time"
	"github.com/zephyyrr/gobii"
)

func main() {
	et, err := tobii.NewEyeTracker()
	defer et.Close()
	//err = et.onConnectionStateChanged();
	if err != nil {
		fmt.Println("Error:",err)
		return
	}

	fmt.Printf("Tracker: %#v\n", et)
	//gaze := et.GetGazeChannel();
	
	time.Sleep(time.Second)
}
