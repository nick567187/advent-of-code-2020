package day10

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func ReadInput() (arr []int) {
	file, err := os.Open("jolt.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		num, _ := strconv.Atoi(text)
		arr = append(arr, num)
	}
	return
}

func JoltPart1(input []int) int {
	sort.Ints(input)
	differences := map[int]int{
		1: 0,
		3: 0,
	}
	var currJolt int
	for _, adapter := range input {
		diff := adapter - currJolt
		differences[diff]++
		currJolt = adapter
	}
	// find cellphone adapter diff
	return differences[1] * (differences[3] + 1)
}

func JoltPart2Helper(input []int) int {
	input = append(input, 0)
	sort.Ints(input)
	input = append(input, input[len(input)-1] + 3)
	return dynamicProgram(input)
}

func memoize(input []int, i int, memo []int) int {
	if v := memo[i]; v != 0 {
		return v
	}
	if i == len(input) - 1 {
		return 1
	}
	total := 0

	if i + 1 < len(input) && input[i + 1] - input[i] <= 3 {
		total += memoize(input, i + 1, memo)
	}
	if i + 2 < len(input) && (input[i + 2] - input[i] <= 3) {
		total += memoize(input, i + 2, memo)
	}
	if i + 3 < len(input) && (input[i + 3] - input[i] <= 3) {
		total += memoize(input, i + 3, memo)
	}
	memo[i] = total
	return total
}

func dynamicProgram(input []int) int {
	dp := make([]int, len(input))
	dp[0] = 1

	for i := 1; i < len(input); i++ {
		currJolt := input[i]
		for j := i - 1; j >= 0; j-- {
			if currJolt - input[j] <= 3 {
				dp[i] += dp[j]
			} else {
				break
			}
		}
	}
	return dp[len(dp)-1]
}