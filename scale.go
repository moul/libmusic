package libmusic

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
