package soteria

// PasswordValidator defines the public interface that describes a valid struct
// that is capable of validating passwords.
type PasswordValidator interface {
	Validate(string) (bool []string)
}
