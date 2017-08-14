package models

import (
	"fmt"
	"strings"
)

// Validation holds information about whether the current model
// is valid or not.  It is designed to be embedded into structs
// that need post-operation validation.
//
// swagger: model
type Validation struct {
	// Validated tracks whether or not the model has been validated.
	// read only: true
	Validated bool `json:",omitempty"`
	// Available tracks whether or not the model passed validation.
	// read only: true
	Available bool `json:",omitempty"`
	// If there are any errors in the validation process, they will be
	// available here.
	// read only: true
	Errors []string `json:",omitempty"`
}

func (v *Validation) ClearValidation() {
	v.Validated = false
	v.Available = false
	v.Errors = []string{}
}

func (v *Validation) Errorf(fmtStr string, args ...interface{}) {
	v.Available = false
	if v.Errors == nil {
		v.Errors = []string{}
	}
	v.Errors = append(v.Errors, fmt.Sprintf(fmtStr, args...))
}

func (v *Validation) AddError(err error) {
	if err != nil {
		if v.Errors == nil {
			v.Errors = []string{}
		}
		v.Errors = append(v.Errors, err.Error())
	}
}

func (v *Validation) HasError() error {
	if len(v.Errors) == 0 {
		return nil
	}
	return v
}

func (v *Validation) Useable() bool {
	return v.Validated
}

func (v *Validation) SetValid() bool {
	v.Validated = v.Validated || len(v.Errors) == 0
	return v.Validated
}

func (v *Validation) SetAvailable() bool {
	v.Available = v.Available || len(v.Errors) == 0
	return v.Available
}

func (v *Validation) Error() string {
	return strings.Join(v.Errors, "\n")
}

func (v *Validation) MakeError(code int, errType string, obj Model) error {
	if len(v.Errors) == 0 {
		return nil
	}
	return &Error{
		Object:        obj,
		Code:          code,
		Type:          errType,
		Messages:      v.Errors,
		containsError: true,
	}
}
