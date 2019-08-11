package output

import (
    "errors"
    "github.com/olehan/marker/validation"
    "reflect"
    "strings"
)

const (
    typeCastError = "interface conversion: output.Output is"
)

var (
    emptyErrorValue = reflect.ValueOf(errors.New(""))
)

// RecoverBadTypeCast recovers only bad interface type casts
// and executes the given callback in that scenario.
func RecoverBadTypeCast(callback func(err error)) {
    if err := toError(recover()); err != nil {
        if strings.Index(err.Error(), typeCastError) > -1 {
            callback(err)
        } else {
            panic(err)
        }
    }
}

// toError checks if the given value is an error and returns it
// if it is an error.
func toError(maybeError interface{}) error {
    err := validation.ValidateTwoValues(reflect.ValueOf(maybeError), emptyErrorValue)
    if err != nil {
        return nil
    }
    return maybeError.(error)
}
