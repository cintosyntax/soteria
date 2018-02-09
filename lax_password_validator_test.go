package soteria

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLaxPasswordValidator(t *testing.T) {
	lpv := &LaxPasswordValidator{}
	Convey("Validate", t, func() {
		Convey("when given an empty string", func() {
			badPassword := ""
			bool, errMsgs := lpv.Validate(badPassword)

			Convey("it returns true", func() {
				So(bool, ShouldBeTrue)
			})
			Convey("it returns no error messages", func() {
				So(errMsgs, ShouldBeEmpty)
			})
		})
	})
}
