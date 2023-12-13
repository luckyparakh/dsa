// https://www.youtube.com/watch?v=DMnDM_sxVig&list=PLgUwDviBIf0oE3gA41TKO2H5bHpPd7fzn&index=47

package main

import (
	"fmt"
	"sort"
)

type DS struct {
	size           []int
	ultimateParent []int
}

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
func (d DS) findUltimateParent(u int) int {
	if u == d.ultimateParent[u] {
		return u
	}
	d.ultimateParent[u] = d.findUltimateParent(d.ultimateParent[u])
	return d.ultimateParent[u]
}
func (d DS) isConnected(u, v int) bool {
	return d.ultimateParent[u] == d.ultimateParent[v]
}
func createDS(nodeCount int) DS {
	ds := DS{
		size:           make([]int, nodeCount+1),
		ultimateParent: make([]int, nodeCount+1),
	}
	for i := 0; i < nodeCount+1; i++ {
		ds.size[i] = 1
		ds.ultimateParent[i] = i
	}
	return ds
}
func spanningTree(V int, edges [][]int) int {
	sum := 0
	ds := createDS(V)
	mst := make([][]int, 0)

	sort.SliceStable(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})
	for _, edge := range edges {
		if !ds.isConnected(edge[0], edge[1]) {
			ds.Union(edge[0], edge[1])
			sum += edge[2]
			mst = append(mst, []int{edge[0], edge[1]})
		}
	}
	fmt.Println(mst)
	return sum
}

func main() {
	fmt.Println(spanningTree(3, [][]int{{0, 1, 5}, {1, 2, 3}, {0, 2, 1}})) //4
	fmt.Println(spanningTree(2, [][]int{{0, 1, 5}}))                       //5
}
