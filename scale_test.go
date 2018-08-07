package libmusic

import "fmt"

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
