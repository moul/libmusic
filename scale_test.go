package libmusic

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestScale(t *testing.T) {
	c.Convey("Testing Scale", t, func() {
		c.So(C4.Scale(MajorChord).Equal(Chord{C4, E4, G4}), c.ShouldBeTrue)
		c.So(C4.Scale(MinorChord).Equal(Chord{C4, Ds4, G4}), c.ShouldBeTrue)

		/*
			c.So(C4.Scale(AugmentedChord).Equal(Chord{C4}), c.ShouldBeTrue)
			c.So(C4.Scale(DiminishedChord).Equal(Chord{C4}), c.ShouldBeTrue)
			c.So(C4.Scale(DominantSeventhChord).Equal(Chord{C4}), c.ShouldBeTrue)
			c.So(C4.Scale(MinorSeventhChord).Equal(Chord{C4}), c.ShouldBeTrue)
			c.So(C4.Scale(MajorSeventhChord).Equal(Chord{C4}), c.ShouldBeTrue)
			c.So(C4.Scale(AugmentedMinorSeventhChord).Equal(Chord{C4}), c.ShouldBeTrue)
			c.So(C4.Scale(DiminishedSeventhChord).Equal(Chord{C4}), c.ShouldBeTrue)
			c.So(C4.Scale(HalfDiminishedSeventhChord).Equal(Chord{C4}), c.ShouldBeTrue)
		*/
	})
}
