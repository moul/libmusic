package libmusic

import (
	"fmt"
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestNote(t *testing.T) {
	c.Convey("Testing Note", t, func() {
		c.Convey("Testing Note.Add", func() {
			c.So(Note{letter: C, octave: 3}.Add(1*Semitone).Equal(Note{letter: Cs, octave: 3}), c.ShouldBeTrue)
			c.So(Note{letter: C, octave: 3}.Add(2*Semitone).Equal(Note{letter: D, octave: 3}), c.ShouldBeTrue)
			c.So(Note{letter: C, octave: 3}.Add(3*Semitone).Equal(Note{letter: Ds, octave: 3}), c.ShouldBeTrue)
			c.So(Note{letter: C, octave: 3}.Add(4*Semitone).Equal(Note{letter: E, octave: 3}), c.ShouldBeTrue)
			c.So(Note{letter: C, octave: 3}.Add(5*Semitone).Equal(Note{letter: F, octave: 3}), c.ShouldBeTrue)
			c.So(Note{letter: C, octave: 3}.Add(6*Semitone).Equal(Note{letter: Fs, octave: 3}), c.ShouldBeTrue)
			c.So(Note{letter: C, octave: 3}.Add(7*Semitone).Equal(Note{letter: G, octave: 3}), c.ShouldBeTrue)
			c.So(Note{letter: C, octave: 3}.Add(8*Semitone).Equal(Note{letter: Gs, octave: 3}), c.ShouldBeTrue)
			c.So(Note{letter: C, octave: 3}.Add(9*Semitone).Equal(Note{letter: A, octave: 3}), c.ShouldBeTrue)
			c.So(Note{letter: C, octave: 3}.Add(10*Semitone).Equal(Note{letter: As, octave: 3}), c.ShouldBeTrue)
			c.So(Note{letter: C, octave: 3}.Add(11*Semitone).Equal(Note{letter: B, octave: 3}), c.ShouldBeTrue)
			c.So(Note{letter: C, octave: 3}.Add(12*Semitone).Equal(Note{letter: C, octave: 4}), c.ShouldBeTrue)
			c.So(Note{letter: C, octave: 3}.Add(13*Semitone).Equal(Note{letter: Cs, octave: 4}), c.ShouldBeTrue)
			c.So(Note{letter: C, octave: 3}.Add(30*Semitone).Equal(Note{letter: Fs, octave: 5}), c.ShouldBeTrue)
			c.So(Note{letter: C, octave: 3}.Add(1*Tone).Equal(Note{letter: D, octave: 3}), c.ShouldBeTrue)
			c.So(Note{letter: C, octave: 3}.Add(1*Octave).Equal(Note{letter: C, octave: 4}), c.ShouldBeTrue)
			c.So(Note{letter: C, octave: 3}.Add(-1*Semitone).Equal(Note{letter: B, octave: 2}), c.ShouldBeTrue)
		})
		c.Convey("Testing Note.String", func() {
			c.So(fmt.Sprintf("%s", C4), c.ShouldEqual, "C4")
			c.So(fmt.Sprintf("%s", Cs4), c.ShouldEqual, "C♯4")
			c.So(fmt.Sprintf("%s", D4), c.ShouldEqual, "D4")
			c.So(fmt.Sprintf("%s", Ds4), c.ShouldEqual, "D♯4")
			c.So(fmt.Sprintf("%s", E4), c.ShouldEqual, "E4")
			c.So(fmt.Sprintf("%s", F4), c.ShouldEqual, "F4")
			c.So(fmt.Sprintf("%s", Fs4), c.ShouldEqual, "F♯4")
			c.So(fmt.Sprintf("%s", G4), c.ShouldEqual, "G4")
			c.So(fmt.Sprintf("%s", Gs4), c.ShouldEqual, "G♯4")
			c.So(fmt.Sprintf("%s", A4), c.ShouldEqual, "A4")
			c.So(fmt.Sprintf("%s", As4), c.ShouldEqual, "A♯4")
			c.So(fmt.Sprintf("%s", B4), c.ShouldEqual, "B4")
		})
	})
}
