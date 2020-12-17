package day6_test

import (
	"testing"

	"github.com/nick567187/advent-of-code-2020/day6"
	"github.com/stretchr/testify/require"
)

func TestCountQuestionsAnyoneYes(t *testing.T) {
	i := [][]string{{"abc"}, {"a", "b", "c"}, {"ab", "ac"}, {"a", "a", "a", "a"}, {"b"}}
	r := day6.CountQuestionsAnyoneYes(i)
	require.Equal(t, 11, r)
}

func TestCountQuestionsAnyoneOriginalInput(t *testing.T) {
	i := day6.ReadCustoms()
	r := day6.CountQuestionsAnyoneYes(i)
	require.Equal(t, 6590, r)
}

func TestCountQuestionsEveryoneYes(t *testing.T) {
	i := [][]string{{"abc"}, {"a", "b", "c"}, {"ab", "ac"}, {"a", "a", "a", "a"}, {"b"}}
	r := day6.CountQuestionsEveryoneYes(i)
	require.Equal(t, 6, r)
}

func TestCountQuestionsEveryoneOriginalInput(t *testing.T) {
	i := day6.ReadCustoms()
	r := day6.CountQuestionsEveryoneYes(i)
	require.Equal(t, 3288, r)
}
