// https://practice.geeksforgeeks.org/problems/alien-dictionary/1
// https://www.codeconvert.ai/golang-to-python-converter
// https://www.youtube.com/watch?v=U3N_je7tWAs&list=PLgUwDviBIf0oE3gA41TKO2H5bHpPd7fzn&index=26

package main

import (
	"fmt"
	"strings"
)

func findOrder(alien_dict []string, N, K int) string {
	// Create adjacency list
	adj := make([][]int, K)
	for i := 0; i < K; i++ {
		adj[i] = []int{}
	}

	// Build adjacency list based on alien dictionary
	for i := 0; i < len(alien_dict)-1; i++ {
		strA := alien_dict[i]
		strB := alien_dict[i+1]
		for j := 0; j < len(strA) && j < len(strB); j++ {
			if strA[j] != strB[j] {
				adj[strA[j]-'a'] = append(adj[strA[j]-'a'], int(strB[j]-'a'))
				break
			}
		}
	}

	// Perform topological sort
	ts := topoSort(adj)

	// Build the resulting string
	sb := strings.Builder{}
	for k := range ts {
		sb.WriteByte(byte('a' + ts[k]))
	}

	return sb.String()
}

func topoSort(adj [][]int) []int {
	inDegree := make([]int, len(adj))
	ans := []int{}
	q := []int{}
	for _, node := range adj {
		for _, neighbor := range node {
			inDegree[neighbor]++
		}
	}
	for k := range inDegree {
		if inDegree[k] == 0 {
			q = append(q, k)
		}
	}
	for len(q) != 0 {
		node := q[0]
		ans = append(ans, node)
		for _, neighbor := range adj[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				q = append(q, neighbor)
			}
		}
		q = q[1:]
	}
	return ans
}

func main() {
	fmt.Println(findOrder([]string{"baa", "abcd", "abca", "cab", "cad", "e"}, 6, 5))
}
