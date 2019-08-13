package marker

import (
    "github.com/olehan/marker/path"
    "testing"
)

func TestMark(t *testing.T) {
    m, err := Mark(path.RelativeToCaller("./examples/many_symbols/target.go"), NewValueMarker("Version", StringValue))
    if err != nil {
        t.Log(err)
        t.FailNow()
    }

    t.Log(*m.(*string))
}

func TestMarkMany(t *testing.T) {
    outputMap, err := MarkMany(
        path.RelativeToCaller("./examples/many_symbols/target.go"),
        NewValueMarker("Version", StringValue),
        NewValueMarker("Depth", IntValue),
        NewValueMarker("Date", Float64Value),
        NewValueMarker("Author", StringValue),
        NewValueMarker("Repo", StringValue),
        NewValueMarker("Name", StringValue),
        NewValueMarker("Private", BoolValue),
    )

    if err != nil {
        t.Log(err)
        t.FailNow()
    }

    for key, value := range outputMap {
        t.Log(key, value)
    }
}
