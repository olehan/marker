package path

import (
    "strings"
    "testing"
)

func TestRemoveLastNode(t *testing.T) {
    testFile := "main.go"
    expectedValue := "/Users/olehan/path/to/project"
    value := strings.Join([]string{expectedValue, testFile}, "/")

    output := RemoveLastNode(value)
    if output != expectedValue {
        t.FailNow()
    }
}

func TestRelativeToCaller(t *testing.T) {
    t.Log(RelativeToCaller("../README.md"))
}

func TestRelativeToExecutable(t *testing.T) {
    t.Log(RelativeToExecutable("../README.md"))
}

func TestRelativeToWorkingDirectory(t *testing.T) {
    t.Log(RelativeToWorkingDirectory("../README.md"))
}
