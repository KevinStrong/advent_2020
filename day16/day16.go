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
}

type TicketDefinition struct {
	fields []FieldDef
}

type Ticket struct {
	values []int
}

func main() {
	prefix := "day16/input/"
	_ = createYourTicket(input.ReadLines(prefix + "your_ticket.txt"))
	otherTickets := createOtherTickets(input.ReadLines(prefix + "other_tickets.txt"))
	restrictions := createRestrictions(input.ReadLines(prefix + "restrictions.txt"))
	invalidValues := 0
	for _, ticket := range otherTickets {
		for _, value := range ticket.values {
			if !isValueValid(value, restrictions) {
				fmt.Printf("Invalid Number %d\n", value)
				invalidValues += value
			}
		}
	}
	fmt.Println(invalidValues)
}

func isValueValid(value int, restrictions TicketDefinition) bool {
	for _, field := range restrictions.fields {
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
	var bothRanges = regexp.MustCompile(`^[\w ]+: (\d+)-(\d+) or (\d+)-(\d+)$`)
	captureGroups := bothRanges.FindStringSubmatch(line)
	firstLower, _ := strconv.Atoi(captureGroups[1])
	firstUpper, _ := strconv.Atoi(captureGroups[2])
	firstRange := Rule{
		lower: firstLower,
		upper: firstUpper,
	}
	secondLower, _ := strconv.Atoi(captureGroups[3])
	secondUpper, _ := strconv.Atoi(captureGroups[4])
	secondRange := Rule{
		lower: secondLower,
		upper: secondUpper,
	}
	return FieldDef{
		rules: []Rule{firstRange, secondRange},
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
		values: values,
	}
}
