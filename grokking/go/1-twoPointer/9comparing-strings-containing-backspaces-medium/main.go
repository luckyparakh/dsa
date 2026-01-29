// https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/problem-challenge-2-comparing-strings-containing-backspaces-medium
package main

// Solution struct simulates a class in Go.
type Solution struct{}

// compare is a method of Solution struct.
// It uses two pointer approach to compare the strings.
func (s *Solution) compare(str1, str2 string) bool {
	l, r := len(str1)-1, len(str2)-1

	// both || or && works
	for l >= 0 && r >= 0 {
		l = s.getNextValidCharIndex(str1, l)
		r = s.getNextValidCharIndex(str2, r)

		// it means both strings are exhausted
		// i.e. l and r are -1 which can happen if there are
		// more or equal backspaces than characters
		if l < 0 && r < 0 {
			return true
		}

		// if one string is exhausted before the other 
		// then they are not equal as we can't compare
		if l < 0 || r < 0 {
			return false
		}

		if str1[l] != str2[r] {
			return false
		}
		l--
		r--
	}
	return true
}

// getNextValidCharIndex finds the next valid character index in the string
// by skipping backspaces and the characters they delete.
// index value can be -1 if number of backspaces are more than or equal to characters.
func (s *Solution) getNextValidCharIndex(str string, index int) int {
	skipCount := 0
	for index >= 0 {
		if str[index] == '#' {
			skipCount++
			index--
		} else if skipCount > 0 {
			skipCount--
			index--
		} else {
			break
		}
	}
	return index
}

// Comparing Strings Containing Backspaces - Intuition & Solution
//
// PROBLEM: Given two strings with '#' (backspace), check if they're equal
// Example: "xy#z" → "xz", "xzz#" → "xz" → Equal!
//          "xy#z" → "xz", "xyz#" → "xy" → Not equal!
//
// THE JOURNEY TO THE SOLUTION:
//
// 1. NAIVE APPROACH - Build the actual strings
//    - Use a stack or array to process each string
//    - For each char: if it's '#', pop; otherwise push
//    - Compare the final strings
//    Time: O(n + m), Space: O(n + m) for two new strings
//
// 2. "Can I do better with space?"
//    The problem says we need to compare. Do I really need to BUILD both strings?
//    What if I could compare character-by-character WITHOUT building?
//
// 3. KEY INSIGHT: Process strings BACKWARDS!
//    Why backwards? Because backspace affects characters to the LEFT!
//
//    "abc#d" → when we see '#' at index 3, it deletes 'c' at index 2
//
//    If we go LEFT-TO-RIGHT: we don't know if a character will be deleted later
//    If we go RIGHT-TO-LEFT: we know immediately how many chars to skip!
//
// 4. THE "AHA!" MOMENT:
//    When moving backwards and we see '#':
//    - Count how many backspaces we have
//    - Skip that many regular characters
//    - Continue until we find a "valid" character (one that survives)
//
// 5. MENTAL MODEL - Two pointers scanning backwards:
//
//    str1: "ab##c"
//               ↑ start here, move left
//    str2: "ad#c"
//              ↑ start here, move left
//
//    Step 1: Find next valid char in str1
//      - See '#' at index 3 → skip count = 1
//      - See '#' at index 2 → skip count = 2
//      - See 'b' at index 1 → skip it (skip count = 1)
//      - See 'a' at index 0 → skip it (skip count = 0)
//      → No valid char! (or we've exhausted the string)
//
//    Actually let's trace "xy#z" vs "xyz#":
//
//    str1 = "xy#z"
//           0123
//
//    Start from end (index 3):
//    - index 3: 'z' → valid! Compare this
//    - Move to index 2: '#' → skip count = 1
//    - Move to index 1: 'y' → skip it (skip count becomes 0)
//    - Move to index 0: 'x' → valid! Compare this
//
//    str2 = "xyz#"
//           0123
//
//    Start from end (index 3):
//    - index 3: '#' → skip count = 1
//    - Move to index 2: 'z' → skip it (skip count becomes 0)
//    - Move to index 1: 'y' → valid! Compare this
//
//    Compare: str1='z' vs str2='y' → NOT EQUAL!
//
// 6. ALGORITHM STRUCTURE:
//    a) Start two pointers at the end of both strings
//    b) For each string, find the next "valid" character:
//       - Count consecutive '#' characters
//       - Skip that many regular characters
//       - Return the index of the valid character
//    c) Compare the valid characters
//    d) If different → return false
//    e) If same → move both pointers left and repeat
//    f) If both pointers reach the start → strings are equal!
//
// 7. HELPER FUNCTION: getNextValidCharIndex(str, index)
//    Purpose: Given a string and current index, find the next valid char moving left
//
//    Logic:
//    - skipCount = 0
//    - While index >= 0:
//        - If str[index] == '#': skipCount++, move left
//        - Else if skipCount > 0: this char is deleted, skipCount--, move left
//        - Else: this char is valid! return index
//    - Return -1 (no valid char found)
//
// 8. WHY THIS WORKS:
//    - Backspace affects LEFT characters only
//    - By going right-to-left, we know exactly what survives
//    - We compare only surviving characters
//    - No extra space needed (just counters and pointers)
//
// 9. EDGE CASES:
//    - All backspaces: "##" vs "" → both empty, equal
//    - More backspaces than chars: "a###" → empty
//    - Different lengths after processing: "ab#" vs "a" → both "a", equal
//    - Empty strings: "" vs "" → equal
//    - No backspaces: "abc" vs "abc" → equal
//
// COMPLEXITY:
// Time: O(n + m) - we traverse each string once from right to left
// Space: O(1) - only use pointers and counters, no extra strings built
//
// PATTERN RECOGNITION - Use this approach when:
// ✓ Problem involves "undoing" or "canceling" operations
// ✓ Effect is directional (backspace affects left)
// ✓ Need O(1) space instead of building result
// ✓ Can process in reverse to know final state
//
// WHY NOT BUILD STRINGS?
// Building: O(n + m) space, O(n + m) time
// Two pointers: O(1) space, O(n + m) time
// Better space complexity!
