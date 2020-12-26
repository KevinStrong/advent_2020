package main

import (
	"advent_2020/input"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	allergies, _, recipes := parseInput(input.ReadLines("day21/input.txt"))
	fmt.Println(allergies)
	// fmt.Println(allIngredients)
	allergyToIngredient := make(map[string]string)
	ingredientsWithAllergies := make(map[string]bool)

	wasMatchFound := true
	for wasMatchFound {
		var ingredientWithAllergy []string
		wasMatchFound, ingredientWithAllergy = collapseIngredientList(allergies, allergyToIngredient)
		for _, ingredient := range ingredientWithAllergy {
			ingredientsWithAllergies[ingredient] = true
		}
	}
	fmt.Println(countRemainingIngredients(recipes, ingredientsWithAllergies))
}

func countRemainingIngredients(recipes [][]string, ingredientsWithAllergies map[string]bool) int {
	count := 0
	for _, ingredientList := range recipes {
		for _, ingredient := range ingredientList {
			if !ingredientsWithAllergies[ingredient] {
				count++
			}
		}
	}
	return count
}

func collapseIngredientList(allergies map[string][][]string, allergyToIngredient map[string]string) (bool, []string) {
	isMatchFound := false
	ingredientThatHasAnAllergy := make([]string, 0)
	for allergy, possibleIngredientLists := range allergies {
		if allergyToIngredient[allergy] == "" {
			intersection := make([]string, 0)
			intersection = append(intersection, possibleIngredientLists[0]...)
			for i := range possibleIngredientLists {
				intersection = createIntersection(intersection, possibleIngredientLists[i])
			}
			if len(intersection) == 1 {
				fmt.Println("Found match for: " + allergy + ":" + intersection[0])
				isMatchFound = true
				ingredientThatHasAnAllergy = append(ingredientThatHasAnAllergy, intersection[0])
				removeIngredientAsPossibleFromAllergies(allergies, intersection[0])
				allergyToIngredient[allergy] = intersection[0]
			}
		}
	}
	return isMatchFound, ingredientThatHasAnAllergy
}

func removeIngredientAsPossibleFromAllergies(allergies map[string][][]string, ingredient string) {
	for allergy := range allergies {
		for ingredientSet := range allergies[allergy] {
			if contains(allergies[allergy][ingredientSet], ingredient) {
				allergies[allergy][ingredientSet] = remove(allergies[allergy][ingredientSet], ingredient)
			}
		}
	}
}

func remove(ingredients []string, ingredient string) []string {
	for i, other := range ingredients {
		if other == ingredient {
			return append(ingredients[:i], ingredients[i+1:]...)
		}
	}
	panic("found but failed to remove:" + ingredient)
}

func contains(ingredients []string, ingredient string) bool {
	for i := range ingredients {
		if ingredients[i] == ingredient {
			return true
		}
	}
	return false
}

func createIntersection(first []string, second []string) []string {
	inFirst := map[string]bool{}
	inSecond := map[string]bool{}
	for _, c := range first {
		inFirst[c] = true
	}

	for _, c := range second {
		inSecond[c] = true
	}

	var result = make([]string, 0)
	for key := range inFirst {
		if inSecond[key] {
			result = append(result, key)
		}
	}
	return result
}

func parseInput(lines []string) (map[string][][]string, []string, [][]string) {
	allIngredients := make(map[string]bool)
	allRecipies := make([][]string, 0)
	allergies := make(map[string][][]string)
	for i := range lines {
		var maskValue = regexp.MustCompile(`^([\w ]+) \(contains ([\w ,]+)\)$`)
		captureGroups := maskValue.FindStringSubmatch(lines[i])
		if len(captureGroups) != 3 {
			panic(strconv.Itoa(len(captureGroups)) + lines[i])
		}
		addToIngredientList(allIngredients, captureGroups[1])
		addToAllergyMap(allergies, captureGroups[1:])
		allRecipies = addToRecipeList(allRecipies, captureGroups[1])
	}
	return allergies, convertToList(allIngredients), allRecipies
}

func addToRecipeList(recipes [][]string, s string) [][]string {
	recipe := strings.Split(s, " ")
	return append(recipes, recipe)
}

func convertToList(ingredients map[string]bool) []string {
	list := make([]string, len(ingredients))
	i := 0
	for s := range ingredients {
		list[i] = s
		i++
	}
	return list
}

func addToAllergyMap(allergies map[string][][]string, ingredientsAndAllergies []string) {
	// for every allergy, add a new entry to the allergies map with all the ingredients
	ingredientsToAdd := strings.Split(ingredientsAndAllergies[0], " ")
	allergiesToAdd := strings.Split(ingredientsAndAllergies[1], ", ")
	for _, currentAllergy := range allergiesToAdd {
		allergies[currentAllergy] = append(allergies[currentAllergy], makeACopy(ingredientsToAdd))
	}
}

func makeACopy(add []string) []string {
	theCopy := make([]string, len(add))
	copy(theCopy, add)
	return theCopy
}

func addToIngredientList(ingredients map[string]bool, s string) {
	ingredientsToAdd := strings.Split(s, " ")
	for i := range ingredientsToAdd {
		ingredients[ingredientsToAdd[i]] = true
	}
}
