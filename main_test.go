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

func TestLoadCommonPasswordsFile(t *testing.T) {
	Convey("when no file is specified", t, func() {
		commonPasswords, err := LoadCommonPasswordsFile("")

		Convey("it should return no error", func() {
			So(err, ShouldBeNil)
		})
		Convey("it should return a empty slice of strings", func() {
			So(commonPasswords, ShouldBeEmpty)
		})
	})

	Convey("when a invalid file is specified", t, func() {
		_, err := LoadCommonPasswordsFile("no path here!")
		Convey("it should return an error", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "open no path here!: no such file or directory")
		})
	})
}
