package terrapi

import (
	"errors"
	"reflect"
)

var ErrNoRegisteredActions = errors.New("no registered actions for terraform to execute")

type InvalidActionError struct {
	Type reflect.Type
}

func (e *InvalidActionError) Error() string {
	if e.Type == nil {
		return "action: (nil)"
	}

	return "action: " + e.Type.String()
}

type TerraformBinaryError struct {
	Path string
	Msg  string
}

func (e *TerraformBinaryError) Error() string {
	return e.Path + ": " + e.Msg
}

type FunctionalOptionError struct {
	Opt string
	//TODO: Read on how to encapsulate/wrap original error like `%w` do in fmt pkg.
	Err error
}

func (e *FunctionalOptionError) Error() string {
	return "error processing functional option" + e.Opt
}
