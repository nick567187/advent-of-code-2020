package day2_test

import (
	"testing"

	"github.com/nick567187/advent-of-code-2020/day2"
	"github.com/stretchr/testify/require"
)

func TestFormatData(t *testing.T) {
	i := []string{"3-9 n: nnnlncrnnnnn"}
	res := day2.FormatData(i)
	require.Equal(t, res, []day2.Policy{{3, 9, "n", "nnnlncrnnnnn"}})
}

func TestParsePolicyRange(t *testing.T) {
	i := "3-9"
	lower, upper := day2.ParsePolicyRange(i)
	require.Equal(t, lower, 3)
	require.Equal(t, upper, 9)
}

func TestValidPasswords(t *testing.T) {
	i := []day2.Policy{{3, 9, "n", "nnnlncrnnnnn"}, {1, 2, "a", "aaa"}}
	res := day2.ValidPasswords(i)
	require.Equal(t, res, 1)
}

func TestPasswordPassesPolicy(t *testing.T) {
	i := day2.Policy{3, 9, "n", "nnnlncrnnnnn"}
	res := day2.PasswordPassesPolicy(i)
	require.Equal(t, res, true)
}

func TestPasswordPassesPolicyExceedsRange(t *testing.T) {
	i := day2.Policy{1, 3, "n", "nnnlncrnnnnn"}
	res := day2.PasswordPassesPolicy(i)
	require.Equal(t, res, false)
}

func TestPasswordPassesPolicyDoesNotMeetRange(t *testing.T) {
	i := day2.Policy{2, 3, "n", "na"}
	res := day2.PasswordPassesPolicy(i)
	require.Equal(t, res, false)
}

func TestPasswordPassesPolicyDoesNotHaveLetter(t *testing.T) {
	i := day2.Policy{1, 3, "a", "nnnlncrnnnnn"}
	res := day2.PasswordPassesPolicy(i)
	require.Equal(t, res, false)
}