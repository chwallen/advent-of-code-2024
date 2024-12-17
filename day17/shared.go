package day17

import (
	"strings"

	"aoc/common"
)

type instruction struct {
	input1 *int
	input2 *int
	output *int
	op     func(*int, *int, *int)
}

type program struct {
	instructions []instruction
	a            int
	b            int
	c            int
	pc           int
	outputIndex  int
	output       []int
}

func parseIntoProgram(lines []string) *program {
	p := &program{}

	p.a = common.Atoi(strings.TrimPrefix(lines[0], "Register A: "))
	p.b = common.Atoi(strings.TrimPrefix(lines[1], "Register B: "))
	p.c = common.Atoi(strings.TrimPrefix(lines[2], "Register C: "))

	tokens := strings.Split(strings.TrimPrefix(lines[4], "Program: "), ",")
	p.instructions = make([]instruction, 0, len(tokens)/2)
	p.output = make([]int, 0, len(tokens))
	p.outputIndex = 0

	for i := 0; i < len(tokens); i += 2 {
		opCode := tokens[i]
		operand := common.Atoi(tokens[i+1])

		var op func(*int, *int, *int)
		var input1, input2, output *int

		switch opCode {
		case "0":
			op = divideAndStore
			input1 = &p.a
			input2 = getComboOperandValue(&operand, &p.a, &p.b, &p.c)
			output = &p.a
		case "1":
			op = xorAndStore
			input1 = &p.b
			input2 = &operand
			output = &p.b
		case "2":
			op = mod8AndStore
			input1 = getComboOperandValue(&operand, &p.a, &p.b, &p.c)
			input2 = nil
			output = &p.b
		case "3":
			op = func(a *int, jumpTarget *int, pc *int) {
				if *a != 0 {
					// Compensate for pc being half of expected and loop adding 1
					*pc = (*jumpTarget / 2) - 1
				}
			}
			input1 = &p.a
			input2 = &operand
			output = &p.pc
		case "4":
			op = xorAndStore
			input1 = &p.b
			input2 = &p.c
			output = &p.b
		case "5":
			op = func(input *int, outIndex *int, _ *int) {
				value := *input % 8
				if len(p.output) == len(tokens) {
					p.output[*outIndex] = value
					*outIndex = (*outIndex + 1) % len(tokens)
				} else {
					p.output = append(p.output, value)
				}
			}
			input1 = getComboOperandValue(&operand, &p.a, &p.b, &p.c)
			input2 = &p.outputIndex
			output = nil
		case "6":
			op = divideAndStore
			input1 = &p.a
			input2 = getComboOperandValue(&operand, &p.a, &p.b, &p.c)
			output = &p.b
		case "7":
			op = divideAndStore
			input1 = &p.a
			input2 = getComboOperandValue(&operand, &p.a, &p.b, &p.c)
			output = &p.c
		default:
			panic("Invalid operation code " + opCode)
		}
		p.instructions = append(p.instructions, instruction{input1, input2, output, op})
	}

	return p
}

func (p *program) execute() []int {
	for p.pc = 0; p.pc < len(p.instructions); p.pc += 1 {
		inst := p.instructions[p.pc]
		inst.op(inst.input1, inst.input2, inst.output)
	}
	return p.output
}

func divideAndStore(numerator *int, factor *int, output *int) {
	*output = *numerator / common.IntPow(2, *factor)
}

func mod8AndStore(numerator *int, _ *int, output *int) {
	*output = *numerator % 8
}

func xorAndStore(left *int, right *int, output *int) {
	*output = *left ^ *right
}

func getComboOperandValue(operand *int, a *int, b *int, c *int) *int {
	switch *operand {
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	default:
		return operand
	}
}
