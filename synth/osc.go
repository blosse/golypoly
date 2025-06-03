package synth

import (
	"math"
	"sync/atomic"
)

type Oscillator struct {
	frequencyBits uint64 // Atomic
	amplitudeBits uint64 // Atomic
	phase float64
	sampleRate float64
	waveform atomic.Value //func(phase float64) float64
}

func NewOscillator(sampleRate float64, waveform func(phase float64) float64) *Oscillator {
	osc := &Oscillator{sampleRate: sampleRate}
	osc.waveform.Store(waveform)
	atomic.StoreUint64(&osc.frequencyBits, math.Float64bits(440))
	atomic.StoreUint64(&osc.amplitudeBits, math.Float64bits(1.0))

	return osc
}

func (osc * Oscillator) SetFrequency(freq float64) {
	atomic.StoreUint64(&osc.frequencyBits, math.Float64bits(freq))
}
func (osc * Oscillator) SetAmplitude(amplitude float64) {
	atomic.StoreUint64(&osc.amplitudeBits, math.Float64bits(amplitude))
}

// Create sample and update phase
func (osc *Oscillator) Process() float64 {
	freq := math.Float64frombits(atomic.LoadUint64(&osc.frequencyBits))
	amplitude := math.Float64frombits(atomic.LoadUint64(&osc.amplitudeBits))

	wave := osc.waveform.Load().(func(float64) float64)
	sample := amplitude * wave(osc.phase)
	osc.phase += 2 * math.Pi * freq / osc.sampleRate
	if osc.phase > 2 * math.Pi {
		osc.phase -= 2 * math.Pi
	}
	return sample
}

func Sine(phase float64) float64 {
	return math.Sin(phase)
}

func Saw(phase float64) float64 {
	return 2*(phase/math.Pi - 1)
}
func Square(phase float64) float64 {
	if math.Sin(phase) > 0 {
		return 1
	} 
	return -1
}
