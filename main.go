package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	pws, err := ExtractPasswordsFromIO()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello, %s!", pws[0])
}

func ExtractPasswordsFromIO() ([]*Password, error) {
	stdinFileInfo, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}

	if (stdinFileInfo.Mode() & os.ModeNamedPipe) == 0 {
		return nil, errors.New("No passwords to validate given")
	}

	stdinContent, readErr := ioutil.ReadAll(os.Stdin)
	if readErr != nil {
		return nil, readErr
	}

	passwordLines := strings.Split(string(stdinContent), "\n")
	pws := make([]*Password, len(passwordLines))
	for i, pw := range passwordLines {
		pws[i] = &Password{
			String: pw,
		}
	}

	return pws, nil
}
