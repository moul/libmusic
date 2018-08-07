package libmusic

import "fmt"

func ExampleNotes_String() {
	fmt.Println(Notes{C4, E4, G4})
	// Output: C4-E4-G4
}

func ExampleNotes_Intervals() {
	fmt.Println(Notes{C4, E4, G4}.Intervals())
	// Output: M3-P5
}
