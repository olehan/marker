package validation

import (
    "errors"
    "fmt"
    "reflect"
)

const (
    valuesNotMatchError = "the marker value and the plugin value does not match"
)

func ValidateTwoValues(target reflect.Value, example reflect.Value) error {
    if !target.IsValid() || !example.IsValid() {
        return errors.New("invalid reflect value")
    }

    v1Kind := target.Kind()
    v2Kind := example.Kind()

    isV1Ptr := v1Kind == reflect.Ptr
    isV2Ptr := v2Kind == reflect.Ptr

    if isV1Ptr != isV2Ptr {
        return errors.New(valuesNotMatchError)
    }
    if isV1Ptr {
        target = target.Elem()
    }
    if isV2Ptr {
        example = example.Elem()
    }

    switch v1Kind {
    case reflect.Struct:
        for i := 0; i < target.NumField(); i++ {
            field := target.Field(i)
            fieldKind := field.Kind()
            exampleValue := example.Field(i)

            if fieldKind != exampleValue.Kind() {
                return errors.New(fmt.Sprintf(
                    "the struct field '%v' doesn't match the markers type",
                    field.String()),
                )
            }

            if fieldKind == reflect.Struct {
                err := ValidateTwoValues(field, exampleValue)
                if err != nil {
                    return err
                }
            }
        }
    default:
        if target.Kind() != example.Kind() {
            return errors.New(valuesNotMatchError)
        }
    }

    return nil
}
