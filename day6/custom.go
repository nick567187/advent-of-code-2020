package day6

import (
	"bufio"
	"log"
	"os"
)

type Groups [][]string
type Group []string

func ReadCustoms() Groups {
	file, err := os.Open("custom.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var groups Groups
	var group Group
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			group = append(group, text)
		} else if text == "" {
			groups = append(groups, group)
			group = Group{}
		}
	}
	if len(group) > 0 {
		groups = append(groups, group)
	}
	return groups
}

func CountQuestionsAnyoneYes(groups Groups) (count int) {
	for _, group := range groups {
		var letters [26]bool
		for _, s := range group {
			for _, v := range s {
				var letterIndex = v - 97
				if !letters[letterIndex] {
					count++
				}
				letters[letterIndex] = true
			}
 		}
	}
	return
}

func CountQuestionsEveryoneYes(groups Groups) (count int) {
	for _, group := range groups {
		var letters [26]int
		var totalInGroup = len(group)
		for _, s := range group {
			for _, v := range s {
				var letterIndex = v - 97
				letters[letterIndex]++
				if letters[letterIndex] == totalInGroup {
					count++
				}
			}
		}
	}
	return
}