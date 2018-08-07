package libmusic

import "strings"

type Notes []Note

func (n Notes) String() string {
	notes := []string{}
	for _, note := range n {
		notes = append(notes, note.String())
	}
	return strings.Join(notes, "-")
}

func (n Notes) Intervals() Intervals {
	intervals := Intervals{}
	for i := 1; i < len(n); i++ {
		intervals = append(intervals, n[0].IntervalTo(n[i]))
	}
	return intervals
}
