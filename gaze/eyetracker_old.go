// +build ignore

package gaze

import (
	"fmt"
	"time"
)

type EyeTracker struct {
	handle   uintptr
	gazeChan chan GazeData
}

func NewEyeTracker() (*EyeTracker, error) {
	//txCreateContext(&hContext, TX_FALSE);
	handle, err := wCreateContext(false)

	if err != nil {
		return nil, err
	}

	return &EyeTracker{handle: handle}, nil
}

type GazeData struct {
	X, Y      float64
	Timestamp time.Time
}

func (d GazeData) String() string {
	return fmt.Sprintf("{X: %f, Y: %f, Time: %d}", d.X, d.Y, d.Timestamp)
}

func (e *EyeTracker) GetGazeChannel() <-chan GazeData {
	if e.gazeChan == nil {
		e.gazeChan = make(chan GazeData)
	}

	return e.gazeChan
}

func (e *EyeTracker) Close() {
	err := wReleaseContext(e.handle)

	if err != nil {
		panic(err.Error())
	}
}

func (e *EyeTracker) onStateChanged() {
	//txRegisterEventHandler(hContext, &hEventHandlerTicket, HandleEvent, NULL);
}
