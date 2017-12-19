package fantree

import (
	"fmt"
	"sync"
)

//TreeNode is the node Of tree.
type TreeNode struct {
	Name     string      //Name of TreeNode should equal to the node
	Value    *Node       //Value of TreeNode is a pointer to the Node
	Previous []*TreeNode // A TreeNode may have many previous
	Next     []*TreeNode // A TreeNode may have many next too.
}

//NewTreeNode construct a TreeNode based on 1 node pointer
func NewTreeNode(nd *Node) *TreeNode {
	treeNode := new(TreeNode)
	treeNode.Name = nd.Name
	treeNode.Value = nd
	treeNode.Previous = []*TreeNode{}
	treeNode.Next = []*TreeNode{}

	return treeNode
}

func findTreeNodeInList(node *TreeNode, treeNodeList []*TreeNode) bool {
	for _, n := range treeNodeList {
		if node == n {
			return true
		}
	}
	return false
}

//Forest is structure to store a forest
type Forest struct {
	NodePool map[string]*TreeNode //NodePool store all TreeNodes of forest in a map
	Roots    []*TreeNode          //Roots is slice pointers to mark the roots in forest.
}

//NewForest construct a new forest based on a list of Node
func NewForest(nodes []*Node) (forest *Forest, err error) {
	forest = new(Forest)
	forest.setupForestNodePool(nodes)
	err = forest.setupForest()

	return forest, err
}

func (frt *Forest) setupForestNodePool(nodeList []*Node) {
	frt.NodePool = make(map[string]*TreeNode)

	//store all node in a map
	for _, node := range nodeList {
		frt.NodePool[node.Name] = NewTreeNode(node)
	}
}

func (frt *Forest) setupForest() error {
	//setup the tree
	for _, treeNode := range frt.NodePool {
		pre := treeNode.Value.PreName
		next := treeNode.Value.NextName

		if preNode, ok := frt.NodePool[pre]; ok {
			if !findTreeNodeInList(preNode, treeNode.Previous) {
				treeNode.Previous = append(treeNode.Previous, preNode)
			}
			if !findTreeNodeInList(treeNode, preNode.Next) {
				preNode.Next = append(preNode.Next, treeNode)
			}
		}
		if nextNode, ok := frt.NodePool[next]; ok {
			if !findTreeNodeInList(nextNode, treeNode.Next) {
				treeNode.Next = append(treeNode.Next, nextNode)
			}
			if !findTreeNodeInList(treeNode, nextNode.Previous) {
				nextNode.Previous = append(nextNode.Previous, treeNode)
			}
		}
	}

	return frt.findRoots()
}

func (frt *Forest) findRoots() error {
	hCount := 0
	for _, node := range frt.NodePool {
		if len(node.Previous) == 0 {
			hCount++
			frt.Roots = append(frt.Roots, node)
		}
	}

	if hCount == 0 {
		return fmt.Errorf("Do not find a root!")
	} else {
		return nil
	}
}

func printForest(roots []*TreeNode) {
	for _, root := range roots {
		fmt.Println("Name:", root.Name)
		if len(root.Next) > 0 {
			printForest(root.Next)
		}
	}
}

//Print will print the node name of forest
func (frt *Forest) Print() {
	if len(frt.Roots) == 0 {
		fmt.Println("There is still no tree in forest")
		return
	}
	printForest(frt.Roots)
}

//GetTrees return all roots in the forest
func (frt *Forest) GetTrees() []*TreeNode {
	return frt.Roots
}

//GetTree return the specified tree(through the name of root node) in the forest.
// If there isn't a tree with the name in the forest an error will be returned.
func (frt *Forest) GetTree(name string) (root *TreeNode, err error) {
	for _, root = range frt.Roots {
		if root.Name == name {
			return root, nil
		}
	}

	return nil, fmt.Errorf("Do not find the tree with a root which name is:[%s]", name)
}

// Pipeline launch goroutines for every nodes in the forest to execute the node's handler concurrently.
// It will Set up a wait group to sync the execution of all nodes.
func (frt *Forest) Pipeline() error {
	var wg sync.WaitGroup

	for _, tp := range frt.NodePool {
		node := tp.Value
		wg.Add(1)

		//fan-in the previous channels for the nodes
		inC := make(chan chan interface{}, len(tp.Previous))

		if len(tp.Previous) > 0 {
			for _, pre := range tp.Previous {
				inC <- pre.Value.OutC
			}
		}
		close(inC)

		//start the goroutine to execute node handler
		go func() {
			node.Handler(node, inC, node.OutC)
			wg.Done()
		}()

	}
	//wait all goroutine done
	// TODO: right now just implement the way to sync data and semaphore between goroutines, need a way to
	// notify the downstream goroutines cancel when errors happened on upstream.
	wg.Wait()
	return nil
}
