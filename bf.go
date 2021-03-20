// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package bf

import (
	"errors"
	"io"
	"strings"
)

const (
	// Plus instruction constant
	Plus byte = '+'
	// Minus instruction constant
	Minus byte = '-'
	// Right instruction constant
	Right byte = '>'
	// Left instruction constant
	Left byte = '<'
	// PutChar instruction constant
	PutChar byte = '.'
	// ReadChar instruction constant
	ReadChar byte = ','
	// JumpIfZero instruction constant
	JumpIfZero byte = '['
	// JumpIfNotZero instruction constant
	JumpIfNotZero byte = ']'
)

// Interpreter brainfuck interpreter type
type Interpreter struct {
	// The instructions that need to be executed
	code io.Reader

	// Instruction list
	instructions string

	// Holds the value of the tape cell index
	instructionPointer int

	// The tape has 30000 cells, each cell hold an integer value from 0 to 255 (default is 0)
	tape [30000]int

	// Points to the tape cell
	dataPointer int

	// Used to read characters
	input io.Reader

	// Used to output characters
	output io.Writer

	buf []byte
}

// NewInterpreter creates an interpreter instance
func NewInterpreter(code io.Reader, in io.Reader, out io.Writer) *Interpreter {
	return &Interpreter{
		code:         code,
		input:        in,
		output:       out,
		instructions: "",
		buf:          make([]byte, 1),
	}
}

// IsValid validates brainfuck code
// it ensures that code doesn't start with `]` and has a matching `[` and `]`
func (i *Interpreter) IsValid() error {
	count := 0

	// Read code instruction from input stream
	i.readAllInstructions()

	for _, op := range strings.Split(i.instructions, "") {
		if op == "[" {
			count++
		} else if op == "]" {
			count--
			if count < 0 {
				return errors.New("Invalid code: ] is before [")
			}
		}
	}

	if count < 0 {
		return errors.New("Invalid code: Mismatched []")
	}

	return nil
}

// Execute executes the brainfuck code
func (i *Interpreter) Execute() error {
	var err error

	// Read code instruction from input stream
	i.readAllInstructions()

	for i.instructionPointer < len(i.instructions) {
		ins := i.instructions[i.instructionPointer]

		switch ins {
		case Plus:
			// Increment (increase by one) the byte at the data pointer.
			i.tape[i.dataPointer]++

			// Update cell value to 0 since it should hold an integer value from 0 to 255 (ASCII Chart)
			if i.tape[i.dataPointer] == 256 {
				i.tape[i.dataPointer] = 0
			}

		case Minus:
			// Decrement (decrease by one) the byte at the data pointer.
			i.tape[i.dataPointer]--

			// Update cell value to 0 since it should hold an integer value from 0 to 255 (ASCII Chart)
			if i.tape[i.dataPointer] == -1 {
				i.tape[i.dataPointer] = 255
			}

		case Right:
			// Increment the data pointer (to point to the next cell to the right).
			i.dataPointer++

		case Left:
			// Decrement the data pointer (to point to the next cell to the left).
			i.dataPointer--

		case ReadChar:
			// Accept one byte of input, storing its value in the byte at the data pointer.
			err = i.readChar()

			if err != nil {
				return err
			}

		case PutChar:
			// Output the byte at the data pointer.
			err = i.putChar()

			if err != nil {
				return err
			}

		case JumpIfZero:
			// If the byte at the data pointer is zero, jump it forward to the command after the matching ] command.
			if i.tape[i.dataPointer] == 0 {
				depth := 1

				for depth != 0 {
					i.instructionPointer++
					switch i.instructions[i.instructionPointer] {
					case JumpIfZero:
						depth++
					case JumpIfNotZero:
						depth--
					}
				}
			}

		case JumpIfNotZero:
			// If the byte at the data pointer is nonzero, jump it back to the command after the matching [ command.
			if i.tape[i.dataPointer] != 0 {
				depth := 1

				for depth != 0 {
					i.instructionPointer--
					switch i.instructions[i.instructionPointer] {
					case JumpIfNotZero:
						depth++
					case JumpIfZero:
						depth--
					}
				}
			}
		}

		i.instructionPointer++
	}

	return nil
}

// readAllInstructions reads all instruction from input stream
func (i *Interpreter) readAllInstructions() {
	// skip if instructions already exit
	if i.instructions != "" {
		return
	}

	p := make([]byte, 1)

	instructions := make([]string, 0)

	for {
		n, err := i.code.Read(p)

		if err == io.EOF {
			break
		}

		instructions = append(instructions, string(p[:n]))
	}

	i.instructions = strings.Join(instructions, "")
}

// readChar reads one byte from the input
func (i *Interpreter) readChar() error {
	n, err := i.input.Read(i.buf)

	if err != nil {
		return err
	}

	if n != 1 {
		return errors.New("wrong bytes read")
	}

	i.tape[i.dataPointer] = int(i.buf[0])

	return nil
}

// putChar writes the content of the current tape cell to the output stream
func (i *Interpreter) putChar() error {
	i.buf[0] = byte(i.tape[i.dataPointer])

	n, err := i.output.Write(i.buf)

	if err != nil {
		return err
	}

	if n != 1 {
		return errors.New("wrong bytes written")
	}

	return nil
}
