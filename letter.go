package libmusic

type Letter int

const (
	C Letter = iota
	Cs
	D
	Ds
	E
	F
	Fs
	G
	Gs
	A
	As
	B
)

var letters = []string{"C", "C♯", "D", "D♯", "E", "F", "F♯", "G", "G♯", "A", "A♯", "B"}

func (l Letter) String() string {
	return letters[l]
}
