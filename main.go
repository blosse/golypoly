package main

import (
	"time"

	"github.com/blosse/golypoly/audio"
	"github.com/blosse/golypoly/synth"
)

func main() {
	// Make cool pew pew noises here
	osc := synth.NewOscillator(44100, synth.Sine)
	synth := synth.Synth{Osc1: *osc, SampleRate: 44100}
	stream := audio.StartAudio(synth)

	time.Sleep(5 * time.Second)

	audio.StopAudio(stream)
}
