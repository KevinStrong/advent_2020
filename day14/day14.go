package day14

import (
	"math"
	"math/bits"
	"regexp"
	"strconv"
	"strings"
)

type Mask struct {
	maskXs       []Wildcard
	maskOnes     uint64
	instructions []Instruction
}

type Instruction struct {
	value    int
	location uint64
}

type Wildcard struct {
	setOnes  uint64
	setZeros uint64
}

func Solve(input []string) int {
	return solvePartTwo(input)
}

// See previous commit for solvePartOne
func solvePartTwo(input []string) int {
	allMasks := createInstructions(input)
	results := processInstructions(allMasks)
	return sumValues(results)
}

func sumValues(results map[uint64]int) int {
	var total = 0
	for memLocation := range results {
		total += results[memLocation]
	}
	return total
}

func processInstructions(masks []Mask) map[uint64]int {
	results := make(map[uint64]int)
	for _, mask := range masks {
		for _, instruction := range mask.instructions {
			for _, wildcardMask := range mask.maskXs {
				location := (instruction.location | mask.maskOnes | wildcardMask.setOnes) & wildcardMask.setZeros
				results[location] = instruction.value
			}
		}
	}
	return results
}

func createInstructions(input []string) []Mask {
	var masks = make([]Mask, 0)
	var currentMask = Mask{
		maskXs:       make([]Wildcard, 0),
		maskOnes:     0,
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
	location, _ := strconv.ParseUint(captureGroups[1], 10, 64)
	value, _ := strconv.Atoi(captureGroups[2])
	return Instruction{
		value:    value,
		location: location,
	}
}

func createMask(line string) Mask {
	var maskValue = regexp.MustCompile(`^mask = (.*)$`)
	captureGroups := maskValue.FindStringSubmatch(line)
	replaceXWithZeros := strings.ReplaceAll(captureGroups[1], "X", "0")

	replaceOnesWithZeros := strings.ReplaceAll(captureGroups[1], "1", "0")
	thenConvertXsToOnes := strings.ReplaceAll(replaceOnesWithZeros, "X", "1")

	maskXs, _ := strconv.ParseUint(thenConvertXsToOnes, 2, 64)
	maskOnes, _ := strconv.ParseUint(replaceXWithZeros, 2, 64)
	return Mask{
		maskXs:       buildWildcards(maskXs),
		maskOnes:     maskOnes,
		instructions: make([]Instruction, 0),
	}
}

func buildWildcards(xs uint64) []Wildcard {
	var allInstructionLocations = make([]Wildcard, 0)
	allInstructionLocations = append(allInstructionLocations, Wildcard{
		setOnes:  0,
		setZeros: math.MaxUint64,
	})
	for i := 0; i < 64; i++ {
		// If the first bit is a 1 then it is a wildcard
		if 1&xs == 1 {
			var tempInstructions = make([]Wildcard, 0)
			for _, loc := range allInstructionLocations {
				tempInstructions = append(tempInstructions, Wildcard{
					setOnes:  loc.setOnes | 1,
					setZeros: loc.setZeros,
				})
				tempInstructions = append(tempInstructions, Wildcard{
					setOnes:  loc.setOnes,
					setZeros: loc.setZeros & (math.MaxUint64 - 1), // All 1s except first bit
				})
			}
			allInstructionLocations = tempInstructions
		}
		xs = bits.RotateLeft64(xs, 1)
		for i := range allInstructionLocations {
			allInstructionLocations[i].setOnes = bits.RotateLeft64(allInstructionLocations[i].setOnes, 1)
			allInstructionLocations[i].setZeros = bits.RotateLeft64(allInstructionLocations[i].setZeros, 1)
		}
	}
	return allInstructionLocations
}
