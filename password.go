package main

// Password defines a valid data structure that encapsulates a password string and
// and its validity
type Password struct {
	String string
	errors []string
}

// PasswordValidator defines the public interface that describes a valid struct
// that is capable of validating passwords.
type PasswordValidator interface {
	Validate(string) (bool []string)
}

func BuildPassword(pw string) *Password {
	return &Password{
		String: pw,
		errors: []string{},
	}
}

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
