// https://www.youtube.com/watch?v=aBxjDBC4M1U&list=PLgUwDviBIf0oE3gA41TKO2H5bHpPd7fzn&index=47

package main

import "fmt"

type DS struct {
	size           []int
	ultimateParent []int
}

func (d *DS) unionBySize(u, v int) {
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
		size:           make([]int, nodeCount+1),
		ultimateParent: make([]int, nodeCount+1),
	}
	for i := 0; i < nodeCount+1; i++ {
		ds.ultimateParent[i] = i
		ds.size[i] = 1
	}
	return ds
}
func main() {
	ds := createDS(7)
	ds.unionBySize(1, 2)
	ds.unionBySize(2, 3)
	ds.unionBySize(4, 5)
	ds.unionBySize(6, 7)
	ds.unionBySize(5, 6)
	fmt.Println(ds.isConnected(3, 7))
	ds.unionBySize(3, 7)
	fmt.Println(ds.size, ds.ultimateParent)
}
