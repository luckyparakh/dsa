package main

import (
	"fmt"
	"math"
)

func minEatingSpeed(piles []int, h int) int {
	var max int
	min := math.MaxInt
	for _, v := range piles {
		if v > max {
			max = v
		}
	}
	// fmt.Println(max)
	speedLowerBound, speedUpperBound := 1, max // Important point that lower bound of speed can't be less than 1 as speed can't be 0 or -ve
	for speedLowerBound <= speedUpperBound {
		m := speedLowerBound + (speedUpperBound-speedLowerBound)/2
		// fmt.Println(e, m, s, min)
		timeTaken := 0
		for _, v := range piles {
			timeTaken += int(math.Ceil(float64(v) / float64(m)))
		}
		if timeTaken <= h {
			if min > m {
				min = m
			}
			speedUpperBound = m - 1
		} else {
			speedLowerBound = m + 1
		}
	}
	return min
}

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
}
