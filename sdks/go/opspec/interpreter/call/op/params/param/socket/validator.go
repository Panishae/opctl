package socket

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./fakeValidator.go --fake-name FakeValidator ./ Validator

import (
	"errors"
	"github.com/opctl/opctl/sdks/go/model"
)

type Validator interface {
	Validate(
		value *model.Value,
	) []error
}

func NewValidator() Validator {
	return _validator{}
}

type _validator struct {
}

// Validate validates a value against a string parameter
func (vdt _validator) Validate(
	value *model.Value,
) []error {
	// handle no value passed
	if nil == value || nil == value.Socket {
		return []error{errors.New("socket required")}
	}

	return []error{}
}
