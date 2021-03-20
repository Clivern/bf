// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package bf

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/franela/goblin"
)

// TestConstants test cases
func TestConstants(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#Constants", func() {
		g.It("It should satisfy all provided test cases", func() {
			g.Assert(string(Plus)).Equal("+")
			g.Assert(string(Minus)).Equal("-")
			g.Assert(string(Right)).Equal(">")
			g.Assert(string(Left)).Equal("<")
			g.Assert(string(PutChar)).Equal(".")
			g.Assert(string(ReadChar)).Equal(",")
			g.Assert(string(JumpIfZero)).Equal("[")
			g.Assert(string(JumpIfNotZero)).Equal("]")
		})
	})
}

// TestType test cases
func TestType(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#InterpreterType", func() {
		g.It("It should satisfy all provided test cases", func() {
			g.Assert(fmt.Sprintf("%T", NewInterpreter(strings.NewReader(""), os.Stdin, os.Stdout))).Equal("*bf.Interpreter")
		})
	})
}

// TestIsValid test cases
func TestIsValid(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#IsValid", func() {
		var tests = []struct {
			input         string
			wantError     bool
			wantErrorText string
		}{
			{"+[----->+++<]>+.---.[--->+<]>++.", false, ""},
			{"+[----->+++<]>+.---.", false, ""},
			{"+[----->+++<]>+.---.++++++.", false, ""},
			{"+[----->+++<[>+.---.++++++.", true, "Invalid code: Mismatched []"},
			{"+[----->+++<]>+++.+++++++++++.", false, ""},
			{"+[----->+++]<]>+++.+++++++++++.", true, "Invalid code: Mismatched []"},
			{"+----->+++]<]>+++.+++++++++++.", true, "Invalid code: ] is before ["},
		}

		for _, tt := range tests {
			g.It(fmt.Sprintf("It should satisfy test case for brainfuck code '%s'", tt.input), func() {
				interpreter := NewInterpreter(strings.NewReader(tt.input), os.Stdin, os.Stdout)

				result := interpreter.IsValid()

				g.Assert(result != nil).Equal(tt.wantError)

				if result != nil {
					g.Assert(result.Error()).Equal(tt.wantErrorText)
				}
			})
		}
	})
}

// TestExecute test cases
func TestExecute(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#Execute", func() {
		var tests = []struct {
			input      string
			wantError  bool
			wantResult string
		}{
			{"+[----->+++<]>+.---.+++++++..+++.++++++++.--------.+++.------.--------.", false, "helloworld"},
			{"+[----->+++<]>+.---.+++++++..+++.>+[--->++<]>+.-[--->+<]>---.+++.------.--------.", false, "helloWorld"},
			{"--------[-->+++<]>.+++[->+++<]>.[--->+<]>----.+.", false, "test"},
			{"--[----->+<]>---.+++++++++.---.[->++++++<]>.+[->+++<]>.+++++++++++++.----.", false, "clivern"},
			{"+[--------->++<]>+.++[->+++<]>++.++++++++++++.+++.----.-------.", false, "sample"},
			{"+[--------->++<]>+.++[->+++<]>++.++++++++++++.+++.----.-------.[->+++<]>+.+.", false, "sample01"},
			{"------[-->+++<]>.--------.+++.------.--------.", false, "world"},
			{"--[----->+<]>-.----.--[--->+<]>---.++.------------.---[->+++<]>-.", false, "earth."},
			{"-[----->+<]>--.+.+.+.+.+.+.+.+.---------.[--->++<]>++.[----->+<]>.-[--->++<]>+.[----->+<]>.-[->++++++<]>..[-->+<]>+..++[----->+<]>+.--.[---->+++<]>.+++[->++<]>.[--->++<]>---.[->+++++<]>.+.+[--->++<]>-.[->+++++<]>++.-[->+++++<]>+.-[->++<]>..[++>---<]>+.", false, "1234567890\":';\\//><-`=12!§?||;"},
			{"+[----->+++<]>.+[--->+<]>+.+[++>---<]>..----.+[->++<]>.[-->+<]>--.[--->+<]>+.---.", false, "gy773h2gd"},
			{"--[----->+<]>--.++++++++++++++.-------------.+++++++++++.-.+[-->+<]>.------.+++[->++<]>.", false, "drepo82j"},
			{"-[----->+<]>---......--...-----------.+.----[->++<]>.-[--->++<]>--.++[-->+++<]>+.-[--->++<]>.----.---.---[->++<]>.", false, "000000...#$@(@*&#@"},
			{">++++++++[-<+++++++++>]<.>>+>-[+]++>++>+++[>[->+++<<+++>]<<]>-----.>-> Comments can be added+++..+++.>-.<<+[>[+>+]>>]<--------------.>>.+++.------.--------.>+.>+.", false, "Hello World!\n"},
		}

		for _, tt := range tests {
			g.It(fmt.Sprintf("It should satisfy test case for brainfuck code '%s'", tt.input), func() {
				buf := new(bytes.Buffer)

				interpreter := NewInterpreter(strings.NewReader(tt.input), os.Stdin, buf)
				err := interpreter.Execute()

				g.Assert(err != nil).Equal(tt.wantError)
				g.Assert(fmt.Sprintf("%s", buf.String())).Equal(tt.wantResult)
			})
		}
	})
}
