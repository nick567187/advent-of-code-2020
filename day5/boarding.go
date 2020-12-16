package day5

import (
	"bufio"
	"errors"
	"log"
	"math"
	"os"
)

func ReadFile() (lines []string, e error) {
	file, err := os.Open("boarding.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

func BinaryBoarding(seatIds []string) (max float64){
	for _, s := range seatIds {
		row, col := GetSeat(s)
		if id := GetSeatId(row, col); id > max {
			max = id
		}
	}
	return max
}

func BinaryBoardingMissingNumber(seatIds []string) float64 {
	// n log n time n space: sort then find missing number
	// n time 2n space: loop to find low and max while caching seatIds
	// create an array of that size
	// place items by index
	// loop through to see which index is undefined
	max := 0.0
	min := math.Inf(1)

	memo := make(map[float64]bool)
	for _, s := range seatIds {
		row, col := GetSeat(s)
		id := GetSeatId(row, col)
		memo[id] = true
		if id > max {
			max = id
		}
		if id < min {
			min = id
		}
	}

	arr := make([]float64, int(max-min+1))
	for k, _ := range memo {
		arr[int(k - min)] = 1
	}

	for i, v := range arr {
		if v == 0 {
			return float64(i) + min
		}
	}
	return 0.0
}

func GetSeat(str string) (row float64, col float64) {
	if len(str) != 10 {
		panic(errors.New("invalid seat"))
	}
	row = BinarySearch(str,'F','B',0,7, 127)
	col = BinarySearch(str,'L','R',7,10,8)
	return
}

func BinarySearch(str string, front, back, i, j uint8, upperLimit float64) float64 {
	var upper = upperLimit
	var lower = 0.0
	for ;i < j; i++ {
		if str[i] == front {
			upper = upper - math.Ceil((upper-lower)/2)
		} else if str[i] == back {
			lower = upper - math.Floor((upper-lower)/2)
		}
	}
	return lower
}

func GetSeatId(row float64, col float64) float64 {
	return row * 8 + col
}