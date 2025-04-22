package main

func characterReplacement(s string, k int) int {
	start, ans := 0, 0
	d := make(map[byte]int)
	for end := range s {
		d[s[end]]++

		for end-start+1-maxFreq(d) > k {
			d[s[start]]--
			start++
		}
		if end-start+1 > ans {
			ans = end - start + 1
		}
	}
	return ans
}

func maxFreq(d map[byte]int) int {
	a := 0
	for _, v := range d {
		if v > a {
			a = v
		}
	}
	return a
}

func characterReplacement1(s string, k int) int {
	start, maxCount, ans := 0, 0, 0
	d := make(map[byte]int)
	for end := range s {
		d[s[end]]++
		if d[s[end]] > maxCount {
			maxCount = d[s[end]]
		}
		/*
		maxCount represents the maximum frequency of any character within the current window.
		It's used to determine if the current window is valid (i.e., if we can make all characters the same with at most k replacements).
		Why maxCount Doesn't Need Recalculation

		Optimistic Assumption:

		The algorithm takes an optimistic approach. It assumes that the maxCount from the larger window will still be relevant to the smaller window.
		The while loop only shrinks the window from the left. It doesn't introduce any new characters that could potentially increase the maximum frequency.
		Maintaining Validity:

		Even if maxCount is slightly "outdated" after shrinking the window, the condition right - left + 1 - maxCount > k still serves its purpose.
		If the condition is true, it definitely means that the window is invalid.
		If the condition becomes false, it means the window is now valid, or at least valid enough.
		The reason this works is because if the window was previously valid, and you are shrinking it, you are either maintaining or lowering the number of characters that need to be changed. Therefore, if the window was previously invalid, it is still invalid after shrinking.
		Efficiency:

		Recalculating maxCount within the while loop would require iterating through the charCounts map, which would increase the time complexity of the algorithm.
		By avoiding this recalculation, we maintain the O(n) time complexity.
		In simpler terms:

		Think of it like this:

		If, in a large window, you found that you had, say, 5 'A's (the maximum count), and you need to remove some characters from the left to make it valid, you're not going to suddenly have more 'A's than you did before.
		Therefore, the old maxCount is still a valid upper limit. If the window is still invalid with that maxCount, then it is definetly invalid. If it becomes valid, then we know that the window is now acceptable.
		Key takeaway:

		The algorithm's correctness relies on the fact that maxCount provides a sufficient upper bound for the maximum frequency, and the while loop ensures that the window is shrunk until it becomes valid.
		*/
		for end-start+1-maxCount > k {
			d[s[start]]--
			start++
		}
		if end-start+1 > ans {
			ans = end - start + 1
		}
	}
	return ans
}
