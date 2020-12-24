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

func JoltPart2Helper(input []int, currentValue, currentIndex int, cache []int) int {
	var total int
	if currentIndex > 0 && currentIndex < len(input) && cache[currentIndex] > 0 {
		return cache[currentIndex]
	}
	for nextIndex := currentIndex + 1; nextIndex <= nextIndex+3 && nextIndex<len(input); nextIndex++ {
		nextValue := input[nextIndex]
		if nextValue-currentValue <= 3 {
			helper := JoltPart2Helper(input, nextIndex, nextValue, cache)
			cache[nextIndex] = helper
			total += helper
		}
	}

	return total
}