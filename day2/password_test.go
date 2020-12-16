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
	res := day2.ValidPasswords(i, day2.PasswordPassesPart1Policy)
	require.Equal(t, res, 1)
}

func TestValidPasswordsPart1OriginalInput(t *testing.T) {
	f := day2.FormatData(day2.Input)
	res := day2.ValidPasswords(f, day2.PasswordPassesPart1Policy)
	require.Equal(t, res, 378)
}


func TestValidPasswordsPart2OriginalInput(t *testing.T) {
	f := day2.FormatData(day2.Input)
	res := day2.ValidPasswords(f, day2.PasswordPassesPart2Policy)
	require.Equal(t, res, 280)
}

func TestPasswordPassesPart1Policy(t *testing.T) {
	i := day2.Policy{3, 9, "n", "nnnlncrnnnnn"}
	res := day2.PasswordPassesPart1Policy(i)
	require.Equal(t, res, true)
}

func TestPasswordPassesPart1PolicyExceedsRange(t *testing.T) {
	i := day2.Policy{1, 3, "n", "nnnlncrnnnnn"}
	res := day2.PasswordPassesPart1Policy(i)
	require.Equal(t, res, false)
}

func TestPasswordPassesPart1PolicyDoesNotMeetRange(t *testing.T) {
	i := day2.Policy{2, 3, "n", "na"}
	res := day2.PasswordPassesPart1Policy(i)
	require.Equal(t, res, false)
}

func TestPasswordPassesPart1PolicyDoesNotHaveLetter(t *testing.T) {
	i := day2.Policy{1, 3, "a", "nnnlncrnnnnn"}
	res := day2.PasswordPassesPart1Policy(i)
	require.Equal(t, res, false)
}

func TestPasswordPassesPart2Policy(t *testing.T) {
	i := day2.Policy{1, 3, "a", "abcde"}
	res := day2.PasswordPassesPart2Policy(i)
	require.Equal(t, res, true)
}

func TestPasswordPassesPart2PolicyFails(t *testing.T) {
	i := day2.Policy{1, 3, "b", "cdefg"}
	res := day2.PasswordPassesPart2Policy(i)
	require.Equal(t, res, false)
}

func TestPasswordPassesPart2PolicyDoesNotMeetRange(t *testing.T) {
	i := day2.Policy{2, 9, "c", "ccccccccc"}
	res := day2.PasswordPassesPart2Policy(i)
	require.Equal(t, res, false)
}