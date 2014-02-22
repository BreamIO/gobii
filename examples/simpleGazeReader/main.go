package main

import (
	"fmt"
	"github.com/zephyyrr/gobii/gaze"
	"log"
	"time"
)

func main() {
	log.Println("Creating tracker...")
	et, err := gaze.AnyConnectedEyeTracker()

	//err = et.onConnectionStateChanged();
	if err != nil {
		log.Fatalln("Error:", err)
	}

	defer et.Close()
	log.Println("Tracker created.")
	log.Println("Connecting to tracker.")
	fmt.Printf("Tracker: %#v\n", et)

	checked(et.Connect())
	log.Println("Connected!")

	//et.startTracking(func (data GazeData) {
	//	fmt.Println(data)
	//})

	time.Sleep(time.Second * 30)
}

func checked(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
