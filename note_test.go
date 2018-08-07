package libmusic

import "fmt"

func ExampleNote_String() {
	fmt.Println(C4)
	fmt.Println(Cs4)
	fmt.Println(D4)
	fmt.Println(Ds4)
	fmt.Println(E4)
	fmt.Println(F4)
	fmt.Println(Fs4)
	fmt.Println(G4)
	fmt.Println(Gs4)
	fmt.Println(A4)
	fmt.Println(As4)
	fmt.Println(B4)
	// Output:
	// C4
	// C♯4
	// D4
	// D♯4
	// E4
	// F4
	// F♯4
	// G4
	// G♯4
	// A4
	// A♯4
	// B4
}

func ExampleNote_Add() {
	fmt.Println(C4.Add(1 * Semitone))
	fmt.Println(C4.Add(2 * Semitone))
	fmt.Println(C4.Add(3 * Semitone))
	fmt.Println(C4.Add(4 * Semitone))
	fmt.Println(C4.Add(5 * Semitone))
	fmt.Println(C4.Add(6 * Semitone))
	fmt.Println(C4.Add(7 * Semitone))
	fmt.Println(C4.Add(8 * Semitone))
	fmt.Println(C4.Add(9 * Semitone))
	fmt.Println(C4.Add(10 * Semitone))
	fmt.Println(C4.Add(11 * Semitone))
	fmt.Println(C4.Add(12 * Semitone))
	fmt.Println(C4.Add(13 * Semitone))
	fmt.Println(C4.Add(30 * Semitone))
	fmt.Println(C4.Add(3 * Tone))
	fmt.Println(C4.Add(3 * Octave))
	fmt.Println(C4.Add(-1*Octave - 2*Tone - 3*Semitone))
	// Output:
	// C♯4
	// D4
	// D♯4
	// E4
	// F4
	// F♯4
	// G4
	// G♯4
	// A4
	// A♯4
	// B4
	// C5
	// C♯5
	// F♯6
	// F♯4
	// C7
	// F2
}
