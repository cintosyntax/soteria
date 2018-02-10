package models

// Password defines a valid data structure that encapsulates a password string and
// and its validity
type Password struct {
	String string
	errors []string
}

// BuildPassword returns a Password struct with its value as the string provided
func BuildPassword(pw string) *Password {
	return &Password{
		String: pw,
		errors: []string{},
	}
}

// GetErrors returns the private collection of errors.
func (p *Password) GetErrors() []string {
	return p.errors
}

// Valid returns true if there are no errors on the string
func (p *Password) Valid() bool {
	if len(p.GetErrors()) > 0 {
		return false
	}
	return true
}
