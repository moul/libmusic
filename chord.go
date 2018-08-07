package libmusic

import "strings"

type Chord []Note

type Notes []Note

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

func (n Notes) String() string {
	notes := []string{}
	for _, note := range n {
		notes = append(notes, note.String())
	}
	return strings.Join(notes, "-")
}
