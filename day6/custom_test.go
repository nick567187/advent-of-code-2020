package day6_test

import (
	"testing"

	"github.com/nick567187/advent-of-code-2020/day6"
	"github.com/stretchr/testify/require"
)

func TestCountCustoms(t *testing.T) {
	i := [][]string{{"abc", "a"}}
	r := day6.CountCustoms(i)
	require.Equal(t, 4, r)
}