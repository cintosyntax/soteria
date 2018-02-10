package models

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBuildPassword(t *testing.T) {
	Convey("when given a password", t, func() {
		somePassword := "goats"
		pw := BuildPassword(somePassword)

		Convey("it a Password model struct", func() {
			So(pw, ShouldHaveSameTypeAs, &Password{})
		})
		Convey("it stores the passwords the string", func() {
			So(pw.String, ShouldEqual, somePassword)
		})
	})
}

func TestPassword(t *testing.T) {
	Convey("Valid", t, func() {
		Convey("with errors it should be false", func() {
			pw := &Password{
				String:        "goat",
				errorMessages: []string{"bad characters", "bad things"},
			}
			So(pw.Valid(), ShouldBeFalse)
		})

		Convey("without errors it should be true", func() {
			pw := &Password{
				String: "goat",
			}

			So(pw.Valid(), ShouldBeTrue)
		})
	})
}
