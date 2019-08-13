<h1 align="center">✏️ Marker</h1>
<p align="center">Mark your golang plugin file to get the best of it</p>

<p align="center">
    <a href="https://travis-ci.org/olehan/marker">
        <img alt="Marker Travis Status" src="https://img.shields.io/travis/olehan/marker?logo=travis">
    </a>
    <a href="https://codecov.io/gh/olehan/marker">
        <img alt="Marker Code Coverage" src="https://codecov.io/gh/olehan/marker/branch/master/graph/badge.svg">
    </a>
    <a href="https://github.com/olehan/marker/blob/master/LICENSE">
        <img alt="Marker License" src="https://img.shields.io/github/license/olehan/marker.svg">
    </a>
    <a href="https://github.com/olehan/marker/releases">
        <img alt="Marker Last Release" src="https://img.shields.io/github/tag/olehan/marker.svg?label=release">
    </a>
</p>

----

<p align="center">
    <strong><a href="#usage">Usage</a></strong>
    |
    <strong><a href="#installation">Installation</a></strong>
    |
    <strong><a href="#license">License</a></strong>
</p>

----

I know that this (✏️) is not a marker, but I do not have anything better ＼(￣▽￣)／

<h2 id="usage">Usage</h2>

*Actually I'd suggest you to go into the [examples directory](examples) to see more usage
variations.*

***Project:***
```
project
├── config.go
└── app/
    └── main.go
```

***config.go***
```go
package main // Notice that the package has to be main to build a plugin.

// Variables and functions also have to be public to import them using marker. 
var ApiUrl = "https://api.github.com"
```

***app/main.go***
```go
package main
 
import (
    "fmt"
    "github.com/olehan/marker"
    "github.com/olehan/marker/path"
)

func main() {
    // Here we're marking our file as a plugin and telling the marker what we need
    // and what it should be.
    output, err := marker.Mark(
        path.RelativeToCaller("../config.go"),
        // We're specifying the variable's name and its type so marker could
        // validate the value that was imported from the plugin.
        marker.NewMarker("ApiUrl", marker.StringValue),
    )

    if err != nil {
        panic(err)
    }

    // Symbol's always pointer.
    apiUrl := output.(*string)   
    fmt.Println("apiUrl:", *apiUrl)
}
```


<h2 id="installation">Installation</h2>

```sh
$ go get github.com/olehan/marker
```


<h2 id="license">License</h2>

[MIT](LICENSE)
