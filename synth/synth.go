package synth

type AudioUnit interface {
	Process() float32
}

type Synth struct {
	Osc1 Oscillator
	SampleRate float64
}

const (
	sine = iota
	square
	triangle
	saw
)

func (synth *Synth) AudioCallback(out []float32) {
	// Call the process functions of all AudioUnits here
	for i := range out {
		out[i] = float32(synth.Osc1.Process())
	}
}
