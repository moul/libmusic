package libmusic

import "fmt"

func ExampleInterval_String() {
	fmt.Printf("%d: %s\n", Semitone, Semitone.String())
	fmt.Printf("%d: %s\n", Tone, Tone.String())
	fmt.Printf("%d: %s\n", MajorThird, MajorThird.String())
	fmt.Printf("%d: %s\n", Fourth, Fourth.String())
	fmt.Printf("%d: %s\n", Octave, Octave.String())
	// Output:
	// 1: m2
	// 2: M2
	// 4: M3
	// 5: P4
	// 12: P8
}
