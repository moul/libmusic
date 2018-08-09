package libmusic

import (
	"fmt"
	"strings"
)

type Hz float64
type Meter float64

const (
	semitoneFrequency = 1.059463094359 // 2^(1/12)
	soundSpeed        = 345.0          // 345 m/s
	a4Freq            = 440.0          // 440hz
)

func (n Note) Frequency() Hz {
	var (
		interval   = A4.IntervalTo(n)
		base       = a4Freq
		multiplier = 1.0
	)

	for interval < -11 {
		interval += 12
		base /= 2
	}

	for interval > 11 {
		interval -= 12
		base *= 2
	}

	for i := 0; i > int(interval); i-- {
		multiplier /= semitoneFrequency
	}

	for i := 0; i < int(interval); i++ {
		multiplier *= semitoneFrequency
	}

	return Hz(base * multiplier)
}

func NoteByFrequency(target Hz) Note {
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
