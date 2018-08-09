package libmusic

import "fmt"

func ExampleNote_Scale() {
	fmt.Println(C4.Scale(MajorScale))
	fmt.Println(C4.Scale(HemiPentatonicScale))
	fmt.Println(C4.Scale(BluesScale))
	fmt.Println(C4.Scale(MajorPentatonicScale))
	fmt.Println(C4.Scale(ChromaticScale))
	fmt.Println(C4.Scale(SpanishScale))
	// Output:
	// C4-D4-E4-F4-G4-A4-B4
	// C4-C♯4-F4-G4-G♯4
	// C4-D♯4-F4-F♯4-G4-A♯4
	// C4-D4-E4-G4-A4
	// C4-C♯4-D4-D♯4-E4-F4-F♯4-G4-G♯4-A4-A♯4-B4
	// C4-C♯4-E4-F4-G4-G♯4-A♯4
}
