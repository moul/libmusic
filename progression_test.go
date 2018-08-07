package libmusic

import "fmt"

func ExampleProgression_String() {
	fmt.Println(Progression{
		C4.Scale(MajorChord),
		E4.Scale(MinorChord),
		G4.Scale(MinorSeventhChord),
	})
	// Output: C Em Gm7
}
