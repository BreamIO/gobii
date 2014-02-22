package main

import (
	"fmt"
<<<<<<< HEAD
	"time"
	tobii "github.com/zephyyrr/gobii/gaze"
)

func main() {
	fmt.Println("Attempting to find a connected EyeTracker...")
	et, err := tobii.AnyConnectedEyeTracker()

	if err != nil {
		fmt.Println("Error:", err)
		return
=======
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
		log.Fatalln("Error:",err)
>>>>>>> 82054268482a78533142b2de476a788fcf4f61d3
	}

	defer et.Close()
	
	fmt.Printf("Tracker: %#v\n", et)
<<<<<<< HEAD
	// gaze := et.GetGazeChannel();
=======
	checked(et.Connect())
	
	//et.startTracking(func (data GazeData) {
	//	fmt.Println(data)
	//})
>>>>>>> 82054268482a78533142b2de476a788fcf4f61d3
	
	// go func(){
	// 	for point := range gaze {
	// 		fmt.Println(point);
	// 	}
	// }()

	time.Sleep(time.Second*30);
	*/
}

func checked (err error) {
	if err != nil {
		log.Fatalln(err);
	}
}
