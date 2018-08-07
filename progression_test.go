package libmusic

import "fmt"

func ExampleProgression_String() {
	fmt.Println(Progression{
		C4.Scale(MajorChord, 1),
		E4.Scale(MinorChord, 1),
		G4.Scale(MinorSeventhChord, 1),
	})
	// Output: C Em Gm7
}
