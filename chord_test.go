package libmusic

import "fmt"

func ExampleChord_Notes() {
	fmt.Println(Chord{C4, E4, G4}.Notes())
	// Output: C4-E4-G4
}

func ExampleChord_String() {
	fmt.Println(Chord{C4, Ds4, G4})
	// Output: Cm
}

func ExampleChord_String_scale() {
	fmt.Println(C4.Scale(MajorChord))
	fmt.Println(C4.Scale(MinorChord))
	fmt.Println(C4.Scale(AugmentedChord))
	fmt.Println(C4.Scale(DiminishedChord))
	fmt.Println(C4.Scale(DominantSeventhChord))
	fmt.Println(C4.Scale(MinorSeventhChord))
	fmt.Println(C4.Scale(MajorSeventhChord))
	fmt.Println(C4.Scale(AugmentedMinorSeventhChord))
	fmt.Println(C4.Scale(DiminishedSeventhChord))
	fmt.Println(C4.Scale(HalfDiminishedSeventhChord))

	// Output:
	// C
	// Cm
	// Caug
	// Cdim
	// C7
	// Cm7
	// Cmaj7
	// Caug7
	// Cdim7
	// C7♭5
}
