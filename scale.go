package libmusic

type ScaleType int

const (
	HemiPentatonicScale ScaleType = iota
	MajorPentatonicScale
	MinorPentatonicScale
	HexatonicWholeToneScale
	BluesScale
	LocrianModeScale
	PhrygianModeScale
	SpanishScale
	NaturalMinorScale
	HarmonicMinorScale
	DorianModeScale
	MelodicMinorScale
	MyxolydianModeScale
	MajorScale
	LydianModeScale
	ChromaticScale

	// aliases

	SynolydianModeScale  = LydianModeScale
	IonianModeScale      = MajorScale
	AeolianModeScale     = NaturalMinorScale
	MessiaenScale        = HexatonicWholeToneScale
	HyperdorianModeScale = AeolianModeScale

	// HypophrygianModeScale
)

var scaleIntervals = []Intervals{
	{1, 4, 2, 1, 4},                      // hemi pentatonic
	{2, 2, 3, 2, 3},                      // major pentatonic
	{3, 2, 2, 3, 2},                      // minor pentatonic
	{2, 2, 2, 2, 2, 2},                   // hexatonic whole tone
	{3, 2, 1, 1, 3, 2},                   // blues
	{1, 2, 2, 1, 2, 2, 2},                // locrian mode
	{1, 2, 2, 2, 1, 2, 2},                // phrygian mode
	{1, 3, 1, 2, 1, 2, 2},                // spanish (gipsy)
	{2, 1, 2, 2, 1, 2, 2},                // natural minor
	{2, 1, 2, 2, 1, 3, 1},                // harmonic minor
	{2, 1, 2, 2, 2, 1, 2},                // dorian mode
	{2, 1, 2, 2, 2, 2, 1},                // melodic minor
	{2, 2, 1, 2, 2, 1, 2},                // myxolydian mode
	{2, 2, 1, 2, 2, 2, 1},                // major
	{2, 2, 2, 1, 2, 2, 1},                // lydian mode
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, // chromatic
}

func (n Note) Scale(kind ScaleType) Notes {
	notes := Notes{n}
	currentNote := n
	for _, interval := range scaleIntervals[kind] {
		currentNote = currentNote.Augment(interval)
		if currentNote.Index() >= n.Augment(Octave).Index() {
			break
		}
		notes = append(notes, currentNote)
	}
	return notes
}
