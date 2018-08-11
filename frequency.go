package libmusic

import (
	"fmt"
	"math/big"
	"strings"
)

type Hz float64
type Meter float64
type FrequencyRatio struct {
	r *big.Rat
	f float64
}

const (
	soundSpeed = 345.0     // 345 m/s
	a4Freq     = Hz(440.0) // 440hz
)

var justScaleFrequencyRatios = []FrequencyRatio{
	{r: big.NewRat(1, 1)}, // unison
	{r: big.NewRat(25, 24)},
	{r: big.NewRat(9, 8)},
	{r: big.NewRat(6, 5)},
	{r: big.NewRat(5, 4)},
	{r: big.NewRat(4, 3)},
	{r: big.NewRat(45, 32)},
	{r: big.NewRat(3, 2)},
	{r: big.NewRat(8, 5)},
	{r: big.NewRat(5, 3)},
	{r: big.NewRat(9, 5)},
	{r: big.NewRat(15, 8)},
	{r: big.NewRat(2, 1)}, // octave
}

var pythagoreanFrequencyRatios = []FrequencyRatio{
	{r: big.NewRat(1, 1)}, // unison
	{r: big.NewRat(256, 243)},
	{r: big.NewRat(9, 8)},
	{r: big.NewRat(32, 27)},
	{r: big.NewRat(81, 64)},
	{r: big.NewRat(4, 3)},
	{r: big.NewRat(729, 512)},
	{r: big.NewRat(3, 2)},
	{r: big.NewRat(128, 81)},
	{r: big.NewRat(27, 16)},
	{r: big.NewRat(16, 9)},
	{r: big.NewRat(243, 128)},
	{r: big.NewRat(2, 1)}, // octave
}

var equalTemperedFrequencyRatios = []FrequencyRatio{
	{f: 1.0},
	{f: 1.0594630943592953}, // 2^(1/12)
	{f: 1.122462048309373},  // (2^(1/12))^2
	{f: 1.1892071150027212}, // (2^(1/12))^3
	{f: 1.2599210498948733}, // (2^(1/12))^4
	{f: 1.3348398541700346}, // (2^(1/12))^5
	{f: 1.4142135623730953}, // (2^(1/12))^6
	{f: 1.4983070768766818}, // (2^(1/12))^7
	{f: 1.5874010519681999}, // (2^(1/12))^8
	{f: 1.6817928305074296}, // (2^(1/12))^9
	{f: 1.7817974362806792}, // (2^(1/12))^10
	{f: 1.8877486253633877}, // (2^(1/12))^11
	{f: 2.0},
}

func (i Interval) JustScaleRatio() FrequencyRatio {
	if i < 0 {
		return (-i).JustScaleRatio().inv()
	}
	if i > 11 {
		return (i - 12).JustScaleRatio().mulByInt(2)
	}
	return justScaleFrequencyRatios[i]
}

func (f FrequencyRatio) inv() FrequencyRatio {
	n := FrequencyRatio{}
	if f.r != nil {
		n.r = big.NewRat(1, 1).Inv(f.r)
	} else {
		n.f = 1 / f.f
	}
	return n
}

func (f FrequencyRatio) mulByInt(by int) FrequencyRatio {
	n := FrequencyRatio{}
	if f.r != nil {
		n.r = big.NewRat(1, 1).Mul(f.r, big.NewRat(int64(by), 1))
	} else {
		n.f = f.f * float64(by)
	}
	return n
}

func (f FrequencyRatio) Float() float64 {
	if f.r != nil {
		ret, _ := f.r.Float64()
		return ret
	}
	return f.f
}

func (f FrequencyRatio) String() string {
	if f.r != nil {
		return f.r.String()
	}
	return fmt.Sprintf("%s", strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.3f", f.f), "0"), "."))
}

func (i Interval) PythagoreanRatio() FrequencyRatio {
	if i < 0 {
		return (-i).PythagoreanRatio().inv()
	}
	if i > 12 {
		return (i - 12).PythagoreanRatio().mulByInt(2)
	}
	return pythagoreanFrequencyRatios[i]
}

func (n Note) JustScaleFrequency() Hz {
	interval := A4.IntervalTo(n)
	ratio := interval.JustScaleRatio()
	return Hz(float64(a4Freq) * ratio.Float())
}

func (i Interval) EqualTemperedRatio() FrequencyRatio {
	if i < 0 {
		return (-i).EqualTemperedRatio().inv()
	}
	if i > 12 {
		return (i - 12).EqualTemperedRatio().mulByInt(2)
	}
	return equalTemperedFrequencyRatios[i]
}

func (n Note) Frequency() Hz {
	interval := A4.IntervalTo(n)
	ratio := interval.EqualTemperedRatio()
	return Hz(float64(a4Freq) * ratio.Float())
}

func NoteByEqualTemperedFrequency(target Hz) Note {
	n := A4
	// octave operations
	for n.Frequency() <= target/2 {
		n = n.Augment(Octave)
	}
	for n.Frequency() >= target*2 {
		n = n.Augment(-Octave)
	}
	// semitone operations
	for n.Frequency() < target {
		n = n.Augment(Semitone)
	}
	for n.Frequency() > target {
		n = n.Augment(-Semitone)
	}
	// round (nearest)
	if n.Augment(Semitone).Frequency()-target < target-n.Frequency() {
		n = n.Augment(Semitone)
	}
	return n
}

func (h Hz) String() string {
	return fmt.Sprintf("%s㎐", strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.3f", h), "0"), "."))
}

func (h Hz) Relative() Hz {
	if h < 0 {
		return -h
	}
	return h
}

func (h Hz) WaveLength() Meter {
	return Meter(1 / h * soundSpeed)
}

func (m Meter) String() string {
	return fmt.Sprintf("%sm/s", strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.3f", m), "0"), "."))
}
