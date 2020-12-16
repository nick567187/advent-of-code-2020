package day5_test

import (
	"testing"

	"github.com/nick567187/advent-of-code-2020/day5"
	"github.com/stretchr/testify/require"
)

func TestBinarySearch(t *testing.T) {
	row := day5.BinarySearch("FBFBBFFRLR", 'F', 'B', 0, 7, 127)
	require.Equal(t, 44.0, row)
	col := day5.BinarySearch("FBFBBFFRLR", 'L', 'R', 7, 10, 8)
	require.Equal(t, 5.0, col)
}

func TestGetSeatId(t *testing.T) {
	s := day5.GetSeatId(5, 7)
	require.Equal(t, 47.0, s)
}

func TestGetSeat(t *testing.T) {
	row, col := day5.GetSeat("FBFBBFFRLR")
	require.Equal(t, 44.0, row)
	require.Equal(t, 5.0, col)
}

func TestBinaryBoarding(t *testing.T) {
	var seats = []string{"FBFBBFFRLR", "BFFFBBFRRR", "FFFBBBFRRR"}
	max := day5.BinaryBoarding(seats)
	require.Equal(t, 567.0, max)
}

func TestBinaryBoardingRawInput(t *testing.T) {
	seats, _ := day5.ReadFile()
	max := day5.BinaryBoarding(seats)
	require.Equal(t, 848.0, max)
}

func TestBinaryBoardingMissingNumber(t *testing.T) {
	var seats = []string{"FBFBBFFRLL", "FBFBBFFRRL"}
	miss := day5.BinaryBoardingMissingNumber(seats)
	require.Equal(t, 357.0, miss)
}

func TestBinaryBoardingMissingNumberRawInput(t *testing.T) {
	seats, _ := day5.ReadFile()
	miss := day5.BinaryBoardingMissingNumber(seats)
	require.Equal(t, 682.0, miss)
}