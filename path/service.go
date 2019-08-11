package path

import (
    "os"
    "path"
    "runtime"
    "strings"
)

func RelativeToCaller(pathToFile string) string {
    _, fileName, _, ok := runtime.Caller(1)
    if !ok {
        panic("could not get caller information")
    }
    return path.Join(RemoveFileFromPath(fileName), pathToFile)
}

func RelativeToExecutable(pathToFile string) string {
    fileName, err := os.Executable()
    if err != nil {
        panic("could not get path name of the executable")
    }
    return path.Join(fileName, pathToFile)
}

func RelativeToWorkingDirectory(pathToFile string) string {
    fileName, err := os.Getwd()
    if err != nil {
        panic("could not get working directory")
    }
    return path.Join(fileName, pathToFile)
}

func RemoveFileFromPath(path string) string {
    fileName := strings.Split(path, "/")
    return strings.Join(fileName[0:len(fileName) - 1], "/")
}
