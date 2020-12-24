package day9_test

import (
	"testing"

	"github.com/nick567187/advent-of-code-2020/day9"
	"github.com/stretchr/testify/require"
)

func TestEncodingPart1(t *testing.T) {
	var i = []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}
	res := day9.EncodingPart1(5, i)
	require.Equal(t, 127, res)
}

func TestEncodingPart1OriginalInput(t *testing.T) {
	i := day9.ReadInput()
	res := day9.EncodingPart1(25, i)
	require.Equal(t, 22477624, res)
}

func TestEncodingPart2(t *testing.T) {
	var i = []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}
	res := day9.EncodingPart2(5, i)
	require.Equal(t, 62, res)
}

func TestEncodingPart2OriginalInput(t *testing.T) {
	i := day9.ReadInput()
	res := day9.EncodingPart2(25, i)
	require.Equal(t, 2980044, res)
}
