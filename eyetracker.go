package tobii

import (
	"time"
)

type EyeTracker struct {
	handle uintptr
	gazeChan chan GazeData
}

func NewEyeTracker() (*EyeTracker, error) {
	handle, err := wCreateContext(false)

	if err != nil {
		return nil, err
	}

	return &EyeTracker{handle: handle}, nil
}

type GazeData struct {
	X, Y float64
	Timestamp time.Time
}

func (e *EyeTracker) GetGazeChannel() <-chan GazeData {
	if e.gazeChan == nil {
		e.gazeChan = make(chan GazeData)
	}

	return e.gazeChan
}
