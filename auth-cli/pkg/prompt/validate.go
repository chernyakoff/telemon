package prompt

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/dongri/phonenumber"
)

type Validate struct{}

func (v Validate) PhoneNumber(input string) error {
	number := phonenumber.Parse(input, "RU")
	if number == "" {
		return errors.New("неправильно указан телефон")
	}
	return nil
}

func (v Validate) AppId(input string) error {
	error := errors.New("неправильно указан app id")
	if len(input) < 6 || len(input) > 15 {
		return error
	}
	if _, err := strconv.Atoi(input); err != nil {
		return error
	}
	return nil
}

func (v Validate) AppHash(input string) error {
	isValid := regexp.MustCompile(`^[0-9a-z]+$`).MatchString
	if len(input) != 32 || !isValid(input) {
		return errors.New("неправильно указан app hash")
	}
	return nil
}

func (v Validate) EmptyInput(input string) error {
	if len(strings.TrimSpace(input)) < 1 {
		return errors.New("this input must not be empty")
	}
	return nil
}

func (v Validate) IntegerNumberInput(input string) error {
	_, err := strconv.ParseInt(input, 0, 64)
	if err != nil {
		return errors.New("invalid number")
	}
	return nil
}
