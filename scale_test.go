package libmusic

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestScale(t *testing.T) {
	c.Convey("Testing Scale", t, func() {
		c.So(C4.Scale(MajorChord).Equal(Chord{C4, E4, G4}), c.ShouldBeTrue)
	})
}
