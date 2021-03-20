// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

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
