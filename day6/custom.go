package day6

import (
	"bufio"
	"log"
	"os"
)

func ReadCustoms() [][]string {
	file, err := os.Open("custom.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var groups [][]string
	var group []string
	var i int
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			group = append(group, text)
		} else if text == "" {
			groups[i] = group
			group = []string{}
		}
	}
	return groups
}

func CountCustoms(customs [][]string) (count int) {
	for _, v := range customs {
		for range v {
			count++
		}
	}
	return
}