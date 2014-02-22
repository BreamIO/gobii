package main

import (
	"fmt"
	"log"
	//"time"
	"github.com/zephyyrr/gobii/gaze"
)

func main() {
	url, err := gaze.ConnectedEyeTracker()
	if err != nil{
		log.Fatalln(err)
	}
	et, err := gaze.EyeTrackerFromURL(url)
	defer et.Close()
	//err = et.onConnectionStateChanged();
	if err != nil {
		log.Fatalln("Error:", err)
	}

	defer et.Close()
	
	fmt.Printf("Tracker: %#v\n", et)

	checked(et.Connect())
	
	//et.startTracking(func (data GazeData) {
	//	fmt.Println(data)
	//})
	
	// go func(){
	// 	for point := range gaze {
	// 		fmt.Println(point);
	// 	}
	// }()

	time.Sleep(time.Second*30)
}

func checked (err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
