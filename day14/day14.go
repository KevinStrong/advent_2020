package day14

import (
	"regexp"
	"strconv"
	"strings"
)

type Mask struct {
	maskOnes     uint64
	maskZeros    uint64
	instructions []Instruction
}

type Instruction struct {
	value    uint64
	location int
}

func Solve(input []string) int {
	return solvePartOne(input)
}

func solvePartOne(input []string) int {
	allMasks := createInstructions(input)
	results := processInstructions(allMasks)
	return int(sumValues(results))
}

func sumValues(results map[int]uint64) uint64 {
	var total uint64 = 0
	for memLocation := range results {
		total += results[memLocation]
	}
	return total
}

func processInstructions(masks []Mask) map[int]uint64 {
	results := make(map[int]uint64)
	for _, mask := range masks {
		for _, instruction := range mask.instructions {
			results[instruction.location] = (mask.maskOnes | instruction.value) & mask.maskZeros
		}
	}
	return results
}

func createInstructions(input []string) []Mask {
	var masks = make([]Mask, 0)
	var currentMask = Mask{
		maskOnes:     0,
		maskZeros:    0,
		instructions: make([]Instruction, 0),
	}
	for _, inputLine := range input {
		if strings.HasPrefix(inputLine, "mask") {
			if len(currentMask.instructions) > 0 {
				masks = append(masks, currentMask)
			}
			currentMask = createMask(inputLine)
		} else {
			currentMask.instructions =
				append(currentMask.instructions, buildMemoryInstructionInstruction(inputLine))
		}
	}
	masks = append(masks, currentMask)
	return masks
}

func buildMemoryInstructionInstruction(line string) Instruction {
	var maskValue = regexp.MustCompile(`^mem\[(\d+)] = (\d+)$`)
	captureGroups := maskValue.FindStringSubmatch(line)
	location, _ := strconv.Atoi(captureGroups[1])
	value, _ := strconv.ParseUint(captureGroups[2], 10, 36)
	return Instruction{
		value:    value,
		location: location,
	}
}

func createMask(line string) Mask {
	var maskValue = regexp.MustCompile(`^mask = (.*)$`)
	captureGroups := maskValue.FindStringSubmatch(line)
	replaceXWithZeros := strings.ReplaceAll(captureGroups[1], "X", "0")
	replaceXWithOnes := strings.ReplaceAll(captureGroups[1], "X", "1")
	maskOnes, _ := strconv.ParseUint(replaceXWithZeros, 2, 64)
	maskZeros, _ := strconv.ParseUint(replaceXWithOnes, 2, 64)
	return Mask{
		maskOnes:     maskOnes,
		maskZeros:    maskZeros,
		instructions: make([]Instruction, 0),
	}
}
