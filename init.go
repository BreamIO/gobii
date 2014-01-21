package tobii

import (
	"log"
)

func init() {
	err := wInitializeSystem()

	if err != nil {
		log.Fatal(err)
	}
}