package day1_test

import (
	"testing"

	"github.com/nick567187/advent-of-code-2020/day1"
	"github.com/stretchr/testify/require"
)

func TestFindsTwoNumbersThatSumTo2020AndMultiplyThem(t *testing.T) {
	i := []int{1920,89,100,50}
	res := day1.FindTwoThatSum2020AndMultiply(i)
	require.Equal(t, res, 1920 * 100)
}

func TestFindsTwoNumbersThatSumTo2020AndMultiplyThemOriginalInput(t *testing.T) {
	res := day1.FindTwoThatSum2020AndMultiply(day1.Input)
	require.Equal(t, res, 842016)
}

func TestFindsThreeNumbersThatSumTo2020AndMultiplyThem(t *testing.T) {
	i := []int{1920,89,99,50,1}
	res := day1.FindThreeThatSum2020AndMultiply(i)
	require.Equal(t, res, 1920 * 99 * 1)
}

func TestFindsThreeNumbersThatSumTo2020AndMultiplyThemOriginalInput(t *testing.T) {
	res := day1.FindThreeThatSum2020AndMultiply(day1.Input)
	require.Equal(t, res, 9199664)
}

