package main

import (
	"errors"
	"flag"
	"github.com/cintosyntax/soteria/pkg/models"
	"github.com/cintosyntax/soteria/pkg/validators"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Parse Arguments
	// Defines flags
	flag.String("validator", "", "defines which validator to use")
	flag.Parse()

	// LoadPasswordValidator

	_, err := ExtractPasswordsFromIO()
	if err != nil {
		// Panic if there was something wrong with the passwords provided in IO.
		panic(err)
	}

	// Read the CommonPasswords in the 2nd parameters
	// Build the NistPasswordValidator

}

// Application specific functions ---

func LoadPasswordValidator(pvName string) (validators.PasswordValidator, error) {
	if pvName == "" {
		return nil, errors.New("validator must be specified")
	}

	pvName = strings.ToLower(pvName)

	switch pvName {
	case "nist":
		return validators.BuildNistPasswordValidator(), nil
	case "lax":
		return validators.BuildLaxPasswordValidator(), nil
	default:
		return nil, errors.New("unknown validator specified")
	}
}

func ExtractPasswordsFromIO() ([]*models.Password, error) {
	stdinFileInfo, err := os.Stdin.Stat()
	if err != nil {
		// Add Test coverage here
		return nil, err
	}

	if (stdinFileInfo.Mode() & os.ModeNamedPipe) == 0 {
		return nil, errors.New("no passwords to validate given")
	}

	stdinContent, readErr := ioutil.ReadAll(os.Stdin)
	if readErr != nil {
		// Add Test Coverage here
		return nil, readErr
	}

	passwordLines := strings.Split(string(stdinContent), "\n")
	pws := make([]*models.Password, len(passwordLines))
	for i, pw := range passwordLines {
		pws[i] = models.BuildPassword(pw)
	}

	return pws, nil
}
