package models

import "github.com/cintosyntax/soteria/pkg/validators"

// Password defines a valid data structure that encapsulates a password string and
// and its validity
type Password struct {
	String        string
	errorMessages []string
}

// BuildPassword returns a Password struct with its value as the string provided
func BuildPassword(pw string) *Password {
	return &Password{
		String:        pw,
		errorMessages: []string{},
	}
}

func (p *Password) Validate(validator validators.PasswordValidator) {
	valid, errMsgs := validator.Validate(p.String)
	if valid == false {
		p.errorMessages = errMsgs
	}
}

// GetErrorMessages returns the private collection of errors.
func (p *Password) GetErrorMessages() []string {
	return p.errorMessages
}

// Valid returns true if there are no errors on the string
func (p *Password) Valid() bool {
	if len(p.GetErrorMessages()) > 0 {
		return false
	}
	return true
}
