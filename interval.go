package libmusic

type Interval int

const (
	PerfectUnison                        Interval = 0
	Semitone, MinorSecond                Interval = 1, 1
	Tone, MajorSecond                    Interval = 2, 2
	MinorThird                           Interval = 3
	MajorThird                           Interval = 4
	Fourth                               Interval = 5
	DiminishedFifth, AugmentedFourth     Interval = 6, 6
	Fifth, DiminishedSixth               Interval = 7, 7
	MinorSixth, AugmentedFifth           Interval = 8, 8
	MajorSixth                           Interval = 9
	MinorSeventh                         Interval = 10
	MajorSeventh                         Interval = 11
	Octave                               Interval = 12
	MinorNinth                           Interval = 13
	MajorNinth                           Interval = 14
	MinorTenth                           Interval = 15
	MajorTenth                           Interval = 16
	Eleventh                             Interval = 17
	AugmentedEleventh, DiminishedTwelfth Interval = 18, 18
	Twelfth, Tritave                     Interval = 19, 19
	MinorThirteenth                      Interval = 20
	MajorThirteenth                      Interval = 21
	MinorFourteenth                      Interval = 22
	MajorFourteen                        Interval = 23
	Fifteenth, DoubleOctave              Interval = 24, 24
)

var (
	intervals = []string{
		"P1",
		"m2",
		"M2",
		"m3",
		"M3",
		"P4",
		"d5",
		"P5",
		"m6",
		"M6",
		"m7",
		"M7",
		"P8", // octave

		"m9",
		"M9",
		"m10",
		"M10",
		"P11",
		"d12",
		"P12",
		"m13",
		"M13",
		"m14",
		"M14",
		"P15", // double octve
	}
)

func (i Interval) String() string {
	return intervals[i]
}
