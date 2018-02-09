package main

import "fmt"

// NistPasswordValidator implements as PasswordValidator which determines if
// a given password string meets NIST requirements.
type NistPasswordValidator struct {
	commonBlackList map[string]bool
}

// BuildNistPasswordValidator returns a pointer to a NistPasswordValidator that
// contains logic for validating if a password string is valid according to NIST
// standards.
func BuildNistPasswordValidator() *NistPasswordValidator {
	return &NistPasswordValidator{
		commonBlackList: make(map[string]bool),
	}
}

// Validate takes the string provided and returns if it valid or invalid. If
// the record is invalid it will return a slice of strings that indicate the
// reason.
//
// TODO - use go functions to add performance boost
func (nsv *NistPasswordValidator) Validate(pw string) (bool, []string) {
	if valid, errMsg := nsv.validateLength(pw); valid == false {
		return false, []string{errMsg}
	}

	if valid, errMsg := nsv.validateNotCommon(pw); valid == false {
		return false, []string{errMsg}
	}

	if valid, errMsg := nsv.validateASCII(pw); valid == false {
		return false, []string{errMsg}
	}

	return true, []string{}
}

// AddToCommonBlackList stores records of blacklisted passwords that are common.
func (nsv *NistPasswordValidator) AddToCommonBlackList(pws []string) {
	for _, pw := range pws {
		if nsv.commonBlackList[pw] == false {
			// No record found.. add it into the data structure
			nsv.commonBlackList[pw] = true
		}
	}
}

func (nsv *NistPasswordValidator) validateLength(s string) (bool, string) {
	minLength := 8
	maxLength := 64

	if len(s) < minLength {
		return false, fmt.Sprintf("Too short (%v character minimum)", minLength)
	} else if len(s) > maxLength {
		return false, fmt.Sprintf("Too long (%v character limit)", maxLength)
	}

	return true, ""
}

func (nsv *NistPasswordValidator) validateASCII(s string) (bool, string) {
	for _, chr := range s {
		if chr < 32 || chr > 126 {
			return false, "Contains illegal characters"
		}
	}

	return true, ""
}

func (nsv *NistPasswordValidator) validateNotCommon(s string) (bool, string) {
	if nsv.commonBlackList[s] == true {
		return false, "Too common"
	}

	return true, ""
}
