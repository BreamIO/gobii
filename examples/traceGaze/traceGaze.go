package main

import (
	"log"

	termbox "github.com/nsf/termbox-go"
	"github.com/zephyyrr/gobii/gaze"
)

const (
	leftMarker  = termbox.ColorGreen
	rightMarker = termbox.ColorRed
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
	termbox.Clear(0, termbox.ColorWhite) //Set bg to white.
	et.StartTracking(func(data *gaze.GazeData) {
		ts := data.TrackingStatus()
		if ts >= gaze.BothEyesTracked && ts != gaze.OneEyeTrackedUnknownWhich {
			termbox.Clear(0, termbox.ColorWhite)
			drawPoint(data.Left().GazePointOnDisplay(), leftMarker)
			drawPoint(data.Right().GazePointOnDisplay(), rightMarker)
			termbox.Flush()
		}
	})

	done := make(chan struct{})

	go func() {
		for {
			event := termbox.PollEvent()
			if event.Type == termbox.EventKey && event.Key == termbox.KeyEsc {
				close(done)
			}
		}
	}()

	<-done
}

func drawPoint(point gaze.Point2D, attributes termbox.Attribute) {
	sizeX, sizeY := termbox.Size()
	x, y := int(float64(sizeX)*point.X()), int(float64(sizeY)*point.Y())
	switch {
	case x < 0:
		x = 0
	case x >= sizeX:
		x = sizeX - 1
	}

	switch {
	case y < 0:
		y = 0
	case y >= sizeY:
		y = sizeY - 1
	}
	termbox.CellBuffer()[sizeX*y+x] = termbox.Cell{'O', attributes, termbox.ColorWhite}
}
