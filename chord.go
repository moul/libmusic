package libmusic

import "fmt"

type Chord Notes

func (c Chord) Equal(other Chord) bool {
	if len(c) != len(other) {
		return false
	}
	for i, note := range c {
		if note != other[i] {
			return false
		}
	}
	return true
}

func (c Chord) Notes() Notes {
	return Notes(c)
}

func (c Chord) String() string {
	for idx := range scaleIntervals {
		chordType := ChordType(idx)
		if c.Equal(c[0].Scale(chordType)) {
			return fmt.Sprintf("%s%s", c[0].letter.String(), chordType.String())
		}
	}
	panic("unknown chord type")
}
