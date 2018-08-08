package libmusic

import "fmt"

func ExampleNote_Frequency() {
	// low-frequency notes
	fmt.Println(C0, C0.Frequency())
	fmt.Println(D0, D0.Frequency())
	fmt.Println(E0, E0.Frequency())
	fmt.Println(F0, F0.Frequency())
	fmt.Println(G0, G0.Frequency())
	fmt.Println(A0, A0.Frequency())
	fmt.Println(B0, B0.Frequency())

	// octaves
	fmt.Println(A0, A0.Frequency())
	fmt.Println(A1, A1.Frequency())
	fmt.Println(A2, A2.Frequency())
	fmt.Println(A3, A3.Frequency())
	fmt.Println(A4, A4.Frequency())
	fmt.Println(A5, A5.Frequency())
	fmt.Println(A6, A6.Frequency())
	fmt.Println(A7, A7.Frequency())
	fmt.Println(A8, A8.Frequency())

	// Output:
	// C0 16.352㎐
	// D0 18.354㎐
	// E0 20.602㎐
	// F0 21.827㎐
	// G0 24.5㎐
	// A0 27.5㎐
	// B0 30.868㎐
	// A0 27.5㎐
	// A1 55㎐
	// A2 110㎐
	// A3 220㎐
	// A4 440㎐
	// A5 880㎐
	// A6 1760㎐
	// A7 3520㎐
	// A8 7040㎐
}
