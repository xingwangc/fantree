// Copyright 2017 Simon Cai <xingwangc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fantree_test

import (
	"github.com/xingwangc/fantree"
)

func ExampleNewLink() {

	nodeList := []*fantree.Node{}
	nodeList = append(nodeList, fantree.NewNode("1"))
	nodeList = append(nodeList, fantree.NewNode("2",
		fantree.SetNodePreName("1")))
	nodeList = append(nodeList, fantree.NewNode("3",
		fantree.SetNodePreName("1")))
	nodeList = append(nodeList, fantree.NewNode("4",
		fantree.SetNodePreName("1")))
	nodeList = append(nodeList, fantree.NewNode("5",
		fantree.SetNodePreName("1")))
	nodeList = append(nodeList, fantree.NewNode("6",
		fantree.SetNodePreName("2")))
	nodeList = append(nodeList, fantree.NewNode("7",
		fantree.SetNodePreName("2")))
	nodeList = append(nodeList, fantree.NewNode("8",
		fantree.SetNodePreName("3")))
	nodeList = append(nodeList, fantree.NewNode("9",
		fantree.SetNodePreName("3")))
	nodeList = append(nodeList, fantree.NewNode("10",
		fantree.SetNodePreName("5")))
	nodeList = append(nodeList, fantree.NewNode("11",
		fantree.SetNodePreName("5")))
	nodeList = append(nodeList, fantree.NewNode("12",
		fantree.SetNodePreName("7")))
	nodeList = append(nodeList, fantree.NewNode("13",
		fantree.SetNodePreName("9")))
	nodeList = append(nodeList, fantree.NewNode("14",
		fantree.SetNodePreName("12")))
	nodeList = append(nodeList, fantree.NewNode("15",
		fantree.SetNodePreName("12")))
	nodeList = append(nodeList, fantree.NewNode("16",
		fantree.SetNodePreName("13")))
	nodeList = append(nodeList, fantree.NewNode("17",
		fantree.SetNodePreName("10")))
	nodeList = append(nodeList, fantree.NewNode("18",
		fantree.SetNodePreName("14")))
	nodeList = append(nodeList, fantree.NewNode("19",
		fantree.SetNodePreName("16"),
		fantree.SetNodeNextName("25")))
	nodeList = append(nodeList, fantree.NewNode("20",
		fantree.SetNodePreName("17"),
		fantree.SetNodeNextName("25")))
	nodeList = append(nodeList, fantree.NewNode("21"))
	nodeList = append(nodeList, fantree.NewNode("22",
		fantree.SetNodePreName("21")))
	nodeList = append(nodeList, fantree.NewNode("23",
		fantree.SetNodePreName("22")))
	nodeList = append(nodeList, fantree.NewNode("24",
		fantree.SetNodePreName("22")))
	nodeList = append(nodeList, fantree.NewNode("25"))

	link, _ := fantree.NewLink(nodeList)

	// Output should unordered like: 21 -> 1 -> 22 -> 3 -> 23 -> 4 -> 24 -> 5 -> 2 -> 10 -> 9 -> 11 -> 6 -> 17 -> 8 -> 20 -> 7 -> 13 -> 12 -> 16 -> 14 -> 19 -> 15 -> 25 -> 18 ->
	link.Print()
}
