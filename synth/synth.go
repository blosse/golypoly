package synth

type AudioUnit interface {
	Process() float32
}

type Synth struct {
	Osc1 Oscillator
	SampleRate float64
}

func (synth *Synth) AudioCallback(out []float32) {
	// Call the process functions of all AudioUnits here
	for i := range out {
		outputSample := synth.Osc1.Process()
		out[i] = float32(outputSample)
	}
}
