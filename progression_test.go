package libmusic

import "fmt"

func ExampleProgression_String() {
	fmt.Println(Progression{
		C4.Chord(MajorChord),
		E4.Chord(MinorChord),
		G4.Chord(MinorSeventhChord),
	})
	// Output: C Em Gm7
}
