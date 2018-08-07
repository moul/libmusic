package libmusic

import "strings"

type Progression []Chord

func (p Progression) String() string {
	chords := []string{}
	for _, chord := range p {
		chords = append(chords, chord.String())
	}
	return strings.Join(chords, " ")
}
