// Copyright 2017 Simon Cai <xingwangc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fantree_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/xingwangc/fantree"
)

func TestSetupForest_3(t *testing.T) {

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

	forest, err := fantree.NewForest(nodeList)
	if err != nil {
		t.Error(err)
	}

	forest.Print()
}

func handler_1(node *fantree.Node, inC chan chan interface{}, outC chan interface{}) {
	var wg sync.WaitGroup

	go func() {
		for in := range inC {
			ch := in
			wg.Add(1)
			go func(ch chan interface{}) {
				<-ch
				wg.Done()
			}(ch)
		}
		wg.Wait()
		fmt.Println("hello 1!")
		close(outC)
	}()
}

func handler_2(node *fantree.Node, inC chan chan interface{}, outC chan interface{}) {
	var wg sync.WaitGroup

	go func() {
		for in := range inC {
			ch := in
			wg.Add(1)
			go func(ch chan interface{}) {
				<-ch
				wg.Done()
			}(ch)
		}
		wg.Wait()
		fmt.Println("hello 2!")
		close(outC)
	}()
}
func handler_3(node *fantree.Node, inC chan chan interface{}, outC chan interface{}) {
	var wg sync.WaitGroup

	go func() {
		for in := range inC {
			ch := in
			wg.Add(1)
			go func(ch chan interface{}) {
				<-ch
				wg.Done()
			}(ch)
		}
		wg.Wait()
		fmt.Println("hello 3!")
		close(outC)
	}()
}
func handler_4(node *fantree.Node, inC chan chan interface{}, outC chan interface{}) {
	var wg sync.WaitGroup

	go func() {
		for in := range inC {
			ch := in
			wg.Add(1)
			go func(ch chan interface{}) {
				<-ch
				wg.Done()
			}(ch)
		}
		wg.Wait()
		fmt.Println("hello 4!")
		close(outC)
	}()
}
func handler_5(node *fantree.Node, inC chan chan interface{}, outC chan interface{}) {
	var wg sync.WaitGroup

	go func() {
		for in := range inC {
			ch := in
			wg.Add(1)
			go func(ch chan interface{}) {
				<-ch
				wg.Done()
			}(ch)
		}
		wg.Wait()
		fmt.Println("hello 5!")
		close(outC)
	}()
}
func handler_6(node *fantree.Node, inC chan chan interface{}, outC chan interface{}) {
	var wg sync.WaitGroup

	go func() {
		for in := range inC {
			ch := in
			wg.Add(1)
			go func(ch chan interface{}) {
				<-ch
				wg.Done()
			}(ch)
		}
		wg.Wait()
		fmt.Println("hello 6!")
		close(outC)
	}()
}
func Test_ForestPipeline(t *testing.T) {
	nodeList := []*fantree.Node{}
	nodeList = append(nodeList, fantree.NewNode("一",
		fantree.SetNodeCommand("一", handler_1)))
	nodeList = append(nodeList, fantree.NewNode("二",
		fantree.SetNodePreName("一"),
		fantree.SetNodeCommand("二", handler_2)))
	nodeList = append(nodeList, fantree.NewNode("三",
		fantree.SetNodePreName("一"),
		fantree.SetNodeCommand("三", handler_3)))
	nodeList = append(nodeList, fantree.NewNode("四",
		fantree.SetNodeNextName("六"),
		fantree.SetNodeCommand("四", handler_4)))
	nodeList = append(nodeList, fantree.NewNode("五",
		fantree.SetNodeNextName("六"),
		fantree.SetNodeCommand("五", handler_5)))
	nodeList = append(nodeList, fantree.NewNode("六",
		fantree.SetNodePreName("四"),
		fantree.SetNodeCommand("六", handler_6)))

	forest, err := fantree.NewForest(nodeList)
	if err != nil {
		t.Error(err)
	}
	var metadata interface{}
	forest.Pipeline(metadata)
}
