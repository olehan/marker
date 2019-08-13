package validation

import (
    "reflect"
    "testing"
)

func TestValidateTwoValues(t *testing.T) {
    type structTargetType struct {
        privateField int
        PublicField  string
        Node         *structTargetType
    }

    stringTarget := "stringValue"
    structTarget := &structTargetType{}
    structTargetValue := reflect.ValueOf(structTarget)
    stringTargetValue := reflect.ValueOf(stringTarget)

    err := ValidateTwoValues(stringTargetValue, reflect.ValueOf(""))
    if err != nil {
        t.FailNow()
    }

    err = ValidateTwoValues(stringTargetValue, reflect.ValueOf(0))
    if err == nil || err.Error() != valuesNotMatchError {
        t.FailNow()
    }

    err = ValidateTwoValues(structTargetValue, reflect.ValueOf(&structTargetType{}))
    if err != nil {
        t.FailNow()
    }

    err = ValidateTwoValues(structTargetValue, reflect.ValueOf(structTargetType{}))
    if err == nil || err.Error() != valuesNotMatchError {
        t.FailNow()
    }
}
