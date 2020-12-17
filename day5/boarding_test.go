package day6_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinarySearch(t *testing.T) {
	row := day5.BinarySearch("FBFBBFFRLR", 'F', 'B', 0, 7, 127)
	require.Equal(t, 44.0, row)
	col := day5.BinarySearch("FBFBBFFRLR", 'L', 'R', 7, 10, 8)
	require.Equal(t, 5.0, col)
}
