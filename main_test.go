package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExtractPasswordsFromIO(t *testing.T) {

	Convey("when nothing is in IO", t, func() {
		pws, err := ExtractPasswordsFromIO()

		Convey("it returns a no password given error", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "No passwords to validate given")
		})

		Convey("it returns no passwords", func() {
			So(pws, ShouldBeEmpty)
		})
	})
}
