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

func ExampleNotes_Ring() {
	c := Notes{C4, E4, G4}.Ring()
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	// Output:
	// C4
	// E4
	// G4
	// C4
	// E4
	// G4
	// C4
	// E4
	// G4
	// C4
}
