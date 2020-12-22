package day8

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Action string
	Value int
}

func ReadCommands() []*Command {
	file, err := os.Open("handheld.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var commands []*Command
	for scanner.Scan() {
		text := scanner.Text()
		commands = ParseAndAddCommand(commands, text)
	}
	return commands
}

func ParseAndAddCommand(arr []*Command, text string) []*Command {
	str := strings.Split(text, " ")
	v, _ := strconv.Atoi(str[1])
	return append(arr, &Command{str[0], v})
}

func ParseCommands(commands []*Command) (accumulator int) {
	var index int
	for {
		command := commands[index]
		if command.Action == "*" {
			return
		} else if command.Action == "acc" {
			accumulator += command.Value
			command.Action = "*"
			index++
		} else if command.Action == "jmp" {
			index += command.Value
		} else if command.Action == "nop" {
			index++
		}
	}
	return
}

func ParseCommandsPart2(commands []*Command) (accumulator int, done bool) {
	done = false
	var index int
	l := len(commands)
	for {
		if index == l {
			done = true
			return
		}

		command := commands[index]
		if command.Action == "*" {
			return
		} else if command.Action == "acc" {
			accumulator += command.Value
			command.Action = "*"
			index++
		} else if command.Action == "jmp" {
			index += command.Value
		} else if command.Action == "nop" {
			index++
		}
	}
	return
}

func ParseCommandsPart2Solution(commands []*Command) int {
	for i, command := range commands {
		if command.Action == "jmp" || command.Action == "nop" {
			c := myCopy(commands)
			if command.Action == "jmp" {
				c[i].Action = "nop"
			} else if command.Action == "nop" {
				c[i].Action = "jmp"
			}
			if acc, done := ParseCommandsPart2(c); done {
				return acc
			}
		}
	}
	return 0
}

func myCopy(src []*Command) []*Command {
	dest := make([]*Command, len(src))
	for i, p := range src {
		if p == nil {
			continue
		}
		v := *p
		dest[i] = &v
	}
	return dest
}