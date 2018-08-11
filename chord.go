package libmusic

import (
	"fmt"
	"sort"
)

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

func (c Chord) copy() Chord {
	copy := Chord{}
	for _, note := range c {
		copy = append(copy, note)
	}
	return copy
}

func (c Chord) SortNotes() Chord {
	copy := c.copy()
	sort.Sort(ByNotes(copy))
	return copy
}

func (c Chord) Simplify() Chord {
	chord := Chord{c[0]}
	letters := map[Letter]bool{}
	letters[c[0].letter] = true
	for _, n := range c.SortNotes() {
		if _, found := letters[n.letter]; !found {
			letters[n.letter] = true
			n.octave = chord[0].octave
			if n.Index() < c[0].Index() {
				n = n.Augment(Octave)
			}
			chord = append(chord, n)
		}
	}
	return chord.SortNotes()
}

func (c Chord) Notes() Notes {
	return Notes(c)
}

func (c Chord) Prepend(n Note) Chord {
	return append(Chord{n}, c...)
}

func (c Chord) String() string {
	perms := c.Permutations()
	for idx := range chordIntervals {
		chordType := ChordType(idx)
		for _, perm := range perms {
			if perm.Equal(perm[0].Chord(chordType)) {
				return fmt.Sprintf("%s%s", perm[0].letter.String(), chordType.String())
			}
		}
	}
	return c.Notes().String()
}

func (c Chord) Normalize() Chord {
	perms := c.Permutations()
	for idx := range chordIntervals {
		chordType := ChordType(idx)
		for _, perm := range perms {
			if perm.Equal(perm[0].Chord(chordType)) {
				return perm
			}
		}
	}
	return c
}

func (c Chord) chordType() ChordType {
	for idx := range chordIntervals {
		chordType := ChordType(idx)
		if c.Equal(c[0].Chord(chordType)) {
			return chordType
		}
	}
	return UnknownChordType
}

func (c Chord) Augment(i Interval) Chord {
	chord := Chord{}
	for _, n := range c {
		chord = append(chord, n.Augment(i))
	}
	return chord
}

func (c Chord) Permutations() []Chord {
	c = c.Simplify()
	permutations := []Chord{c}
	for i := 1; i < len(c); i++ {
		permutations = append(permutations, c.Prepend(c[i].Augment(-1*Octave)).Simplify().Augment(Octave))
	}
	return permutations
}

type ChordType int

const (
	MajorChord ChordType = iota
	MinorChord
	AugmentedChord
	DiminishedChord
	DominantSeventhChord
	MinorSeventhChord
	MajorSeventhChord
	AugmentedMinorSeventhChord
	DiminishedSeventhChord
	HalfDiminishedSeventhChord

	UnknownChordType = -1
)

var chordIntervals = []Intervals{
	{MajorThird, Fifth},
	{MinorThird, Fifth},
	{MajorThird, AugmentedFifth},
	{MinorThird, DiminishedFifth},
	{MajorThird, Fifth, MinorSeventh},
	{MinorThird, Fifth, MinorSeventh},
	{MajorThird, Fifth, MajorSeventh},
	{MajorThird, AugmentedFifth, MinorSeventh},
	{MinorThird, DiminishedFifth, MajorSixth},
	{MinorThird, DiminishedFifth, MinorSeventh},
}

var chordNames = []string{
	"",
	"m",
	"aug",
	"dim",
	"7",
	"m7",
	"maj7",
	"aug7",
	"dim7",
	"7♭5",
}

func (c ChordType) String() string {
	return chordNames[c]
}

func (c ChordType) Intervals() Intervals {
	return chordIntervals[c]
}

func (n Note) Chord(kind ChordType) Chord {
	chord := Chord{n}
	for _, interval := range kind.Intervals() {
		chord = append(chord, n.Augment(interval))
	}
	return chord
}

// COFRight returns the next chord when you go clockwise around the circle of fifths
func (c Chord) COFRight() Chord {
	return c.Augment(Fifth)
}

// COFLeft returns the next chord when you go clockwise around the circle of fifths
func (c Chord) COFLeft() Chord {
	return c.Augment(-Fifth)
}

func (c Chord) COFDown() Chord {
	switch c.chordType() {
	case MajorChord:
		return c[0].Augment(-3 * Semitone).Chord(MinorChord)
	case MajorSeventhChord:
		return c[0].Augment(-3 * Semitone).Chord(MinorSeventhChord)
	}
	return c
}

func (c Chord) COFUp() Chord {
	switch c.chordType() {
	case MinorChord:
		return c[0].Augment(3 * Semitone).Chord(MajorChord)
	case MinorSeventhChord:
		return c[0].Augment(3 * Semitone).Chord(MajorSeventhChord)
	}
	return c
}
