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
	for idx := range scaleIntervals {
		chordType := ChordType(idx)
		for _, perm := range perms {
			if perm.Equal(perm[0].Scale(chordType, 1)) {
				return fmt.Sprintf("%s%s", perm[0].letter.String(), chordType.String())
			}
		}
	}
	panic("unknown chord type")
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
)

var scaleIntervals = []Intervals{
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
	return scaleIntervals[c]
}

func (n Note) Scale(kind ChordType, numOctaves int) Chord {
	chord := Chord{n}
	for octave := 0; octave < numOctaves; octave++ {
		if octave > 0 {
			chord = append(chord, n.Augment(Interval(octave)*Octave))
		}
		for _, interval := range kind.Intervals() {
			chord = append(chord, n.Augment(interval+Interval(octave)*Octave))
		}
	}
	return chord
}
