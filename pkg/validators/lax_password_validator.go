package validators

// LaxPasswordValidator describes a validator that accepts ANY password. Warning
// never use this in production!
type LaxPasswordValidator struct{}

// BuildLaxPasswordValidator returns a LaxPasswordValidator with it's defaults.
func BuildLaxPasswordValidator() *LaxPasswordValidator {
	return &LaxPasswordValidator{}
}

// Validate returns true no matter what password you give.
func (lpw *LaxPasswordValidator) Validate(pw string) (bool, []string) {
	return true, []string{}
}

// AddToCommonBlackList does not do anything. LaxPasswordValidator does not care
// about a blacklist.
func (lpw *LaxPasswordValidator) AddToCommonBlackList(pws []string) {
	return
}
