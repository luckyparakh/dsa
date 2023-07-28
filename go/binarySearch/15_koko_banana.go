package main

// https://leetcode.com/problems/koko-eating-bananas/
// https://www.youtube.com/watch?v=2JSQIhPcHQg&list=PL_z_8CaSLPWeYfhtuKHj-9MpYb6XQJ_f2&index=21
import "fmt"

func minEatingSpeed(piles []int, h int) int {
	minRateToEat:=1
	maxRateToEat :=findMax(piles)
	return func(s, e int) int {
		min := s
		for s <= e {
			m := s + (e-s)/2
			if isValid(piles, m, h) {
				min = m
				e = m - 1
			} else {
				s = m + 1
			}
		}
		return min
	}(minRateToEat, maxRateToEat)
}
func isValid(piles []int, k, h int) bool {
	d := make([]int, len(piles))
	copy(d, piles)
	neededH := 0
	for i := 0; i < len(d); i++ {
		// Consider if d[i]=15 and k=7; then t=2 and r=1; it means koko will need 3 iter to eat
		// Consider if d[i]=14 and k=7; then t=2 and r=0; it means koko will need 2 iter to eat
		// So r != 0 means we need one more iter to eat remaining bananas
		t := d[i] / k
		r := d[i] % k
		if r == 0 {
			neededH += t
		} else {
			neededH += t + 1
		}
		if neededH > h {
			return false
		}
	}
	return true
}
func findMax(nums []int) int {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Println(minEatingSpeed([]int{30, 30, 30}, 30))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
	fmt.Println(minEatingSpeed([]int{312884470}, 312884469))
	fmt.Println(minEatingSpeed([]int{873375536, 395271806, 617254718, 970525912, 634754347, 824202576, 694181619, 20191396, 886462834, 442389139, 572655464, 438946009, 791566709, 776244944, 694340852, 419438893, 784015530, 588954527, 282060288, 269101141, 499386849, 846936808, 92389214, 385055341, 56742915, 803341674, 837907634, 728867715, 20958651, 167651719, 345626668, 701905050, 932332403, 572486583, 603363649, 967330688, 484233747, 859566856, 446838995, 375409782, 220949961, 72860128, 998899684, 615754807, 383344277, 36322529, 154308670, 335291837, 927055440, 28020467, 558059248, 999492426, 991026255, 30205761, 884639109, 61689648, 742973721, 395173120, 38459914, 705636911, 30019578, 968014413, 126489328, 738983100, 793184186, 871576545, 768870427, 955396670, 328003949, 786890382, 450361695, 994581348, 158169007, 309034664, 388541713, 142633427, 390169457, 161995664, 906356894, 379954831, 448138536}, 943223529))
}
