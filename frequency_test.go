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

func ExampleNote_ByFrequency() {
	targets := []Hz{440, 880, 1, 666, 42, 4242.4242, 1000, 1234}
	for _, target := range targets {
		note := NoteByFrequency(target)
		freq := note.Frequency()
		diff := Hz(target - freq).Relative()
		fmt.Printf("target=%s, note=%s, note.freq=%s, diff=%s\n", freq.String(), note.String(), freq, diff)
	}
	// Output:
	// target=440㎐, note=A4, note.freq=440㎐, diff=0㎐
	// target=880㎐, note=A5, note.freq=880㎐, diff=0㎐
	// target=1.022㎐, note=C-4, note.freq=1.022㎐, diff=0.022㎐
	// target=659.255㎐, note=E5, note.freq=659.255㎐, diff=6.745㎐
	// target=41.203㎐, note=E1, note.freq=41.203㎐, diff=0.797㎐
	// target=4186.009㎐, note=C8, note.freq=4186.009㎐, diff=56.415㎐
	// target=987.767㎐, note=B5, note.freq=987.767㎐, diff=12.233㎐
	// target=1244.508㎐, note=D♯6, note.freq=1244.508㎐, diff=10.508㎐
}
