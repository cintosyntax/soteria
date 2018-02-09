package soteria

// Password defines a valid data structure that encapsulates a password string and
// and its validity
type Password struct {
	String    string
	Errors    string
	Validator *PasswordValidator
}
