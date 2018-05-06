package main

import "fmt"

func emptyValuesInSlice(params ...string) (bool, []string, error) {
	var hasEmpty bool
	params, err := forEach(func(uris string) bool {
		notEmpty := len(uris) > 0
		hasEmpty = !notEmpty
		return notEmpty
	}, params...)
	return hasEmpty, params, err
}

func forEach(verifier func(param string) bool, params ...string) ([]string, error) {
	var verifiedParams []string
	var err error
	for index, param := range params {
		if verifier(param) {
			verifiedParams = append(verifiedParams, param)
		} else {
			err = fmt.Errorf(`parameter "%s" on %d position is invalid`, param, index)
		}
	}
	return verifiedParams, err
}
