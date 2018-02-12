package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/cintosyntax/soteria/pkg/models"
	"github.com/cintosyntax/soteria/pkg/validators"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	validatorPtr := flag.String("validator", "nist", "defines which validator to use")
	commonPasswordFilePtr := flag.String("cpf", "", "defines a file that contains weak passwords")
	flag.Parse()

	passwordValidator, pvLoadErr := LoadPasswordValidator(*validatorPtr)
	if pvLoadErr != nil {
		panic(pvLoadErr)
	}

	commonPasswords, cpLoadErr := LoadCommonPasswordsFile(*commonPasswordFilePtr)
	if cpLoadErr != nil {
		panic(cpLoadErr)
	}

	passwordValidator.AddToCommonBlackList(commonPasswords)

	passwords, pwExtractErr := ExtractPasswordsFromIO()
	if pwExtractErr != nil {
		panic(pwExtractErr)
	}

	for _, pw := range passwords {
		pw.Validate(passwordValidator)
		if pw.Valid() == false {
			pwInvalidErrMsg := buildErrorDisplayMessage(pw)
			fmt.Println(pwInvalidErrMsg)
		} else {
			fmt.Printf("%s: OK\n", pw.String)
		}
	}

}

// Application specific functions ---

func buildErrorDisplayMessage(pw *models.Password) string {
	errorString := strings.Join(pw.GetErrorMessages(), ", ")

	formattedPassword := replaceIllegalCharacters(pw.String, "*")

	return fmt.Sprintf("%s: %s", formattedPassword, errorString)
}

func replaceIllegalCharacters(str string, rs string) string {
	var re = regexp.MustCompile(`([^\x00-\x7F])`)
	s := re.ReplaceAllString(str, rs)
	return s
}
func LoadCommonPasswordsFile(fileName string) ([]string, error) {
	if fileName == "" {
		return []string{}, nil
	}

	fileData, readErr := ioutil.ReadFile(fileName)
	if readErr != nil {
		return nil, readErr
	}

	passwordsString := string(fileData)
	passwords := strings.Split(passwordsString, "\n")
	return passwords, nil
}

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
		return nil, err
	}

	if (stdinFileInfo.Mode() & os.ModeNamedPipe) == 0 {
		return nil, errors.New("no passwords to validate given")
	}

	stdinContent, readErr := ioutil.ReadAll(os.Stdin)
	if readErr != nil {
		return nil, readErr
	}

	rawLines := strings.Split(string(stdinContent), "\n")
	passwordLines := []string{}
	for _, pw := range rawLines {
		if pw != "" {
			passwordLines = append(passwordLines, pw)
		}
	}

	pws := make([]*models.Password, len(passwordLines))
	for i, pw := range passwordLines {
		pws[i] = models.BuildPassword(pw)
	}

	return pws, nil
}
