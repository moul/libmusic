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

func (n Notes) Ring() chan Note {
	c := make(chan Note, 1)
	go func() {
		for {
			for _, note := range n {
				c <- note
			}
		}
	}()
	return c
}

type ByNotes []Note

func (n ByNotes) Len() int           { return len(n) }
func (n ByNotes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByNotes) Less(i, j int) bool { return n[i].Index() < n[j].Index() }
