package day3_test

import (
	"testing"

	"github.com/nick567187/advent-of-code-2020/day3"
	"github.com/stretchr/testify/require"
)

func TestFormatInput(t *testing.T) {
	res := day3.FormatInput(day3.Input)
	require.Equal(t, res, 5)
}

func TestCalculateSlope(t *testing.T) {
	i := day3.FormatInput(day3.Input)
	res := day3.CalculateSlope(i, 3, 1)
	require.Equal(t, res, 171)
}

func TestCalculateMultipleSlope(t *testing.T) {
	i := day3.FormatInput(day3.Input)
	res := day3.CalculateSlope(i, 1, 1)
	res1 := day3.CalculateSlope(i, 3, 1)
	res2 := day3.CalculateSlope(i, 5, 1)
	res3 := day3.CalculateSlope(i, 7, 1)
	res4 := day3.CalculateSlope(i, 1, 2)
	require.Equal(t, res * res1 * res2 * res3 * res4, 1206576000)
}
