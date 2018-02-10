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
			So(err.Error(), ShouldEqual, "no passwords to validate given")
		})

		Convey("it returns no passwords", func() {
			So(pws, ShouldBeEmpty)
		})
	})
}

func TestLoadPasswordValidator(t *testing.T) {
	Convey("when no validator is specified", t, func() {
		_, err := LoadPasswordValidator("")
		Convey("it should return an error", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "validator must be specified")
		})
	})

	Convey("when a unknown validator is specified", t, func() {
		_, err := LoadPasswordValidator("spythisvalidator")
		Convey("it should return an error", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "unknown validator specified")
		})
	})
}
