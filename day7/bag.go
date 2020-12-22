package day7

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rules [][]string

type Rule struct {
	Color string
	NumBags int
}

func ReadRules() map[string][]Rule {
	file, err := os.Open("bag_rules.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rules = make(map[string][]Rule)
	for scanner.Scan() {
		text := scanner.Text()
		if err := ParseStringToRule(text, rules); err != nil {
			panic(err)
		}
	}
	return rules
}

func Part1(graph map[string][]Rule, target string) (count int) {
	for key, rules := range graph {
		if target == key {
			continue
		}
		found := false
		for _, rule := range rules {
			if found {
				break
			}
			if target == rule.Color || Dfs(graph, rule.Color, target) {
				found = true
				count++
			}
		}
	}
	return
}

func Dfs(graph map[string][]Rule, color, target string) bool {
	rules, ok := graph[color]
	if !ok {
		return false
	}

}

func ParseStringToRule(str string, rules map[string][]Rule) error {
	bags := strings.Split(str, "bags contain")
	target := strings.TrimSpace(bags[0])
	contains := strings.Split(strings.Trim(bags[1], " ."), "bag")

	var out []Rule
	for _, bag := range contains {
		if bag == "" {
			continue
		}
		trim := strings.Trim(bag, " .,")
		num, _ := strconv.Atoi(string(trim[0]))
		trim = trim[2:]
		out = append(out, Rule{trim, num})
	}
	rules[target] = out
	return nil
}
