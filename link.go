// Copyright 2017 Simon Cai <xingwangc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fantree

import (
	"fmt"
	"sync"
)

// LinkNode is the basic data structure of the link
type LinkNode struct {
	Name     string    //Name of LinkNode should be the name of node
	Value    *Node     //Pointer to data node
	Previous *LinkNode //Every LinkNode at most has 1 previous
	Next     *LinkNode //Every LinkNode at most has 1 next
}

// NewLinkNode construct a new LinkNode based on the Node
func NewLinkNode(nd *Node) *LinkNode {
	linkNode := new(LinkNode)
	linkNode.Name = nd.Name
	linkNode.Value = nd

	return linkNode
}

// Tail will find the tail of the Link
func (lk *LinkNode) Tail() (tail *LinkNode, err error) {
	if lk == nil {
		return tail, fmt.Errorf("Link is empty!")
	}

	if lk.Next == nil {
		return lk, nil
	}

	return lk.Next.Tail()
}

// MergeLink merge multiple links to 1 link. The rules are:
//1. Suppose every link in the list has the same prority. So randomly
// choose one as the begining.
//2. If 1 node exists in more than 1 links, links will be merged around
// around that node.
//3. If nodes in 2 links are irrelevant, they will be merged alternately.
func MergeLink(heads []*LinkNode) (head *LinkNode, err error) {
	if len(heads) == 0 {
		return head, fmt.Errorf("Should provide at least 1 head")
	}

	//Delete the uninitialized link
	for i := 0; i < len(heads); {
		if heads[i] == nil {
			heads = append(heads[:i], heads[i+1:]...)
			continue
		}
		i++
	}

	head = heads[0]
	if len(heads) == 1 {
		return head, nil
	}

	for i := 1; i < len(heads); i++ {
		position := head

		for current := heads[i]; current != nil; {
			findFlag := false
			next := current.Next
			for node := position; node != nil; node = node.Next {
				if current.Value == node.Value {
					position = node
					findFlag = true
					break
				}
			}
			if !findFlag {
				tmp := position.Next
				position.Next = current
				current.Previous = position
				if tmp == nil {
					break
				} else {
					current.Next = tmp
					tmp.Previous = current
					position = tmp
				}
			}
			current = next
		}
	}

	return head, nil
}

// NewLink will construct a link from a Node list
func NewLink(nodeList []*Node) (head *LinkNode, err error) {
	forest, err := NewForest(nodeList)
	if err != nil {
		return nil, err
	}

	return forest.ToLink()
}

// Print the Link
func (lk *LinkNode) Print() {
	for node := lk; node != nil; node = node.Next {
		fmt.Printf("%s -> ", node.Name)
	}
	fmt.Printf("\n")
}

// Pipeline of the Link will start goroutines to execute user defined
// handlers. But all goroutines are executing synchronously in sequence
// the link defined.
func (lk *LinkNode) Pipeline(metadata interface{}) (output interface{}, err error) {
	input := make(chan interface{})

	go func() {
		input <- metadata
	}()

	var wg sync.WaitGroup
	for lp := lk; lp != nil; lp = lp.Next {
		node := lp.Value
		inC := make(chan chan interface{}, 1)

		if lp.Previous == nil {
			inC <- input
		} else {
			inC <- lp.Previous.Value.OutC
		}
		close(inC)

		if lp.Next == nil {
			wg.Add(1)
			go func() {
				output = <-node.OutC
				wg.Done()
			}()
		}

		wg.Add(1)
		go func() {
			node.Cmd.Handler(node, inC, node.OutC)
			wg.Done()
		}()
	}
	wg.Wait()
	return output, nil
}
