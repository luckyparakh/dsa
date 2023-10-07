// https://www.youtube.com/watch?v=m9zawMC6QAI
// https://leetcode.com/problems/implement-trie-prefix-tree/
package main

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{
			children: make([]*Node, 26),
		},
	}
}

func (this *Trie) Insert(word string) {
	tNode := this.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if tNode.children[index] == nil {
			newNode := &Node{children: make([]*Node, 26)}
			tNode.children[index] = newNode
		}
		if i == len(word)-1 {
			tNode.children[index].eow = true
		}
		tNode = tNode.children[index]
	}
}

func (this *Trie) Search(word string) bool {
	tNode := this.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if tNode.children[index] == nil {
			return false
		}
		if i == len(word)-1 {
			if !tNode.children[index].eow {
				return false
			}
		}
		tNode = tNode.children[index]
	}
	return true
}

func (this *Trie) StartsWith(prefix string) bool {
	tNode := this.root
	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a'
		if tNode.children[index] == nil {
			return false
		}
		tNode = tNode.children[index]
	}
	return true
}

type Node struct {
	children []*Node
	eow      bool
}
