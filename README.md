## Brainfuck Interpreter

A [Brainfuck](https://nickdesaulniers.github.io/blog/2015/05/25/interpreter-compiler-jit/) interpreter written in `Go`

[![Build Status](https://github.com/clivern/bf/workflows/Build/badge.svg)](https://github.com/clivern/bf/actions/workflows/build.yml)
[![Docs](https://godoc.org/github.com/clivern/bf?status.svg)](https://godoc.org/github.com/clivern/bf)
[![Go Report Card](https://goreportcard.com/badge/github.com/clivern/bf)](https://goreportcard.com/report/github.com/clivern/bf)


### Installation

To install `bf` package, you need to install Go (version `1.13+` is required) and setup your Go project first.

```bash
$ mkdir example
$ cd example
$ go mod init example.com
```

- Then you can use the below `Go` command to install the latest version of `bf` package.

```bash
$ go get -u github.com/clivern/bf
```

Or the following for a specific version `vx.x.x`

```bash
go get -u github.com/clivern/bf@vx.x.x
```

- Import the package in your code.

```golang
import "github.com/clivern/bf"
```


### Quick start

Add the following code in `example.go` file

```bash
$ cat example.go
```

```golang
package main

import (
    "bytes"
    "fmt"
    "os"
    "strings"

    "github.com/clivern/bf"
)

func main() {
    buf := new(bytes.Buffer)

    interpreter := bf.NewInterpreter(
        strings.NewReader("-[------->+<]>-.-[->+++++<]>++.+++++++..+++.[--->+<]>-----.---[->+++<]>.-[--->+<]>---.+++.------.--------."),
        os.Stdin,
        buf,
    )

    err := interpreter.IsValid()

    if err != nil {
        panic(fmt.Sprintf("Invalid code: %s", err.Error()))
    }

    err = interpreter.Execute()

    if err != nil {
        panic(fmt.Sprintf("Invalid code: %s", err.Error()))
    }

    fmt.Println(buf.String()) // Hello World
}
```

Run `example.go`

```bash
$ go run example.go
```

It will return `Hello World`
