// https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/dutch-national-flag-problem-medium
package main

// Solution struct to hold methods
type Solution struct{}

func (s *Solution) sort(arr []int) []int {
	l, i, r := 0, 0, len(arr)-1
	for i <= r {
		switch arr[i] {
		case 0:
			arr[l], arr[i] = arr[i], arr[l]
			i++
			l++
		case 1:
			i++
		case 2:
			arr[r], arr[i] = arr[i], arr[r]
			r--
		}
	}
	return arr
}


// Dutch National Flag Problem - Intuition & Solution
//
// PROBLEM: Sort array containing only 0s, 1s, and 2s in-place
//
// NAIVE APPROACH: Use sort.Ints() → O(n log n)
// But we're not leveraging that we only have 3 distinct values!
//
// KEY INSIGHT: Instead of "sorting", think "partitioning into 3 regions"
// - All 0s should go LEFT
// - All 2s should go RIGHT
// - All 1s naturally end up in MIDDLE
//
// STRATEGY: Maintain 3 pointers to track region boundaries
//
// Visual Model:
// [0, 0, 0, | 1, 1, 1, | ?, ?, ?, | 2, 2, 2]
//           ↑           ↑          ↑
//          low          i         high
//
// Regions (Invariants - always true during algorithm):
// 1. arr[0...low-1]     = all 0s (finalized)
// 2. arr[low...i-1]     = all 1s (finalized)
// 3. arr[i...high]      = unknown (not yet processed)
// 4. arr[high+1...n-1]  = all 2s (finalized)
//
// DECISION RULES:
// When arr[i] == 0:
//   - Belongs in LEFT region
//   - Swap with arr[low]
//   - Expand 0s region: low++
//   - Move forward: i++ (safe because arr[low] was a 1 or 0)
//
// When arr[i] == 1:
//   - Already in correct region
//   - Just move forward: i++
//
// When arr[i] == 2:
//   - Belongs in RIGHT region
//   - Swap with arr[high]
//   - Expand 2s region: high--
//   - DON'T move i (need to examine what we just swapped from high!)
//
// WHY i++ for 0 but NOT for 2?
// - After swapping with low: we get element from "1s region" (known, safe)
// - After swapping with high: we get element from "unknown region" (must check!)
//
// TRACE EXAMPLE:
// arr = [1, 0, 2, 1, 0]
// 
// Initial: low=0, i=0, high=4
// [1, 0, 2, 1, 0]
//
// i=0, arr[0]=1 → just move i
// [1, 0, 2, 1, 0], low=0, i=1, high=4
//
// i=1, arr[1]=0 → swap with low, move both
// [0, 1, 2, 1, 0], low=1, i=2, high=4
//
// i=2, arr[2]=2 → swap with high, move high only
// [0, 1, 0, 1, 2], low=1, i=2, high=3
//
// i=2, arr[2]=0 → swap with low, move both
// [0, 0, 1, 1, 2], low=2, i=3, high=3
//
// i=3, arr[3]=1 → just move i
// [0, 0, 1, 1, 2], low=2, i=4, high=3
//
// i=4 > high=3 → STOP!
// Result: [0, 0, 1, 1, 2] ✓
//
// COMPLEXITY:
// Time: O(n) - each element examined at most twice
// Space: O(1) - in-place swaps only
//
// PATTERN RECOGNITION - Use this when:
// ✓ Partitioning into exactly 3 groups
// ✓ Only 3 distinct values/categories
// ✓ Need in-place solution
// ✓ One pass desired
//
// COMMON MISTAKES:
// ❌ Moving i after swapping with high
// ❌ Using i < high instead of i <= high
// ✓ Always use i <= high to process all elements