# Soteria
[![Build Status](https://travis-ci.org/cintosyntax/soteria.svg?branch=master)](https://travis-ci.org/cintosyntax/soteria)

A Golang password validation utility used to validate a list of passwords against the [NIST](https://www.nist.gov/) password standards.

# Requirements

If you haven't already installed Go, follow the steps [here](https://golang.org/doc/install). This project requires Go to be installed.

# Installation

If you haven't already installed Go, follow the steps [here](https://golang.org/doc/install). This project requires Go to be installed.

```bash
cd $GOPATH/src/
git clone https://github.com/cintosyntax/soteria.git
go build
```

# Usage

Basic usage. You must pipe the a newline delimited document of passwords.
```bash
cat password_list.txt | ./soteria -validator=nist -cfp=weak_password_list.txt

# *_*: Too short (8 character minimum), Contains illegal characters
# l*ng: Too short (8 character minimum), Contains illegal characters
# password: Too common
```

### Options

- validator - this flag option can change the validator to use. By default this is "nist". Another option is "lax", which uses a validator that accepts everything.
- cfp - this flag defines the directory to a common password list that can be used by the selected validator to ignore them. By default this is not set.
