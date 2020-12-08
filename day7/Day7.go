package day7

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Rules struct {
	bag    string
	amount int
}

func Solve(ruleInput []string) int {
	ruleMap := buildRulesMap(ruleInput)
	fmt.Printf("Part One Solution: %d\n", solvePartOne(ruleMap))
	return solvePartTwo(ruleMap)
}

func solvePartTwo(ruleMap map[string][]Rules) int {
	return getAllBagsInside("shiny gold bag", ruleMap)
}

func getAllBagsInside(bag string, ruleMap map[string][]Rules) int {
	var bagCount int
	for _, nextedBag := range ruleMap[bag] {
		bagCount += nextedBag.amount * (getAllBagsInside(nextedBag.bag, ruleMap) + 1) // +1 to count this bag
	}
	fmt.Printf("A %s bag contains %d \n", bag, bagCount)
	return bagCount
}

func solvePartOne(ruleMap map[string][]Rules) int {
	bagsThatContainGold := make(map[string]bool)
	bagsThatContainGold["shiny gold bag"] = true
	nextLevelOfBags := expand(bagsThatContainGold, ruleMap)
	for len(nextLevelOfBags) != 0 {
		for k, v := range nextLevelOfBags {
			bagsThatContainGold[k] = v
		}
		nextLevelOfBags = expand(nextLevelOfBags, ruleMap)
	}
	return len(bagsThatContainGold) - 1
}

func expand(bagsToFind map[string]bool, allRules map[string][]Rules) map[string]bool {
	expandedBags := make(map[string]bool)
	for key, rules := range allRules {
		for _, rule := range rules {
			if bagsToFind[rule.bag] {
				expandedBags[key] = true
			}
		}
	}
	return expandedBags
}

func buildRulesMap(ruleInput []string) map[string][]Rules {
	rules := make(map[string][]Rules)
	for _, rule := range ruleInput {
		var leftAndRight = regexp.MustCompile(`^(.*) contain(.*)$`)
		captureGroups := leftAndRight.FindStringSubmatch(rule)
		// fmt.Printf("Full Left:%s\n", captureGroups[1])
		// fmt.Printf("Full Right:%s\n", captureGroups[2])
		var left string = captureGroups[1]
		var allRights []Rules = splitRightSide(captureGroups[2])
		rules[left] = allRights
	}
	return rules
}

func splitRightSide(allRights string) []Rules {
	if allRights == "" {
		return make([]Rules, 0)
	}
	var rules []Rules //nolint:prealloc
	allRulesInput := strings.Split(allRights, ",")
	for _, rule := range allRulesInput {
		var amountRegex = regexp.MustCompile(` (\d)* (.*)`)
		amountAndBag := amountRegex.FindStringSubmatch(rule)
		amount, _ := strconv.Atoi(amountAndBag[1])
		bag := amountAndBag[2]
		// fmt.Printf("Right:%d:%d:%s\n", i, amount, bag)
		rule := Rules{
			bag:    bag,
			amount: amount,
		}
		rules = append(rules, rule)
	}
	return rules
}
