package main

import (
	"advent_2020/input"
	"fmt"
	"regexp"
	"strings"
)

var prefix = "day19/input/"
var lines []string
var rules map[string]Rule

func main() {
	lines = input.ReadLines(prefix + "rules.txt")
	rules := createRules(lines)
	values := input.ReadLines(prefix + "fields.txt")
	linesThatMatchRuleZero := 0
	for i := range values {
		match := rules["0"].match(values[i])
		if match {
			fmt.Println("Match found:", values[i])
			linesThatMatchRuleZero++
		}
	}
	fmt.Println(linesThatMatchRuleZero)
}

type Rule struct {
	RuleHelper
}

type RuleHelper interface {
	matchHelper(input string) (bool, string)
}

type Ands struct {
	RulesAndedTogether []Rule
}

type Ors struct {
	RulesOrdTogether []Rule
}

type Direct struct {
	value string
}

//match is used
func (rule Rule) match(input string) bool {
	success, remaining := rule.matchHelper(input)
	return success && remaining == ""
}

func (rule Ors) matchHelper(input string) (bool, string) {
	for i := range rule.RulesOrdTogether {
		match, remaining := rule.RulesOrdTogether[i].matchHelper(input)
		if match {
			return true, remaining
		}
	}
	return false, ""
}

func (rule Ands) matchHelper(input string) (bool, string) {
	allMatches := true
	remaining := input
	for i := range rule.RulesAndedTogether {
		var thisMatch bool
		thisMatch, remaining = rule.RulesAndedTogether[i].matchHelper(remaining)
		allMatches = allMatches && thisMatch
		if !allMatches {
			return false, ""
		}
	}
	return allMatches, remaining
}

func (rule Direct) matchHelper(input string) (bool, string) {
	return string(input[0]) == rule.value, input[1:]
}

func createRules(lines []string) map[string]Rule {
	rules = make(map[string]Rule)
	for i := range lines {
		createRole(lines[i])
	}
	return rules
}

var Empty Rule

func createRole(line string) Rule {
	var maskValue = regexp.MustCompile(`^(\d+): (.*)$`)
	captureGroups := maskValue.FindStringSubmatch(line)
	key := captureGroups[1]
	if rules[key] != Empty {
		return rules[key]
	}
	value := createRuleFromSpec(captureGroups[2])
	rules[key] = value
	return rules[key]
}

func createRuleFromSpec(s string) Rule {
	orGroupsStrings := strings.Split(strings.TrimSpace(s), "|")
	orGroups := make([]Rule, len(orGroupsStrings))
	for i := range orGroups {
		orGroups[i] = makeAndGroup(orGroupsStrings[i])
	}
	if len(orGroups) == 1 {
		return orGroups[0]
	}
	return Rule{Ors{RulesOrdTogether: orGroups}}
}

func makeAndGroup(andGroupString string) Rule {
	trimedAndGroup := strings.TrimSpace(andGroupString)
	andGroupsStrings := strings.Split(trimedAndGroup, " ")
	andGroups := make([]Rule, len(andGroupsStrings))
	for i := range andGroups {
		if andGroupsStrings[i] == "\"a\"" || andGroupsStrings[i] == "\"b\"" {
			andGroups[i] = Rule{Direct{value: string(andGroupsStrings[i][1])}}
		} else {
			andGroups[i] = createRuleByName(andGroupsStrings[i])
		}
	}
	if len(andGroups) == 1 {
		return andGroups[0]
	}
	return Rule{Ands{RulesAndedTogether: andGroups}}
}

func createRuleByName(name string) Rule {
	return createRole(getRuleSpecByName(name))
}

func getRuleSpecByName(name string) string {
	for ruleIndex := range lines {
		var maskValue = regexp.MustCompile(`^(\d+):.*$`)
		captureGroups := maskValue.FindStringSubmatch(lines[ruleIndex])
		if captureGroups[1] == name {
			return lines[ruleIndex]
		}
	}
	panic(name)
}
