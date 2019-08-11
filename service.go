package marker

import (
    "errors"
    "fmt"
    "github.com/olehan/marker/output"
    "github.com/olehan/marker/utils"
    "github.com/olehan/marker/validation"
    "log"
    "math/rand"
    "os"
    "os/exec"
    "path"
    "plugin"
    "reflect"
    "time"
)

const (
    tempFolder = "/tmp"
    imageNameLen = 10
)

var (
    letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func Mark(filePath string, marker *Marker) (output.Output, error) {
    p, clean, err := newPluginByPath(filePath)
    if err != nil {
        return nil, err
    }
    defer clean()

    return outputFromPlugin(p, marker.Name, marker.Value)
}

func MarkMany(filePath string, markers ...*Marker) (output.Outputs, error) {
    p, clean, err := newPluginByPath(filePath)
    if err != nil {
        return nil, err
    }
    defer clean()

    outputMap := output.NewOutputs()
    for _, marker := range markers {
        o, err := outputFromPlugin(p, marker.Name, marker.Value)
        if err != nil {
            return nil, err
        }

        outputMap[marker.Name] = o
    }
    return outputMap, nil
}

func newPluginByPath(filePath string) (*plugin.Plugin, func(), error) {
    f, err := os.Stat(filePath)
    if err != nil {
        return nil, nil, errors.New(fmt.Sprintln("can't access the file:", err))
    }

    if f.IsDir() {
        return nil, nil, errors.New("path navigates to directory instead of the file")
    }

    outputPath := newOutputPath()

    err = buildImage(outputPath, filePath)
    if err != nil {
        return nil, nil, err
    }

    p, err := plugin.Open(outputPath)
    if err != nil {
        return nil, nil, err
    }

    return p, newCleanerFunc(outputPath), nil
}

func newCleanerFunc(imageFile string) func() {
    return func() {
        err := os.Remove(imageFile)
        if err != nil {
            log.Println("couldn't remove temp image file:", err)
        }
    }
}

func newOutputPath() string {
    return path.Join(tempFolder, utils.RandomStringRunes(imageNameLen))
}

func outputFromPlugin(plugin *plugin.Plugin, name string, value reflect.Value) (output.Output, error) {
    symbol, err := plugin.Lookup(name)
    if err != nil {
        return nil, err
    }

    err = validation.ValidateTwoValues(reflect.ValueOf(symbol).Elem(), value)
    if err != nil {
        return nil, errors.New(fmt.Sprintf("error with symbol '%s': %s", name, err.Error()))
    }

    return symbol, nil
}

func buildImage(outputPath, filePath string) error {
    cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", outputPath, filePath)
    err := cmd.Run()
    if err != nil {
        return err
    }
    return nil
}
