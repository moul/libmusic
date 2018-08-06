package libmusic

type ScaleType uint32

const (
	Major ScaleType = iota
	Minor
)

func (n Note) Scale(kind ScaleType) Chord {
	switch kind {
	case Major:
		return Chord{n, n.Add(4 * Semitone), n.Add(7 * Semitone)}
	case Minor:
		return Chord{n, n.Add(3 * Semitone), n.Add(7 * Semitone)}
	}
	panic("no such scale type")
}
