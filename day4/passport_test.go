package day4_test

import (
	"testing"

	"github.com/nick567187/advent-of-code-2020/day4"
	"github.com/stretchr/testify/require"
)

func TestComputeValidPassports1(t *testing.T) {
	res := day4.ComputeValidPassports("passports1.txt", true)
	require.Equal(t, res, 1)
}

func TestComputeValidPassportsAllNoValidation(t *testing.T) {
	res := day4.ComputeValidPassports("passports.txt", false)
	require.Equal(t, res, 254)
}

func TestComputeValidPassportsAllWithValidation(t *testing.T) {
	res := day4.ComputeValidPassports("passports.txt", true)
	require.Equal(t, res, 184)
}

func TestComputeValidPassportsAllWithValidationPasses(t *testing.T) {
	res := day4.ComputeValidPassports("passports2.txt", true)
	require.Equal(t, res, 4)
}

func TestComputeValidPassportsAllWithValidationFails(t *testing.T) {
	res := day4.ComputeValidPassports("passports3.txt", true)
	require.Equal(t, res, 0)
}

func TestComputeValidPassportsAllWithValidationPassesOne(t *testing.T) {
	res := day4.ComputeValidPassports("passports1.txt", true)
	require.Equal(t, res, 1)
}

func TestContainsValidFields(t *testing.T) {
	s := "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm"
	res := day4.ContainsValidFields(s)
	require.Equal(t, res, true)
}

func TestContainsValidFieldsFails(t *testing.T) {
	s := "ecl:gry pid:860033327 eyr:2020 hcb:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm"
	res := day4.ContainsValidFields(s)
	require.Equal(t, res, false)
}
