package main

import (
	"fmt"
	"time"
	tobii "github.com/zephyyrr/gobii"
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
	gaze := et.GetGazeChannel();
	
	go func(){
		for point := range gaze {
			fmt.Println(point);
		}
	}()
	
	time.Sleep(time.Second*30);
}
