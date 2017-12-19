// Copyright 2017 Simon Cai <xingwangc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fantree_test

import (
	"github.com/xingwangc/fantree"
)

func Example_setupForest_1() {
	nodeList := []*fantree.Node{}
	nodeList = append(nodeList, fantree.NewNode("一"))
	nodeList = append(nodeList, fantree.NewNode("二",
		fantree.SetNodePreName("一")))
	nodeList = append(nodeList, fantree.NewNode("三",
		fantree.SetNodePreName("一")))
	nodeList = append(nodeList, fantree.NewNode("四",
		fantree.SetNodeNextName("六")))
	nodeList = append(nodeList, fantree.NewNode("五",
		fantree.SetNodeNextName("六")))
	nodeList = append(nodeList, fantree.NewNode("六",
		fantree.SetNodePreName("四")))

	// forest should like:
	//	     二
	//	   /
	//	一
	//         \
	//	     三
	//	四
	//	  \
	//	    六
	//	  /
	//	五
	forest, _ := fantree.NewForest(nodeList)

	forest.Print()
	// Unordered output:
	// Name: 一
	// Name: 三
	// Name: 二
	// Name: 五
	// Name: 六
	// Name: 四
	// Name: 六
}

func Example_setupForest_2() {

	nodeList := []*fantree.Node{}
	nodeList = append(nodeList, fantree.NewNode("一"))
	nodeList = append(nodeList, fantree.NewNode("二",
		fantree.SetNodePreName("一")))
	nodeList = append(nodeList, fantree.NewNode("三",
		fantree.SetNodePreName("一")))
	nodeList = append(nodeList, fantree.NewNode("四",
		fantree.SetNodePreName("二"),
		fantree.SetNodeNextName("六")))
	nodeList = append(nodeList, fantree.NewNode("五",
		fantree.SetNodePreName("三"),
		fantree.SetNodeNextName("六")))
	nodeList = append(nodeList, fantree.NewNode("六",
		fantree.SetNodePreName("四")))

	// forest should like:
	//	     二 - 四
	//	   /         \
	//	一             六
	//         \        /
	//	     三 - 五
	forest, _ := fantree.NewForest(nodeList)

	forest.Print()
	// Unordered output:
	// Name: 一
	// Name: 二
	// Name: 四
	// Name: 六
	// Name: 三
	// Name: 五
	// Name: 六
}
