package libmusic

import "fmt"

func ExampleNote_Scale() {
	fmt.Println(C4.Scale(MajorScale))
	fmt.Println(C4.Scale(HemiPentatonicScale))
	fmt.Println(C4.Scale(MinorBluesScale))
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

func ExampleScale_Chord() {
	scales := []Notes{
		C4.Scale(MajorScale),
		C4.Scale(HemiPentatonicScale),
		C4.Scale(MinorBluesScale),
		C4.Scale(MajorPentatonicScale),
		C4.Scale(ChromaticScale),
		C4.Scale(SpanishScale),
	}
	for _, scale := range scales {
		fmt.Println(scale.Chord().Notes(), " - ", scale.Chord(), " - ", scale)
	}
	// Output:
	// C4-E4-G4  -  C  -  C4-D4-E4-F4-G4-A4-B4
	// C4-F4-G♯4  -  Fm  -  C4-C♯4-F4-G4-G♯4
	// C4-F4-G4  -  C4-F4-G4  -  C4-D♯4-F4-F♯4-G4-A♯4
	// C4-E4-A4  -  Am  -  C4-D4-E4-G4-A4
	// C4-D4-E4  -  C4-D4-E4  -  C4-C♯4-D4-D♯4-E4-F4-F♯4-G4-G♯4-A4-A♯4-B4
	// C4-E4-G4  -  C  -  C4-C♯4-E4-F4-G4-G♯4-A♯4
}
