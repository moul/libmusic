package libmusic

import "fmt"

type Note struct {
	letter Letter
	octave int
}

func (n Note) Add(interval Interval) Note {
	note := n
	note.letter += Letter(interval)
	for note.letter > 11 {
		note.letter -= 12
		note.octave++
	}
	for note.letter < 0 {
		note.letter += 12
		note.octave--
	}
	return note
}

func (n Note) String() string {
	return fmt.Sprintf("%s%d", letters[n.letter], n.octave)
}

func (n Note) Equal(other Note) bool {
	return n.letter == other.letter && n.octave == other.octave
}

// func (n Note) Interval(note Note) Interval {}
//func (n Note) RelativeInterval(note Note) Interval {}
