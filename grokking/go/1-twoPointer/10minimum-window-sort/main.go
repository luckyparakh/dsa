// https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/problem-challenge-3-minimum-window-sort-medium
package main

import (
	"fmt"
)

type Solution struct{}

func (s *Solution) sort(arr []int) int {
	fmt.Println("-------------")
	low, high := 0, len(arr)-1
	minInWindow := 100_00_00
	maxInWindow := -100_00_00

	// Key Insight:
	// There will only one unsorted subarray in the array
	// So we can find the first element which is out of order from the start
	// and the first element which is out of order from the end
	// The subarray between these two elements is the unsorted subarray

	// Why arr[i] <= arr[i+1] ?
	// Because we are looking for the first element which is out of order
	// So we need to find the first element which is greater than the next element
	// If we use arr[i] < arr[i+1], we will miss the case where there are duplicate elements
	for i := 0; i < len(arr)-1 && arr[i] <= arr[i+1]; i++ {
		low++
	}
	fmt.Println("low", low)
	if low == len(arr)-1 {
		return 0
	}
	for i := len(arr) - 1; i > 0 && arr[i-1] <= arr[i]; i-- {
		high--
	}
	fmt.Println("high", high)
	for i := low; i <= high; i++ {
		if arr[i] < minInWindow {
			minInWindow = arr[i]
		}
		if arr[i] > maxInWindow {
			maxInWindow = arr[i]
		}
	}
	fmt.Println("minWinVal", minInWindow, "maxWinVal", maxInWindow)

	// Why not arr[i] >= minInWindow ?
	// Because we are looking for the first element which is greater than minInWindow
	// to include it in the sorting window
	// This will keep the sorting window as small as possible
	for i := low - 1; i >= 0 && arr[i] > minInWindow; i-- {
		low--
	}
	fmt.Println("lowNew", low)
	for i := high + 1; i <= len(arr)-1 && arr[i] < maxInWindow; i++ {
		high++
	}
	fmt.Println("highNew", high)
	return high - low + 1
}

func main() {
	s := Solution{}
	fmt.Println(s.sort([]int{1, 2, 5, 3, 7, 10, 9, 12})) // 5
	fmt.Println(s.sort([]int{1, 2, 5}))                  // 0
	fmt.Println(s.sort([]int{3, 3, 2, 2}))               // 4
	fmt.Println(s.sort([]int{1, 1, 1}))                  // 0
	fmt.Println(s.sort([]int{2, 3, 3, 2}))               // 3
}

// Minimum Window Sort - Intuition & Solution
//
// PROBLEM: Find the length of the smallest subarray which when sorted will sort the whole array
// Example: [1, 2, 5, 3, 7, 10, 9, 12]
//          If we sort [5, 3, 7, 10, 9] → we get [1, 2, 3, 5, 7, 9, 10, 12] ✓
//          Length = 5
//
// THE JOURNEY TO THE SOLUTION:
//
// 1. NAIVE APPROACH: Try every possible subarray
//    for i := 0; i < n; i++ {
//        for j := i; j < n; j++ {
//            // Check if sorting [i...j] sorts the whole array
//        }
//    }
//    Time: O(n³) — way too slow!
//
// 2. "What's the problem with brute force?"
//    We're checking EVERY possible window independently.
//    We're not using any insight about WHERE the problem actually is!
//
// 3. KEY INSIGHT: The array has THREE regions
//
//    [Sorted Left | Unsorted Middle | Sorted Right]
//     ✓ ✓ ✓ ✓    |  ✗ ✗ ✗ ✗       | ✓ ✓ ✓
//
//    - Left part: Already in correct sorted order
//    - Middle part: These elements are causing trouble!
//    - Right part: Already in correct sorted order
//
//    We only need to sort the MIDDLE part!
//
// 4. "How do I find where the sorted-left ends?"
//
//    Walk from LEFT and ask: "Is each element ≤ the next one?"
//
//    [1, 2, 5, 3, 7, 10, 9, 12]
//     ✓  ✓  ✓  ✗
//
//    1 ≤ 2? YES, keep going
//    2 ≤ 5? YES, keep going
//    5 ≤ 3? NO! Stop here!
//
//    Left sorted region ends at index 1 (element 2)
//    Problem starts at index 2 (element 5)
//
// 5. "How do I find where the sorted-right starts?"
//
//    Walk from RIGHT and ask: "Is each element ≥ the previous one?"
//
//    [1, 2, 5, 3, 7, 10, 9, 12]
//                       ✗  ✓
//
//    12 is fine (nothing after it)
//    9 ≤ 12? YES, keep going left
//    10 ≥ 9? NO! Stop here!
//
//    Right sorted region starts at index 7 (element 12)
//    Problem ends at index 6 (element 9)
//
// 6. "So the problem window is [2...6]? Done?"
//
//    NOT SO FAST! Here's the TRICKY part:
//
//    The elements in [2...6] are: [5, 3, 7, 10, 9]
//
//    Smallest in window: 3
//    Largest in window: 10
//
//    Now think: Where should 3 ACTUALLY go in a sorted array?
//    Where should 10 ACTUALLY go?
//
//    They might need to move BEYOND our initial boundaries!
//
// 7. "Why would they need to move beyond?"
//
//    Look at the LEFT sorted section: [1, 2]
//    Our window has minimum = 3
//    Is 3 bigger than everything in [1, 2]? YES! So it can stay after them ✓
//
//    But what if the left section was [1, 2, 4]?
//    And our window min = 3?
//    Then 3 needs to move LEFT past the 4! We need to expand our window!
//
// 8. ALGORITHM - Finding and Adjusting Boundaries:
//
//    Step 1: Find initial "problem zone" boundaries
//      - Scan LEFT: stop when arr[i] > arr[i+1] (first decrease)
//      - Scan RIGHT: stop when arr[i-1] > arr[i] (first increase backwards)
//      - This gives us initial low and high
//
//    Step 2: Find min and max IN the problem zone
//      - These are the extreme values that need to find their place
//
//    Step 3: Extend LEFT boundary if needed
//      - Walk left from low: if any element > minInWindow, include it
//      - Why? Because it's BIGGER than something that needs to go there!
//
//    Step 4: Extend RIGHT boundary if needed
//      - Walk right from high: if any element < maxInWindow, include it
//      - Why? Because it's SMALLER than something that needs to go there!
//
// 9. WHY THIS IS NOT SLIDING WINDOW:
//
//    SLIDING WINDOW problems:
//      ✓ Window actively SLIDES (expands right, shrinks left)
//      ✓ You TRY many different window positions
//      ✓ You're SEARCHING for the best window dynamically
//      ✓ Window changes based on a running constraint (sum, distinct count)
//
//    THIS problem:
//      ✗ No sliding — we CALCULATE boundaries directly
//      ✗ We don't try different positions — we DETECT the exact one
//      ✗ We're not searching — we're MEASURING where problems are
//      ✗ No running constraint to maintain — we're analyzing structure
//
// 10. THE REAL PATTERN: "Boundary Detection from Both Ends"
//
//     Think of it like:
//
//     Sliding Window = "I'm a scout EXPLORING with a movable window"
//     This Problem = "I'm a surveyor MEASURING where the problem is"
//
//     Both use pointers, but DIFFERENT purposes!
//
// 11. SIMPLE LITMUS TEST:
//
//     Q: "Do I expand/shrink window based on what's inside it?"
//
//     Sliding Window: YES!
//       "Sum too big? Shrink! Sum okay? Expand!"
//
//     This Problem: NO!
//       "I scan once from each side to FIND boundaries,
//        then ADJUST based on extreme values.
//        No dynamic sliding happening!"
//
// EXAMPLE TRACE:
//
// arr = [1, 2, 5, 3, 7, 10, 9, 12]
//
// Step 1: Find initial low (left boundary)
//   i=0: 1 ≤ 2? YES, low++ (low=1)
//   i=1: 2 ≤ 5? YES, low++ (low=2)
//   i=2: 5 ≤ 3? NO! Stop.
//   low = 2
//
// Step 2: Find initial high (right boundary)
//   i=7: 12 (last element, continue)
//   i=6: 9 ≤ 12? YES, high-- (high=6)
//   i=5: 10 ≥ 9? NO! Stop.
//   high = 6
//
// Step 3: Find min and max in window [2...6]
//   Window: [5, 3, 7, 10, 9]
//   minInWindow = 3
//   maxInWindow = 10
//
// Step 4: Extend left boundary?
//   Check left of low: arr[1] = 2
//   Is 2 > 3 (minInWindow)? NO! Don't extend.
//   low stays 2
//
// Step 5: Extend right boundary?
//   Check right of high: arr[7] = 12
//   Is 12 < 10 (maxInWindow)? NO! Don't extend.
//   high stays 6
//
// Answer: high - low + 1 = 6 - 2 + 1 = 5 ✓
//
// WHY arr[i] <= arr[i+1] (not <)?
//   We want to CONTINUE while things are in order.
//   <= handles duplicates: [1, 2, 2, 5] — the 2s are in order!
//   If we used <, we'd stop at first duplicate incorrectly.
//
// WHY arr[i] > minInWindow (not >=)?
//   We extend the boundary when elements are STRICTLY greater than min.
//   If arr[i] == minInWindow, it's already in the right relative position.
//   We only need to include elements that are MISPLACED (bigger than what should come).
//
// COMPLEXITY:
//   Time: O(n) - four linear scans (left, right, find min/max, extend boundaries)
//   Space: O(1) - only using pointers and variables
//
// PATTERN RECOGNITION:
//   ✓ Finding boundaries/ranges in partially sorted data
//   ✓ Detecting "problem zones" by scanning from both ends
//   ✓ Using max-so-far and min-so-far to detect order violations
//   ✓ Adjusting boundaries based on extreme values
//
// NOT SLIDING WINDOW BECAUSE:
//   ✗ No dynamic expansion/shrinkage based on running constraint
//   ✗ Not trying multiple window positions
//   ✗ Calculating exact boundaries, not exploring possibilities
