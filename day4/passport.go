package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var Input = [][]string{}

func ComputeValidPassports(str string, validate bool) int {
	file, err := os.Open(str)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var passport string
	var count int
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" && passport != "" {
			if validate {
				if ContainsValidFields(passport) && ValidateFieldValues(passport) {
					count++
				}
			} else {
				if ContainsValidFields(passport) {
					count++
				}
			}
			passport = ""
		}
		if passport == "" {
			passport += text
		} else {
			passport += " " + text
		}
	}

	if passport != "" && ContainsValidFields(passport) && ValidateFieldValues(passport) {
		count++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return count
}

func ContainsValidFields(str string) bool {
	validFields := map[string]bool{"byr": false, "iyr": false, "eyr": false, "hgt": false, "hcl": false, "ecl": false, "pid": false}
	var substr string
	for i := 0; i < len(str); i++ {
		s := string(str[i])
		if len(substr) == 3 {
			if _, ok := validFields[substr]; ok {
				validFields[substr] = true
			}
			substr += s
		} else if s == " " {
			substr = ""
		} else {
			substr += s
		}
	}
	for _, v := range validFields {
		if v == false {
			return false
		}
	}
	return true
}

func ValidateFieldValues(str string) bool {
	s := strings.Split(str, " ")
	for i := 0; i < len(s); i++ {
		a := strings.Split(s[i], ":")
		fmt.Println(s, "error ehre", a[0], a[1])
		if validateFieldValue(a[0], a[1]) == false {
			return false
		}
	}
	return true
}

func validateFieldValue(f string, v string) bool {
	switch f {
	case "byr":
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		return i >= 1920 && i <= 2002
	case "iyr":
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		return i >= 2010 && i <= 2020
	case "eyr":
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		return i >= 2020 && i <= 2030
	case "hgt":
		cmOrIn := ""
		if len(v) <= 2 {
			return false
		}
		cm := lastTwoCharsAreString(v, "cm")
		in := lastTwoCharsAreString(v, "in")
		if cm {
			cmOrIn = "cm"
		}
		if in {
			cmOrIn = "in"
		}
		if !cm && !in {
			return false
		}

		r := []rune(v)
		numStr := string(r[:len(v)-2])
		n, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		if cmOrIn == "cm" {
			return n >= 150 && n <= 193
		} else if cmOrIn == "in" {
			return n >= 59 && n <= 76
		} else {
			return false
		}
	case "hcl":
		if string(v[0]) != "#" {
			return false
		}
		if len(v) != 7 {
			return false
		}
		var count int
		for i := 1; i < 7; i++ {
			s := string(v[i])
			if isLetterAToF(s) || isNumber0Through9(s) {
				count++
			}
		}
		return count == 6
	case "ecl":
		m := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, val := range m {
			if v == val {
				return true
			}
		}
		return false
	case "pid":
		if len(v) != 9 {
			return false
		}
		for i := 0; i < 9; i++ {
			if isNumber0Through9(string(v[i])) == false {
				return false
			}
		}
		return true
	case "cid":
		return true
	default:
		return false
	}
}

func lastTwoCharsAreString(str string, expect string) bool {
	s := string(str[len(str)-2]) + string(str[len(str)-1])
	if s == expect {
		return true
	}
	return false
}

func isNumber0Through9(char string) bool {
	i, err := strconv.Atoi(char)
	if err != nil {
		return false
	}
	return i >= 0 && i <= 9
}

func isLetterAToF(char string) bool {
	return char >= "a" && char <= "f"
}

func FormatInput(input [][]string) [][]string {
	n := make([][]string, len(input))
	for i, v := range input {
		n[i] = strings.Split(v[0], "")
	}
	return n
}

func CalculateSlope(input [][]string, xSlope int, ySlope int) int {
	var count int
	x := 0
	for y := 0; y+1 < len(input); {
		x += xSlope
		y += ySlope
		l := len(input[0]) - 1
		if x > l {
			x = x - l - 1
		}
		if string(input[y][x]) == "#" {
			count++
		}
	}
	return count
}
