package libmusic

import "fmt"

func ExampleChord_Notes() {
	fmt.Println(Chord{C4, E4, G4}.Notes())
	// Output: C4-E4-G4
}

func ExampleChord_String() {
	fmt.Println(Chord{C4, Ds4, G4})
	fmt.Println(C4.Scale(MinorChord, 1))
	fmt.Println(C4.Scale(MinorChord, 5))
	// Output:
	// Cm
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
	for _, perm := range C4.Scale(MajorChord, 1).Permutations() {
		fmt.Printf("chord=%s, notes=%s\n", perm.Notes(), perm)
	}
	// Output:
	// chord=C4-E4-G4, notes=C
	// chord=E4-G4-C5, notes=C
	// chord=G4-C5-E5, notes=C
}

func ExampleChord_String_scale() {
	fmt.Println(C4.Scale(MajorChord, 1))
	fmt.Println(C4.Scale(MinorChord, 1))
	fmt.Println(C4.Scale(AugmentedChord, 1))
	fmt.Println(C4.Scale(DiminishedChord, 1))
	fmt.Println(C4.Scale(DominantSeventhChord, 1))
	fmt.Println(C4.Scale(MinorSeventhChord, 1))
	fmt.Println(C4.Scale(MajorSeventhChord, 1))
	fmt.Println(C4.Scale(AugmentedMinorSeventhChord, 1))
	fmt.Println(C4.Scale(DiminishedSeventhChord, 1))
	fmt.Println(C4.Scale(HalfDiminishedSeventhChord, 1))

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

func ExampleNote_Scale() {
	fmt.Println(C4.Scale(MajorChord, 1).Notes())
	fmt.Println(C4.Scale(MinorChord, 1).Notes())
	fmt.Println(C4.Scale(AugmentedChord, 1).Notes())
	fmt.Println(C4.Scale(DiminishedChord, 1).Notes())
	fmt.Println(C4.Scale(DominantSeventhChord, 1).Notes())
	fmt.Println(C4.Scale(MinorSeventhChord, 1).Notes())
	fmt.Println(C4.Scale(MajorSeventhChord, 1).Notes())
	fmt.Println(C4.Scale(AugmentedMinorSeventhChord, 1).Notes())
	fmt.Println(C4.Scale(DiminishedSeventhChord, 1).Notes())
	fmt.Println(C4.Scale(HalfDiminishedSeventhChord, 1).Notes())

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

func ExampleNote_Scale_big() {
	fmt.Println(C4.Scale(MajorChord, 5).Notes())

	// Output: C4-E4-G4-C5-E5-G5-C6-E6-G6-C7-E7-G7-C8-E8-G8
}
