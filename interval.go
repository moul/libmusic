package libmusic

type Interval int

const (
	// tones

	Semitone Interval = 1
	Tone              = 2

	// intervals

	MinorSecond                      = 1
	MajorSecond                      = 2
	MinorThird                       = 3
	MajorThird                       = 4
	Fourth                           = 5
	DiminishedFifth, AugmentedFourth = 6, 6
	Fifth, DiminishedSixth           = 7, 7
	MinorSixth, AugmentedFifth       = 8, 8
	MajorSixth                       = 9
	MinorSeventh                     = 10
	MajorSeventh                     = 11
	Octave                           = 12
	MinorNinth                       = 13
	MajorNinth                       = 14
	MinorTenth                       = 15
	MajorTenth                       = 16
	Eleventh                         = 17
	Twelfth, Tritave                 = 19, 19
	MinorThirteenth                  = 20
	MajorThirteenth                  = 21
	MinorFourteenth                  = 22
	MajorFourteen                    = 23
	Fifteenth, DoubleOctave          = 24, 24
)
