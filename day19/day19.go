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
	fmt.Println("Rule created")
	values := input.ReadLines(prefix + "fields.txt")
	fmt.Println("Matching Values")
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
	matchHelper(input string) (bool, []string)
	wireUpPlaceholder(rules map[string]Rule, depth int) Rule
}

type Ands struct {
	RulesAndedTogether []Rule
}

type Ors struct {
	name             string
	RulesOrdTogether []Rule
}

type Direct struct {
	value string
}

type Placeholder struct {
	ruleName string
}

func (rule Rule) match(input string) bool {
	success, remaining := rule.matchHelper(input)
	atLeastOneMatchesFullInput := false
	for i := range remaining {
		atLeastOneMatchesFullInput = atLeastOneMatchesFullInput || remaining[i] == ""
	}
	return success && atLeastOneMatchesFullInput
}

func (rule Placeholder) matchHelper(_ string) (bool, []string) {
	panic("Calling matchHelper on a placeholder")
}

func (rule Ors) matchHelper(input string) (bool, []string) {
	if len(input) == 0 {
		return false, []string{""}
	}
	possibleRemainings := make([]string, 0)
	for i := range rule.RulesOrdTogether {
		isMatch, remaining := rule.RulesOrdTogether[i].matchHelper(input)
		if isMatch {
			possibleRemainings = append(possibleRemainings, remaining...)
		}
	}
	return len(possibleRemainings) > 0, possibleRemainings
}

func (rule Ands) matchHelper(input string) (bool, []string) {
	if len(input) == 0 {
		return false, []string{""}
	}
	nextRemaining := make([]string, 1)
	nextRemaining[0] = input
	var remaining []string
	for andIndex := range rule.RulesAndedTogether {
		remaining = nextRemaining
		nextRemaining = make([]string, 0)
		for potentialRemainingIndex := range remaining {
			thisMatch, thisRemaining := rule.RulesAndedTogether[andIndex].matchHelper(remaining[potentialRemainingIndex])
			if thisMatch {
				nextRemaining = createIntersection(nextRemaining, thisRemaining)
			}
		}
		if len(nextRemaining) == 0 {
			return false, make([]string, 0)
		}
	}
	return true, nextRemaining
}

func createIntersection(first []string, second []string) []string {
	encountered := map[string]bool{}

	for _, c := range first {
		encountered[c] = true
	}
	for _, c := range second {
		encountered[c] = true
	}

	var result = make([]string, len(encountered))
	i := 0
	for key := range encountered {
		result[i] = key
		i++
	}
	return result
}

func (rule Direct) matchHelper(input string) (bool, []string) {
	remaining := []string{input[1:]}
	return string(input[0]) == rule.value, remaining
}

func (rule Placeholder) wireUpPlaceholder(rules map[string]Rule, _ int) Rule {
	return rules[rule.ruleName]
}

func (rule Ors) wireUpPlaceholder(rules map[string]Rule, depth int) Rule {
	depth++
	if depth > 100 {
		return Rule{rule}
	}
	for i := range rule.RulesOrdTogether {
		rule.RulesOrdTogether[i] = rule.RulesOrdTogether[i].wireUpPlaceholder(rules, depth)
	}
	return Rule{rule}
}

func (rule Ands) wireUpPlaceholder(rules map[string]Rule, depth int) Rule {
	depth++
	if depth > 100 {
		return Rule{rule}
	}
	for i := range rule.RulesAndedTogether {
		rule.RulesAndedTogether[i] = rule.RulesAndedTogether[i].wireUpPlaceholder(rules, depth)
	}
	return Rule{rule}
}

func (rule Direct) wireUpPlaceholder(_ map[string]Rule, _ int) Rule {
	return Rule{rule}
}

func createRules(lines []string) map[string]Rule {
	rules = make(map[string]Rule)
	for i := range lines {
		createRole(lines[i])
	}
	fmt.Println("Wire up placeholders")
	wireUpRecursivePlaceholderValues(rules)
	return rules
}

func wireUpRecursivePlaceholderValues(r map[string]Rule) {
	for ruleKey, ruleValue := range r {
		fmt.Println("Wiring up: ", ruleKey)
		r[ruleKey] = ruleValue.wireUpPlaceholder(r, 0)
	}
}

var Empty Rule

func createRole(line string) Rule {
	var maskValue = regexp.MustCompile(`^(\d+): (.*)$`)
	captureGroups := maskValue.FindStringSubmatch(line)
	key := captureGroups[1]
	if rules[key] != Empty {
		return rules[key]
	}
	rules[key] = Rule{Placeholder{ruleName: key}}
	value := createRuleFromSpec(captureGroups[1], captureGroups[2])
	rules[key] = value
	return rules[key]
}

func createRuleFromSpec(name string, spec string) Rule {
	orGroupsStrings := strings.Split(strings.TrimSpace(spec), "|")
	orGroups := make([]Rule, len(orGroupsStrings))
	for i := range orGroups {
		orGroups[i] = makeAndGroup(orGroupsStrings[i])
	}
	if len(orGroups) == 1 {
		return orGroups[0]
	}
	return Rule{Ors{name: name, RulesOrdTogether: orGroups}}
}

func makeAndGroup(andGroupString string) Rule {
	trimmedAndGroup := strings.TrimSpace(andGroupString)
	andGroupsStrings := strings.Split(trimmedAndGroup, " ")
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
