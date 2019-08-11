package marker

import (
    "reflect"
)

type (
    Marker struct {
        Name  string
        Value reflect.Value
    }
)

var (
    BoolValue = reflect.ValueOf(false)
    IntValue = reflect.ValueOf(0)
    Int8Value = reflect.ValueOf(int8(0))
    Int16Value = reflect.ValueOf(int16(0))
    Int32Value = reflect.ValueOf(int32(0))
    Int64Value = reflect.ValueOf(int64(0))
    UintValue = reflect.ValueOf(uint(0))
    Uint8Value = reflect.ValueOf(uint8(0))
    Uint16Value = reflect.ValueOf(uint16(0))
    Uint32Value = reflect.ValueOf(uint32(0))
    Uint64Value = reflect.ValueOf(uint64(0))
    UintptrValue = reflect.ValueOf(uintptr(0))
    Float32Value = reflect.ValueOf(float32(0.0))
    Float64Value = reflect.ValueOf(0.0)
    StringValue = reflect.ValueOf("")
    Complex64Value = reflect.ValueOf(complex64(0.0))
    Complex128Value = reflect.ValueOf(complex128(0.0))
)

func NewMarker(name string, value interface{}) *Marker {
    return NewValueMarker(name, reflect.ValueOf(value))
}

func NewValueMarker(name string, value reflect.Value) *Marker {
    return &Marker{name, value}
}
