package libmusic

import "fmt"

func ExampleNote_Scale() {
	fmt.Println(C4.Scale(MajorChord).Notes())
	fmt.Println(C4.Scale(MinorChord).Notes())
	fmt.Println(C4.Scale(AugmentedChord).Notes())
	fmt.Println(C4.Scale(DiminishedChord).Notes())
	fmt.Println(C4.Scale(DominantSeventhChord).Notes())
	fmt.Println(C4.Scale(MinorSeventhChord).Notes())
	fmt.Println(C4.Scale(MajorSeventhChord).Notes())
	fmt.Println(C4.Scale(AugmentedMinorSeventhChord).Notes())
	fmt.Println(C4.Scale(DiminishedSeventhChord).Notes())
	fmt.Println(C4.Scale(HalfDiminishedSeventhChord).Notes())

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
	// C4-D♯4-F♯4-A4
}
