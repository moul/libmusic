package libmusic

import "fmt"

func ExampleChord_NotesString() {
	fmt.Println(Chord{C4, E4, G4}.Notes())
	// Output: C4-E4-G4
}
