// https://leetcode.com/problems/accounts-merge/

package main

import (
	"fmt"
	"sort"
)

// DS represents a disjoint set data structure
type DS struct {
	size           []int
	ultimateParent []int
}

// Union merges two sets by updating the ultimate parent and size
func (d DS) Union(u, v int) {
	pu, pv := d.findUltimateParent(u), d.findUltimateParent(v)
	if pu == pv {
		return
	}
	if d.size[pu] >= d.size[pv] {
		d.ultimateParent[pv] = pu
		d.size[pu] += d.size[pv]
	} else {
		d.ultimateParent[pu] = pv
		d.size[pv] += d.size[pu]
	}
}

// findUltimateParent finds the ultimate parent of a node and performs path compression
func (d DS) findUltimateParent(u int) int {
	if u == d.ultimateParent[u] {
		return u
	}
	d.ultimateParent[u] = d.findUltimateParent(d.ultimateParent[u])
	return d.ultimateParent[u]
}

// isConnected checks if two nodes are connected
func (d DS) isConnected(u, v int) bool {
	return d.ultimateParent[u] == d.ultimateParent[v]
}

// createDS creates a disjoint set with the given node count
func createDS(nodeCount int) DS {
	ds := DS{
		size:           make([]int, nodeCount),
		ultimateParent: make([]int, nodeCount),
	}
	for i := 0; i < nodeCount; i++ {
		ds.size[i] = 1
		ds.ultimateParent[i] = i
	}
	return ds
}
func accountsMerge(accounts [][]string) [][]string {
	// Map to store the index of the account that contains the email
	emailMap := map[string]int{}
	// Create disjoint sets to keep track of the relationships between accounts
	ds := createDS(len(accounts))

	// Iterate through each account
	for node, emails := range accounts {
		// Iterate through each email in the account
		for k, email := range emails {
			// Skip the first element which is the name of the account
			if k != 0 {
				// Check if the email already exists in the map
				if _, found := emailMap[email]; !found {
					// If not found, add the email to the map with the index of the current account
					emailMap[email] = node
				} else {
					// If found, merge the current account with the account that contains the email
					ds.Union(emailMap[email], node)
				}
			}
		}
	}

	// Create a temporary array to store the merged accounts
	temp := make([][]string, len(accounts))
	// Iterate through the ultimate parent of each account
	for k, v := range ds.ultimateParent {
		// If the account is its own ultimate parent, create a new empty array for it
		if k == v {
			temp[k] = make([]string, 0)
			temp[k] = append(temp[k], accounts[k][0]) // Add the account name to the array
		}
	}

	// Iterate through the email map
	for k, v := range emailMap {
		// Find the ultimate parent of the account that contains the email
		pv := ds.findUltimateParent(v)
		// Append the email to the array of the ultimate parent account
		temp[pv] = append(temp[pv], k)
	}

	// Sort the emails in each account array and add them to the final result
	ans := make([][]string, 0)
	for i := 0; i < len(temp); i++ {
		// Check if the length of temp[i] is greater than 0
		if len(temp[i]) > 0 {
			// Sort the elements in temp[i] starting from index 1
			sort.Strings(temp[i][1:])

			// Append temp[i] to ans
			ans = append(ans, temp[i])
		}
	}

	return ans
}

func main() {
	fmt.Println(accountsMerge([][]string{
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"Mary", "mary@mail.com"},
		{"John", "Aohnnybravo@mail.com"}}))
}
