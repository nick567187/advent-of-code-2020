package day7_test

import (
	"testing"

	"github.com/nick567187/advent-of-code-2020/day7"
	"github.com/stretchr/testify/require"
)

func TestParseStringToRule(t *testing.T) {
	i := "muted coral bags contain 1 bright magenta bag, 1 dim aqua bag."
	var rules = make(map[string][]day7.Rule)
	day7.ParseStringToRule(i, rules)
	o := rules["muted coral"]
	require.Equal(t, []day7.Rule{{"bright magenta", 1},{"dim aqua", 1}}, o)
}