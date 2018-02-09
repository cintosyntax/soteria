package main

// LaxPasswordValidator describes a validator that accepts ANY password. Warning
// never use this in production!
type LaxPasswordValidator struct{}

// Validate returns true no matter what password you give.
func (lpw *LaxPasswordValidator) Validate(pw string) (bool, []string) {
	return true, []string{}
}
