package main

import (
	"log"

	"github.com/gordonklaus/portaudio"
)

type AudioSettings struct {
	sampleRate float64
	numInputChannels int
	numOutputChannels int
	framesPerBuffer int
}

// Do i really need this?
var audioSettings = AudioSettings {
	sampleRate: 44100,
	numInputChannels: 0,
	numOutputChannels: 2,
	framesPerBuffer: 64,
}

func AudioCallback() {}

func StartAudio() *portaudio.Stream {
	err := portaudio.Initialize()
	if err != nil {
		log.Fatalf("Unable to initialize portaudio", err)
	}

	stream, err := portaudio.OpenDefaultStream(
		audioSettings.numInputChannels,
		audioSettings.numOutputChannels,
		audioSettings.sampleRate,
		audioSettings.framesPerBuffer,
		AudioCallback,
		/*Add data pointer here*/)
	if err != nil {
		log.Fatalf("Unable to open default stream", err)
	}

	err = stream.Start()
	if err != nil {
		log.Fatalf("Unable to start stream", err)
	}
	return stream
}


func StopAudio(stream *portaudio.Stream) {
	err := stream.Stop()
	if err != nil {
		log.Fatalf("Unable to stop stream", err)
	}

	err = stream.Close()
	if err != nil {
		log.Fatalf("Unable to close stream", err)
	}

	err = portaudio.Terminate()
	if err != nil {
		log.Fatalf("Unable to terminate portaudio", err)
	}
}
