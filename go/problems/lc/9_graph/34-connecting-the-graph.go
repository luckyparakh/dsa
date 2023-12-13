// https://practice.geeksforgeeks.org/problems/connecting-the-graph/1

package main

import "fmt"

type DS struct {
	size           []int
	ultimateParent []int
	extraEdge      int
}

func (d *DS) Union(u, v int) {
	pu, pv := d.findUltimateParent(u), d.findUltimateParent(v)
	if pu == pv {
		d.extraEdge++
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
func (d *DS) findUltimateParent(u int) int {
	if u == d.ultimateParent[u] {
		return u
	}
	d.ultimateParent[u] = d.findUltimateParent(d.ultimateParent[u])
	return d.ultimateParent[u]
}
func (d *DS) isConnected(u, v int) bool {
	return d.ultimateParent[u] == d.ultimateParent[v]
}
func createDS(nodeCount int) *DS {
	ds := &DS{
		size:           make([]int, nodeCount),
		ultimateParent: make([]int, nodeCount),
	}
	for i := 0; i < nodeCount; i++ {
		ds.size[i] = 1
		ds.ultimateParent[i] = i
	}
	return ds
}
func solve(n int, adj [][]int) int {
	ds := createDS(n)
	for _, edge := range adj {
		ds.Union(edge[0], edge[1])
	}
	noEdgeNeed := 0
	for k, v := range ds.ultimateParent {
		if k == v {
			noEdgeNeed++
		}
	}
	// fmt.Println(ds.extraEdge)
	// fmt.Println(ds.ultimateParent)
	if ds.extraEdge >= noEdgeNeed-1 {
		return noEdgeNeed - 1
	}
	return -1
}

func main() {
	fmt.Println(solve(4, [][]int{{0, 1}, {0, 2}, {1, 2}}))
	fmt.Println(solve(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}))
}
