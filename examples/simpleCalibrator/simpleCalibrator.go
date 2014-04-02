package main

import (
	"log"
	"time"

	"github.com/zephyyrr/gobii/gaze"
	termbox "github.com/nsf/termbox-go"
)

func main() {
	termbox.Init()
	defer termbox.Close()
	
	et, err := gaze.AnyEyeTracker()
	
	if err != nil {
		log.Println(err)
		return
	}
	
	et.Connect()
	defer et.Close()
	
	done := make(chan struct{})
	
	et.StartCalibration(func (err error) {
		if err != nil {
			log.Println("Could not start calibrating:", err)
			close(done)
			return
		}
		calibratePoint(et, 0.1, 0.1)
		calibratePoint(et, 0.9, 0.1)
		calibratePoint(et, 0.9, 0.9)
		calibratePoint(et, 0.1, 0.9)
		calibratePoint(et, 0.5, 0.5)
		et.ComputeAndSetCalibration(func(err error) {
			et.StopCalibration(func (err error) {
				defer close(done)
				if err != nil {
					log.Println("Could not start calibrating:", err)
					return
				}
			})
		})
	})
	
	<-done
}

const marker = termbox.ColorDefault | termbox.AttrBold

func calibratePoint(et *gaze.EyeTracker, x, y float64) {
	termbox.Clear(0 ,0)
	sizeX, sizeY := termbox.Size()
	dx, dy := int(float64(sizeX)*x), int(float64(sizeY)*y)
	termbox.CellBuffer()[sizeX*dy + dx] = termbox.Cell{'#', marker, termbox.ColorDefault}
	termbox.Flush()
	//log.Printf("Termbox Size: (%d, %d)", sizeX, sizeY)
	log.Printf("Calibrating (%.3f, %.3f)", x, y)
	done := make(chan struct{})
	time.Sleep(1*time.Second) //Minimum time. Gives user time to react.
	et.AddPointToCalibration(gaze.NewPoint2D(x, y), func (err error) {
		close(done) //Synchronizing async call.
	})
	<-done
}