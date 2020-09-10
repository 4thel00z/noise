package main

import (
	"flag"
	"github.com/gen2brain/beeep"
)

var (
	frequency = flag.Float64("frequency", beeep.DefaultFreq, "frequency for the noise")
	duration  = flag.Int("duration", beeep.DefaultDuration, "duration for the noise")
)

func main() {
	flag.Parse()
	err := beeep.Beep(*frequency, *duration)
	if err != nil {
		panic(err)
	}
}
