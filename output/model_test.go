package output

import (
    "plugin"
    "testing"
)

func TestNewOutputs(t *testing.T) {
    expectedKey := "key1"
    expectedValue := plugin.Symbol(&expectedKey)

    outputs := NewOutputs()
    outputs[expectedKey] = expectedValue

    if !outputs.Has(expectedKey) {
        t.FailNow()
    }

    if outputs.Get(expectedKey) != expectedValue {
        t.FailNow()
    }
}
