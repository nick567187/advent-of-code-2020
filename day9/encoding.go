package day9

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadInput() (arr []int) {
	file, err := os.Open("encoding.txt")
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

func EncodingPart1(preambleLength int, input []int) int {
	m := make(map[int][]int)
	for i := 0; i < preambleLength; i++ {
		for j := 0; j < preambleLength; j++ {
			if i == j {
				continue
			}
			m[input[i]] = append(m[input[i]], input[i] + input[j])
		}
	}

	var left int
	for i := preambleLength; i < len(input); i++ {
		if !hasSumInMap(input[i], m) {
			return input[i]
		}
		delete(m, input[left])
		addNewKeyToMap(m, input[i])
		left++
	}
	return 0
}

func addNewKeyToMap(m map[int][]int, key int) {
	m[key] = []int{}
	for k, _ := range m {
		m[key] = append(m[key], key + k)
	}
}

func hasSumInMap(num int, m map[int][]int) bool {
	for _, v1 := range m {
		for _, v2 := range v1 {
			if v2 == num {
				return true
			}
		}
	}
	return false
}

// find contiguous sum range to target and add min and max
func EncodingPart2(preambleLength int, input []int) int {
	target := EncodingPart1(preambleLength, input)

	arr, _ := findContiguousSumRange(target, input)
	min, max := MinMax(arr)
	return min + max
}

func findContiguousSumRange(target int, input []int) ([]int, error) {
	var left, right, sum int
	sum = input[left]
	right++
	for right < len(input) {
		if sum == target {
			return input[left:right], nil
		}
		sum += input[right]
		right++
		for sum > target {
			sum -= input[left]
			left++
		}
	}
	return nil, nil
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}