# Soteria
[![Build Status](https://travis-ci.org/cintosyntax/soteria.svg?branch=master)](https://travis-ci.org/cintosyntax/soteria)

A Golang password validation utility (named after the Greek goddess of safety) used to validate a list of passwords against the [NIST](https://www.nist.gov/) password standards. 


# Installation

If you haven't already installed Go, follow the steps [here](https://golang.org/doc/install). This project requires Go to be installed.

```bash
go get github.com/cintosyntax/soteria
cd $GOPATH/src/github.com/cintosyntax/soteria
go build
```

# Usage

Basic usage. You must pipe the a newline delimited document of passwords.
```bash
cat password_list.txt | ./soteria -validator=nist -cfp=weak_password_list.txt

# *_*: Too short (8 character minimum), Contains illegal characters
# tinyelephantmusic: OK
# l*ng: Too short (8 character minimum), Contains illegal characters
# password: Too common
# thisisagoodpassword: OK
```

#### Options
- validator - this flag option can change the validator to use. By default this is "nist". Another option is "lax", which uses a validator that accepts everything.
- cfp - this flag defines the directory to a common password list that can be used by the selected validator to ignore them. By default this is not set.

# Contributing (Optional)

This project uses glide as a package manager which handles installation of development dependencies such as goconvey (the testing library).

- Follow these steps to install glide [here](https://github.com/Masterminds/glide)
- In the project root directory run `glide install`
- After the dependencies have installed, you can run the tests with `go test`