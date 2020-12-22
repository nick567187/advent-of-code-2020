package day8_test

import (
	"testing"

	"github.com/nick567187/advent-of-code-2020/day8"
	"github.com/stretchr/testify/require"
)

func TestParseAndAddCommand(t *testing.T) {
	var command []*day8.Command
	command = day8.ParseAndAddCommand(command, "jmp +1")
	require.Equal(t, []day8.Command{{"jmp", 1}}, command)

	command = day8.ParseAndAddCommand(command, "acc -34")
	require.Equal(t, []day8.Command{{"jmp", 1}, {"acc", -34}}, command)
}

func TestParseCommands(t *testing.T) {
	var commands = []*day8.Command{
		{"nop", 0},
		{"acc", 1},
		{"jmp", 4},
		{"acc", 3},
		{"jmp", -3},
		{"acc", -99},
		{"acc", 1},
		{"jmp", -4},
		{"acc", 6},
	}
	res := day8.ParseCommands(commands)
	require.Equal(t, 5, res)
}

func TestParseCommandsOriginalInput(t *testing.T) {
	var commands = day8.ReadCommands()
	res := day8.ParseCommands(commands)
	require.Equal(t, 1939, res)
}

func TestParseCommandsPart2Finishes(t *testing.T) {
	var commands = []*day8.Command{
		{"nop", 0},
		{"acc", 1},
		{"jmp", 4},
		{"acc", 3},
		{"jmp", -3},
		{"acc", -99},
		{"acc", 1},
		{"nop", -4},
		{"acc", 6},
	}
	res, done := day8.ParseCommandsPart2(commands)
	require.Equal(t, true, done)
	require.Equal(t, 8, res)
}

func TestParseCommandsPart2Solution(t *testing.T) {
	var commands = day8.ReadCommands()
	res := day8.ParseCommandsPart2Solution(commands)
	require.Equal(t, 2212, res)
}