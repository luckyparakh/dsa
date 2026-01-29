# 2Pointers
## Problem Without 2 pointer
* We brute-force pairs or subarrays with nested loop. Lots of re-scan of same data, o(n2) time.
* Symptoms: double loop, clunky index math, difficulty using sorting.
* Root Cause: 
  * We treat each combination independently, so we repeat work instead of using order/invariants to rule out large groups of possibilities at once.
  * The "fail" was that the first loop didn't "talk" to the second loop. The second loop just blindly checked everything.
## What is it
* Keep 2 indices, move them according to a rule.
* Instead of one person standing still while another person walks the whole line, what if we had two people starting at different spots and moving toward each other based on what they see.
* This will collapse o(n2) into o(n)

## Patterns
* Opposite‑ends pattern: In a sorted array, if current sum is too big, move the right pointer left to get smaller; if too small, move left pointer right to get larger. You never need to revisit rejected regions.
* Sliding‑window pattern: Expand right to explore; when a constraint breaks (sum too big, too many distinct, duplicate appears), shrink left until the invariant/rule holds again. Each index only moves forward.
* Fast/slow pattern: For linked lists, move one pointer 2 steps and the other 1 step to detect cycles or find the middle without extra memory.

## Step‑by‑Step Reasoning Chains:

* Opposite‑ends (e.g., find a pair with target sum in a sorted array)
  1. What do we need? A pair whose sum equals target
  2. What structure helps? Sorting gives monotonic sums as pointers move.
  3. Decision rule: If sum > target, decrease it by moving right pointer left. If sum < target, increase it by moving left pointer right.
  4. Why is this safe? Sorted order guarantees that moving the “too‑large” side is the only way to decrease; we never miss a valid pair.
  5. Complexity: Each pointer moves at most n steps → O(n).
* Sliding window (e.g., longest subarray with sum ≤ K or no duplicates)
  1. Need: Best contiguous window satisfying a constraint.
  2. Explore by expanding right; keep a running state (sum, counts, set).
  3. When the constraint breaks, shrink left to restore it.
  4. Why does this work? The constraint is “shrinkable/monotone”: once invalid, only removing from the left can fix it.
  5. Complexity: Left and right each move at most n steps → O(n).

## When to use
* Data is sorted (or can be sorted) and you need pair/triple comparisons with a monotone predicate (sum, difference, product).
* You’re optimizing a contiguous subarray/string with a “shrinkable” constraint (sum ≤ K, at most K distinct, no duplicates).
* Linked lists questions: cycle detection, middle node, kth from end (fast/slow).
* Merging or intersecting two sorted arrays/lists.

## Mini Checklist:
### 2 pointer
* Is there an order or a shrinkable invariant?
* What invariant will I maintain (sum, distinct count, sorted ordering)? An invariant should be “always true” during the scan and guide movement.
* Exact move rule: when do I move left vs right?
* Does each pointer advance at most n times (no backtracking)?
* Edge cases covered: empty input, single element, duplicates, negatives.

### Sliding Window
* Constraint: shrinkable and monotone under moves (expanding r makes metric non-decreasing; shrinking l makes it non-increasing).
* State: O(1) updates (sum, counts, set/map).
* Move rule: iterate r; add a[r]; while violates, remove a[l] and l++; record best/answer.
* Progress: each index advances at most once; no backtracking.
* Gotchas: negatives break sum monotonicity for “sum ≤ K”; use other techniques (prefix sums) when not shrinkable.

## Problems
### 2 pointer
1. Pair Sum (Sorted Array):
  * Order: array sorted ascending; moving `l` increases sum, moving `r` decreases sum (monotone).
  * Invariant: all candidate pairs remain within [l, r]; rejected regions can’t contain the target given sorted order. Movement keeps the search space valid and shrinks it.
  * Move rule: if `a[l] + a[r] < target` → `l++`; if `a[l] + a[r] > target` → `r--`; if equal → return the pair.
  * Progress: each pointer moves inward at most `n` times; no backtracking.
  * Edge cases: empty/single element; duplicates; very large values. If original indices are required after sorting, carry indices along (not needed here since input is already sorted).

2. Remove Duplicates (Sorted Array):
  * Order: array sorted ascending; duplicates are adjacent.
  * Invariant: `[0...insertIdx-1]` contains unique elements in sorted order; `insertIdx` is where the next unique value goes. Everything before `insertIdx` is finalized; `scanIdx` explores ahead.
  * Move rule: iterate `scanIdx` from 1 to n-1; if `a[scanIdx] != a[scanIdx-1]` (new unique value), copy `a[scanIdx]` to `a[insertIdx]` and `insertIdx++`. Always increment `scanIdx`.
  * Progress: `scanIdx` scans the array once; `insertIdx` advances only when a unique is found; no backtracking.
  * Edge cases: empty array, single element, all duplicates, no duplicates. Return `insertIdx` as the count of unique elements.

3. Squares of Sorted Array:
  * Order: array sorted ascending (may contain negatives); after squaring, largest squares are at the extremes (most negative or most positive values).
  * Invariant: `result[pos...n-1]` is filled with the largest remaining squares in sorted order; search space is `[l, r]`. The largest square among remaining candidates is always at one of the two ends.
  * Move rule: compare `a[l]²` vs `a[r]²`; place the larger one at `result[pos]`; move the pointer that gave the larger square (`l++` if left was larger, `r--` if right was larger); decrement `pos`.
  * Progress: each pointer moves inward at most `n` times; `pos` decrements from `n-1` to `0`; no backtracking. Fill result backward (largest to smallest).
  * Edge cases: empty array, single element, all negatives, all positives, mix of negatives and positives. Complexity: O(n) vs naive square-then-sort O(n log n).
  
4. Triplet Sum to Zero (Unsorted Array):
  * Order: sort array ascending first; then for each fixed element `i`, the remaining subarray `[i+1...n-1]` is sorted, making two-pointer sum monotone.
  * Invariant: for fixed `nums[i]`, all candidate pairs `(l, r)` remain within `[i+1, r]`; rejected regions can't form valid triplets given sorted order. Target is `-nums[i]` for the pair sum. Uniqueness: skip duplicate values at positions `i`, `l`, and `r` to avoid duplicate triplets.
  * Move rule: outer loop iterates `i` from 0 to `n-3` (skip if `nums[i] == nums[i-1]` for `i > 0`); inner two-pointer with `l = i+1`, `r = n-1`: if `nums[i] + nums[l] + nums[r] < 0` → `l++`; if `> 0` → `r--`; if `== 0` → record triplet `[nums[i], nums[l], nums[r]]`, then skip duplicates: `while l < r && nums[l] == nums[l+1]` do `l++`; `while l < r && nums[r] == nums[r-1]` do `r--`; then `l++` and `r--`.
  * Progress: for each fixed `i`, pointers `l` and `r` traverse the subarray once without backtracking; outer loop runs `O(n)` times; total `O(n²)`.
  * Edge cases: fewer than 3 elements (return empty); all duplicates; no valid triplets; mix of negatives, positives, zeros. Complexity: `O(n log n)` sort + `O(n²)` search = `O(n²)` total.
  
5. Triplet Sum Closest to Target (Unsorted Array):
  * Order: sort array ascending first; then for each fixed element `i`, the remaining subarray `[i+1...n-1]` is sorted, making sum adjustments monotone (moving `l` right increases sum, moving `r` left decreases sum).
  * Invariant: `closestSum` tracks the triplet sum with minimum `|sum - target|` seen so far; for fixed `nums[i]`, all candidate pairs `(l, r)` remain within `[i+1, r]`; sorted order ensures we can "walk toward" the target by moving appropriate pointer.
  * Move rule: outer loop iterates `i` from 0 to `n-3`; inner two-pointer with `l = i+1`, `r = n-1`: compute `currentSum = nums[i] + nums[l] + nums[r]`; if `|currentSum - target| < |closestSum - target|` (or equal distance but smaller sum), update `closestSum = currentSum`; if `currentSum < target` → `l++`; if `currentSum > target` → `r--`; if `currentSum == target` → return immediately (can't get closer).
  * Progress: for each fixed `i`, pointers `l` and `r` traverse the subarray once without backtracking, converging toward middle; outer loop runs `O(n)` times; total `O(n²)`.
  * Edge cases: fewer than 3 elements (impossible to form triplet); duplicates (handled naturally, don't skip); exact target match (early return); multiple equally close sums (take smallest). Complexity: `O(n log n)` sort + `O(n²)` search = `O(n²)` total.