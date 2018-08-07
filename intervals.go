package libmusic

import "strings"

type Intervals []Interval

func (i Intervals) String() string {
	intervals := []string{}
	for _, interval := range i {
		intervals = append(intervals, interval.String())
	}
	return strings.Join(intervals, "-")
}
