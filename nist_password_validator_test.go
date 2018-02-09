package soteria

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNistPasswordValidator(t *testing.T) {
	npv := BuildNistPasswordValidator()

	Convey("Validate", t, func() {
		Convey("when given a password less than 8 characters long", func() {
			shortPassword := "short"

			valid, errMsgs := npv.Validate(shortPassword)
			Convey("it returns false", func() {
				So(valid, ShouldBeFalse)
			})
			Convey("it returns a password too short error message", func() {
				So(errMsgs, ShouldContain, "Too short (8 character minimum)")
			})
		})

		Convey("when given a password that is over 64 characters long", func() {
			longPassword := make([]rune, 65)
			for i := range longPassword {
				// Set each character to the letter 'B' (ASCII code = 66)
				longPassword[i] = 66
			}

			valid, errMsgs := npv.Validate(string(longPassword))
			Convey("it says the password is invalid", func() {
				So(valid, ShouldBeFalse)
			})

			Convey("it returns a password too long error message", func() {
				So(errMsgs, ShouldContain, "Too long (64 character limit)")
			})
		})

		Convey("when given a password with non ASCII characters", func() {
			notAllASCII := "ಠ_ಠ password goats"

			valid, errMsgs := npv.Validate(notAllASCII)
			Convey("it says the password is invalid", func() {
				So(valid, ShouldBeFalse)
			})

			Convey("it returns a invalid characters given error", func() {
				So(errMsgs, ShouldContain, "Contains illegal characters")
			})
		})

		Convey("when given a common password", func() {
			commonPassword := "password"
			npv.AddToCommonBlackList([]string{commonPassword})

			valid, errMsgs := npv.Validate(commonPassword)
			Convey("it says the password is invalid", func() {
				So(valid, ShouldBeFalse)
			})

			Convey("it returns a password too common error", func() {
				So(errMsgs, ShouldContain, "Too common")
			})
		})

	})
}
