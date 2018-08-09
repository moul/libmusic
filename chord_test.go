package libmusic

import "fmt"

func ExampleChord_Notes() {
	fmt.Println(Chord{C4, E4, G4}.Notes())
	// Output: C4-E4-G4
}

func ExampleChord_String() {
	fmt.Println(Chord{C4, Ds4, G4})
	fmt.Println(C4.Chord(MinorChord))
	// Output:
	// Cm
	// Cm
}

func ExampleChord_Simplify() {
	chord := Chord{C4, E4, C5, G3, E6}
	fmt.Println(chord)
	fmt.Println(chord.Notes())
	fmt.Println(chord.Simplify())
	fmt.Println(chord.Simplify().Notes())
	// Output:
	// C
	// C4-E4-C5-G3-E6
	// C
	// C4-E4-G4
}

func ExampleChord_Permutations() {
	for _, perm := range C4.Chord(MajorChord).Permutations() {
		fmt.Printf("chord=%s, notes=%s\n", perm.Notes(), perm)
	}
	// Output:
	// chord=C4-E4-G4, notes=C
	// chord=E4-G4-C5, notes=C
	// chord=G4-C5-E5, notes=C
}

func ExampleChord_String_scale() {
	fmt.Println(C4.Chord(MajorChord))
	fmt.Println(C4.Chord(MinorChord))
	fmt.Println(C4.Chord(AugmentedChord))
	fmt.Println(C4.Chord(DiminishedChord))
	fmt.Println(C4.Chord(DominantSeventhChord))
	fmt.Println(C4.Chord(MinorSeventhChord))
	fmt.Println(C4.Chord(MajorSeventhChord))
	fmt.Println(C4.Chord(AugmentedMinorSeventhChord))
	fmt.Println(C4.Chord(DiminishedSeventhChord))
	fmt.Println(C4.Chord(HalfDiminishedSeventhChord))

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

func ExampleNote_Chord() {
	fmt.Println(C4.Chord(MajorChord).Notes())
	fmt.Println(C4.Chord(MinorChord).Notes())
	fmt.Println(C4.Chord(AugmentedChord).Notes())
	fmt.Println(C4.Chord(DiminishedChord).Notes())
	fmt.Println(C4.Chord(DominantSeventhChord).Notes())
	fmt.Println(C4.Chord(MinorSeventhChord).Notes())
	fmt.Println(C4.Chord(MajorSeventhChord).Notes())
	fmt.Println(C4.Chord(AugmentedMinorSeventhChord).Notes())
	fmt.Println(C4.Chord(DiminishedSeventhChord).Notes())
	fmt.Println(C4.Chord(HalfDiminishedSeventhChord).Notes())

	// Output:
	// C4-E4-G4
	// C4-D♯4-G4
	// C4-E4-G♯4
	// C4-D♯4-F♯4
	// C4-E4-G4-A♯4
	// C4-D♯4-G4-A♯4
	// C4-E4-G4-B4
	// C4-E4-G♯4-A♯4
	// C4-D♯4-F♯4-A4
	// C4-D♯4-F♯4-A♯4
}
