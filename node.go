// Copyright 2017 Simon Cai <xingwangc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fantree

// NodeHander is a type of handler function. User just need to follow the
// signature to implement a custom function for the node to handle specific
// computation.
type NodeHander func(self *Node, in chan chan interface{}, out chan interface{})

//Node is the basic structure of fantree.
type Node struct {
	PreName  string           // The name of previous node
	NextName string           // The name of next node
	Name     string           // Name of the node, which should be unquely identify a node
	OutC     chan interface{} // OutC is a channel, which should be passed to the Handler and set by Handler to announce the Handler is done. Then nodes which after and depends on this can be unblocked.
	Value    interface{}      //Value is used to store data of node.

	Handler NodeHander //Handler is implemented to do some custom computation
}

// SetNodePreName set the PreName for node when call NewNode to construct
// a node. To expose a function to do thar means the field
// is optional for initializing a node.
func SetNodePreName(pre string) func(nd *Node) {
	return func(nd *Node) {
		nd.PreName = pre
	}
}

// SetNodeNextName set the NextName for node when call NewNode to construct
// a node. To expose a function to do thar means the field
// is optional for initializing a node.
func SetNodeNextName(next string) func(nd *Node) {
	return func(nd *Node) {
		nd.NextName = next
	}
}

// SetNodeValue set the value for node when call NewNode to construct a node.
// To expose a function to do that means the field
// is optional for initializing a node.
func SetNodeValue(value interface{}) func(nd *Node) {
	return func(nd *Node) {
		nd.Value = value
	}
}

// SetNodeHandler set handler for node when call NewNode to construct a node.
// To expose a function to do that means the field
// is optional for initializing a node. If do not set a custom handler
// NewNode will use a defaultHandler which just close the OutC to make sure
// node behind would not be blocked forever.
func SetNodeHandler(handler NodeHander) func(nd *Node) {
	return func(nd *Node) {
		nd.Handler = handler
	}
}

func defaultHandler(self *Node, in chan chan interface{}, out chan interface{}) {
	close(out)
}

// NewNode will construct a new node with the spcified name.
// And user also can call the API exposed to set other fileds which are not
// necessary.
func NewNode(name string, options ...func(nd *Node)) *Node {
	node := new(Node)
	node.Name = name
	node.OutC = make(chan interface{})
	node.Handler = defaultHandler

	for _, opt := range options {
		opt(node)
	}

	return node
}
