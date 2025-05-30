package synth

type Oscillator struct {
	phase float64
	amplitude float64
	wave int
}
type Synth struct {
	osc1 Oscillator
	osc2 Oscillator
	osc3 Oscillator
	osc4 Oscillator
}

const (
	sine = iota
	square
	triangle
	saw
)
	
	
