package main

import (
	"advent_2020/input"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	lower int
	upper int
}

type FieldDef struct {
	rules []Rule
	name  string
}

type TicketDefinition struct {
	fields []FieldDef
}

type Ticket struct {
	values  []int
	isValid bool
}

func main() {
	prefix := "day16/input/"
	yourTicket := createYourTicket(input.ReadLines(prefix + "your_ticket.txt"))
	otherTickets := createOtherTickets(input.ReadLines(prefix + "other_tickets.txt"))
	restrictions := createRestrictions(input.ReadLines(prefix + "restrictions.txt"))
	otherValidTickets := possibleTicketsOnly(otherTickets, restrictions)
	allTickets := append(otherValidTickets, yourTicket)
	realTicketOrder := findFieldOrder(allTickets, restrictions.fields)
	fmt.Println("Done")
	for i := range realTicketOrder.fields {
		fmt.Printf("%d, %s\n", i, realTicketOrder.fields[i].name)
	}
	departureMultiple := getDepartureMultiple(realTicketOrder, yourTicket)
	fmt.Printf("Departure Multiple: %d\n", departureMultiple)
}

func getDepartureMultiple(order TicketDefinition, ticket Ticket) int {
	multiple := 1
	for i, field := range order.fields {
		if strings.HasPrefix(field.name, "departure") {
			multiple *= ticket.values[i]
		}
	}
	return multiple
}

func findFieldOrder(tickets []Ticket, restrictions []FieldDef) TicketDefinition {
	rowsToPossibleRestrictions := buildColumnToPossibleRestrictions(tickets, restrictions)
	solution, success := solveTwoDimensionalArray(rowsToPossibleRestrictions)
	if !success {
		panic("We failed boss!")
	}
	return convertToTicketDefinition(solution, restrictions)
}

func convertToTicketDefinition(solution [][]int, allRestrictions []FieldDef) TicketDefinition {
	orderedRestrictions := make([]FieldDef, len(allRestrictions))
	for i := range solution {
		if len(solution[i]) != 1 {
			fmt.Println("Solution is broken", len(solution[i]))
		}
		restrictionIndexForThisColumn := solution[i][0]
		orderedRestrictions[i] = allRestrictions[restrictionIndexForThisColumn]
	}
	return TicketDefinition{fields: orderedRestrictions}
}

func solveTwoDimensionalArray(rowsToPossibleRestrictions [][]int) ([][]int, bool) {
	if isSolved(rowsToPossibleRestrictions) {
		return rowsToPossibleRestrictions, true
	}
	if !isSolutionPossible(rowsToPossibleRestrictions) {
		return rowsToPossibleRestrictions, false
	}
	rowsToPossibleRestrictions = copy2DArray(rowsToPossibleRestrictions)

	rowsToPossibleRestrictions = deduceWhatWeCan(rowsToPossibleRestrictions)

	if isSolved(rowsToPossibleRestrictions) {
		return rowsToPossibleRestrictions, true
	}

	var guessedSolution [][]int
	var success bool
	for i := range rowsToPossibleRestrictions {
		if len(rowsToPossibleRestrictions[i]) > 1 {
			for i2 := range rowsToPossibleRestrictions[i] {
				guessedSolution = guessThisIndex(rowsToPossibleRestrictions, i, i2)
				guessedSolution, success = solveTwoDimensionalArray(guessedSolution)
				if success {
					return guessedSolution, true
				}
			}
		}
	}
	return rowsToPossibleRestrictions, false
}

func guessThisIndex(restrictions [][]int, i int, i2 int) [][]int {
	guess := copy2DArray(restrictions)
	removeFromArray(guess, i, i2)
	return guess
}

// Returns a copy of the array that as been solved with a simple sudoku style solver
func deduceWhatWeCan(restrictions [][]int) [][]int {
	updatedArray, changeMade := attemptToReduce(restrictions)
	for changeMade {
		updatedArray, changeMade = attemptToReduce(updatedArray)
	}
	return updatedArray
}

func attemptToReduce(restrictions [][]int) ([][]int, bool) {
	changeMade := false
	for i := range restrictions {
		if len(restrictions[i]) == 1 {
			restrictions, changeMade = removeFromTwoDArray(restrictions, restrictions[i][0])
		}
	}
	return restrictions, changeMade
}

func removeFromTwoDArray(problem [][]int, valueToRemove int) ([][]int, bool) {
	changeMade := false
	for i := range problem {
		changeMade = removeFromArray(problem, i, valueToRemove) || changeMade
	}
	return problem, changeMade
}

func removeFromArray(updatedArray [][]int, i int, valueToRemove int) bool {
	for i2 := range updatedArray[i] {
		if updatedArray[i][i2] == valueToRemove && len(updatedArray[i]) != 1 {
			updatedArray[i] = remove(updatedArray[i], i2)
			return true
		}
	}
	return false
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func copy2DArray(restrictions [][]int) [][]int {
	copyOfArray := make([][]int, len(restrictions))
	for i := range restrictions {
		copyOfArray[i] = make([]int, len(restrictions[i]))
		copy(copyOfArray[i], restrictions[i])
	}
	return copyOfArray
}

func isSolutionPossible(restrictions [][]int) bool {
	for i := range restrictions {
		if len(restrictions[i]) < 1 {
			return false
		}
	}
	return true
}

func isSolved(restrictions [][]int) bool {
	for i := range restrictions {
		if len(restrictions[i]) != 1 {
			return false
		}
	}
	return true
}

// Index of first array matches the column number, then the second array contains all indexs of restrictions that match
func buildColumnToPossibleRestrictions(tickets []Ticket, restrictions []FieldDef) [][]int {
	numberOfColumns := len(restrictions)
	rowToRestrictions := make([][]int, numberOfColumns)
	for i := 0; i < numberOfColumns; i++ {
		rowToRestrictions[i] = buildPossibleRestrictionsForColumn(getColumn(tickets, i), restrictions)
	}
	return rowToRestrictions
}

func getColumn(tickets []Ticket, i int) []int {
	column := make([]int, len(tickets))
	for i2 := range tickets {
		column[i2] = tickets[i2].values[i]
	}
	return column
}

func buildPossibleRestrictionsForColumn(column []int, restrictions []FieldDef) []int {
	possibleRestrictions := make([]int, 0)
	for i, restriction := range restrictions {
		if doAllColumnValuesMeetRestriction(column, restriction) {
			possibleRestrictions = append(possibleRestrictions, i)
		}
	}
	return possibleRestrictions
}

func doAllColumnValuesMeetRestriction(column []int, restriction FieldDef) bool {
	for _, columnValue := range column {
		if !meetsRestriction(restriction, columnValue) {
			return false
		}
	}
	return true
}

func meetsRestriction(restriction FieldDef, i int) bool {
	for _, rule := range restriction.rules {
		if i >= rule.lower && i <= rule.upper {
			return true
		}
	}
	return false
}

func possibleTicketsOnly(otherTickets []Ticket, restrictions TicketDefinition) []Ticket {
	validTickets := make([]Ticket, 0)
	for _, ticket := range otherTickets {
		for _, value := range ticket.values {
			if !isValuePossible(value, restrictions.fields) {
				ticket.isValid = false
			}
		}
		if ticket.isValid {
			validTickets = append(validTickets, ticket)
		}
	}
	return validTickets
}

func isValuePossible(value int, restrictions []FieldDef) bool {
	for _, field := range restrictions {
		for _, rule := range field.rules {
			if rule.lower <= value && rule.upper >= value {
				return true
			}
		}
	}
	return false
}

func createRestrictions(lines []string) TicketDefinition {
	fields := make([]FieldDef, 0)
	for _, line := range lines {
		fields = append(fields, createFieldDef(line))
	}
	return TicketDefinition{fields: fields}
}

func createFieldDef(line string) FieldDef {
	var bothRanges = regexp.MustCompile(`^([\w ]+): (\d+)-(\d+) or (\d+)-(\d+)$`)
	captureGroups := bothRanges.FindStringSubmatch(line)
	firstLower, _ := strconv.Atoi(captureGroups[2])
	firstUpper, _ := strconv.Atoi(captureGroups[3])
	firstRange := Rule{
		lower: firstLower,
		upper: firstUpper,
	}
	secondLower, _ := strconv.Atoi(captureGroups[4])
	secondUpper, _ := strconv.Atoi(captureGroups[5])
	secondRange := Rule{
		lower: secondLower,
		upper: secondUpper,
	}
	return FieldDef{
		rules: []Rule{firstRange, secondRange},
		name:  captureGroups[1],
	}
}

func createOtherTickets(lines []string) []Ticket {
	tickets := make([]Ticket, 0)
	for _, line := range lines {
		tickets = append(tickets, createTicket(line))
	}
	return tickets
}

func createYourTicket(lines []string) Ticket {
	return createTicket(lines[0])
}

func createTicket(line string) Ticket {
	values := make([]int, 0)
	for _, numberString := range strings.Split(line, ",") {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			panic(err)
		}
		values = append(values, number)
	}
	return Ticket{
		values:  values,
		isValid: true,
	}
}
