// https://www.youtube.com/watch?v=aBxjDBC4M1U&list=PLgUwDviBIf0oE3gA41TKO2H5bHpPd7fzn&index=47

package main

import "fmt"

type DS struct {
	rank           []int
	ultimateParent []int
}

func (d *DS) unionByRank(u, v int) {
	pu, pv := d.findUltimateParent(u), d.findUltimateParent(v)
	if pu == pv {
		return
	}
	if d.rank[pu] > d.rank[pv] {
		d.ultimateParent[pv] = pu
	} else if d.rank[pu] == d.rank[pv] {
		d.ultimateParent[pv] = pu
		d.rank[pu]++
	} else {
		d.ultimateParent[pu] = pv
	}
}
func (d *DS) findUltimateParent(x int) int {
	if x == d.ultimateParent[x] {
		return x
	}
	d.ultimateParent[x] = d.findUltimateParent(d.ultimateParent[x])
	return d.ultimateParent[x]
}
func (d *DS) isConnected(u, v int) bool {
	return d.findUltimateParent(u) == d.findUltimateParent(v)
}

func createDS(nodeCount int) DS {
	ds := DS{
		rank:           make([]int, nodeCount+1),
		ultimateParent: make([]int, nodeCount+1),
	}
	for i := 0; i < nodeCount+1; i++ {
		ds.ultimateParent[i] = i
	}
	return ds
}
func main() {
	ds := createDS(7)
	ds.unionByRank(1, 2)
	ds.unionByRank(2, 3)
	ds.unionByRank(4, 5)
	ds.unionByRank(6, 7)
	ds.unionByRank(5, 6)
	fmt.Println(ds.isConnected(3, 7))
	ds.unionByRank(3, 7)
	fmt.Println(ds.rank, ds.ultimateParent)
}
