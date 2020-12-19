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
		fmt.Printf("%d, %s", i, realTicketOrder.fields[i].name)
	}
	departureMultiple := getDepartureMultiple(realTicketOrder, yourTicket)
	fmt.Printf("Departure Multiple: %d\n", departureMultiple)
}

func getDepartureMultiple(order TicketDefinition, ticket Ticket) int {
	multiple := 1
	for i, field := range order.fields {
		if strings.HasPrefix(field.name, "Departure") {
			multiple *= ticket.values[i]
		}
	}
	return multiple
}

func findFieldOrder(tickets []Ticket, restrictions []FieldDef) TicketDefinition {
	potentialFieldOrders := findValidFieldOrders(make([]FieldDef, 0), restrictions, tickets)
	if len(potentialFieldOrders) != 1 {
		fmt.Printf("You have %d valid orders\n", len(potentialFieldOrders))
		// panic("")
	}
	return potentialFieldOrders[0]
}

func validateTicketDefinition(ticket Ticket, potentials []TicketDefinition) []TicketDefinition {
	validTicketDefinitions := make([]TicketDefinition, 0)
	for i := range potentials {
		if isTicketDefinitionValid(ticket, potentials[i]) {
			validTicketDefinitions = append(validTicketDefinitions, potentials[i])
		}
	}
	return validTicketDefinitions
}

func isTicketDefinitionValid(ticket Ticket, definition TicketDefinition) bool {
	ticket.isValid = true
	for i, value := range ticket.values {
		if !validate(value, definition.fields[i]) {
			ticket.isValid = false
		}
	}
	return ticket.isValid
}

func validate(value int, def FieldDef) bool {
	for _, rule := range def.rules {
		if value <= rule.upper && value >= rule.lower {
			return true
		}
	}
	return false
}

func findValidFieldOrders(appliedRestrictions []FieldDef, restrictionsToApply []FieldDef, tickets []Ticket) []TicketDefinition {
	if len(appliedRestrictions) < 7 {
		fmt.Printf("Depth #: %d\n", len(appliedRestrictions))
	}
	validTicketDefinitions := make([]TicketDefinition, 0)
	if !areAllTicketsPossible(restrictionsToApply, tickets) {
		return validTicketDefinitions
	}
	if len(tickets[0].values) == 0 && len(restrictionsToApply) == 0 {
		fmt.Printf("Found a valid solution!\n")
		definitions := append(validTicketDefinitions, TicketDefinition{fields: appliedRestrictions})
		return definitions
	}
	for restrictionIndex, restriction := range restrictionsToApply {
		if allFirstFieldsMeetRestriction(restriction, tickets) {
			validFieldOrders := findValidFieldOrders(
				append(appliedRestrictions, restriction),
				removeRestriction(restrictionsToApply, restrictionIndex),
				removeFirstValueFromAllTickets(tickets))
			validTicketDefinitions = append(validTicketDefinitions, validFieldOrders...)
		}
	}
	return validTicketDefinitions
}

func areAllTicketsPossible(restrictions []FieldDef, tickets []Ticket) bool {
	for _, ticket := range tickets {
		for _, value := range ticket.values {
			if !isValuePossible(value, restrictions) {
				return false
			}
		}
	}
	return true
}

func allFirstFieldsMeetRestriction(restriction FieldDef, tickets []Ticket) bool {
	for _, ticket := range tickets {
		if !meetsRestriction(restriction, ticket.values[0]) {
			return false
		}
	}
	return true
}

func removeFirstValueFromAllTickets(tickets []Ticket) []Ticket {
	updatedTickets := make([]Ticket, len(tickets))
	for index, ticket := range tickets {
		updatedTickets[index] = removeFirstValue(ticket)
	}
	return updatedTickets
}

func removeFirstValue(ticket Ticket) Ticket {
	updatedValues := make([]int, len(ticket.values)-1)
	copy(updatedValues, ticket.values[1:])
	return Ticket{values: updatedValues}
}

func removeRestriction(apply []FieldDef, i int) []FieldDef {
	updatedValues := make([]FieldDef, len(apply))
	copy(updatedValues, apply)
	updatedValues = append(updatedValues[:i], updatedValues[i+1:]...)
	return updatedValues
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
