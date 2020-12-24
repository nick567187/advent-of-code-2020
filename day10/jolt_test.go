package day10_test

import (
	"testing"

	"github.com/nick567187/advent-of-code-2020/day10"
	"github.com/stretchr/testify/require"
)

func TestEncodingPart1(t *testing.T) {
	var i = []int{
		16,
		10,
		15,
		5,
		1,
		11,
		7,
		19,
		6,
		12,
		4,
	}
	res := day10.JoltPart1(i)
	require.Equal(t, 35, res)
}

func TestEncodingPart1Longer(t *testing.T) {
	var i = []int{
		28,
		33,
		18,
		42,
		31,
		14,
		46,
		20,
		48,
		47,
		24,
		23,
		49,
		45,
		19,
		38,
		39,
		11,
		1,
		32,
		25,
		35,
		8,
		17,
		7,
		9,
		4,
		2,
		34,
		10,
		3,
	}
	res := day10.JoltPart1(i)
	require.Equal(t, 220, res)
}

func TestEncodingPart1OriginalInput(t *testing.T) {
	i := day10.ReadInput()
	res := day10.JoltPart1(i)
	require.Equal(t, 2380, res)
}

func TestEncodingPart2Short(t *testing.T) {
	var i = []int{
		0,
		16,
		10,
		15,
		5,
		1,
		11,
		7,
		19,
		6,
		12,
		4,
	}
	res := day10.JoltPart2Helper(i)
	require.Equal(t, 8, res)
}

func TestEncodingPart2Longer(t *testing.T) {
	var i = []int{
		28,
		33,
		18,
		42,
		31,
		14,
		46,
		20,
		48,
		47,
		24,
		23,
		49,
		45,
		19,
		38,
		39,
		11,
		1,
		32,
		25,
		35,
		8,
		17,
		7,
		9,
		4,
		2,
		34,
		10,
		3,
	}
	res := day10.JoltPart2Helper(i)
	require.Equal(t, 19208, res)
}

func TestEncodingPart2OriginalInput(t *testing.T) {
	i := day10.ReadInput()
	res := day10.JoltPart2Helper(i)
	require.Equal(t, 2380, res)
}