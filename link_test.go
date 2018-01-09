// Copyright 2017 Simon Cai <xingwangc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fantree_test

import (
	"sync"
	"testing"

	"github.com/xingwangc/fantree"
)

func link_handler_1(node *fantree.Node, inC chan chan interface{}, outC chan interface{}) {
	var wg sync.WaitGroup

	go func() {
		for in := range inC {
			ch := in
			wg.Add(1)
			go func(ch chan interface{}) {
				meta := <-ch
				if value, ok := meta.(int); ok {
					outC <- value + 1
				}
				wg.Done()
			}(ch)
		}
		wg.Wait()
		close(outC)
	}()
}

func link_handler_2(node *fantree.Node, inC chan chan interface{}, outC chan interface{}) {
	var wg sync.WaitGroup

	go func() {
		for in := range inC {
			ch := in
			wg.Add(1)
			go func(ch chan interface{}) {
				meta := <-ch
				if value, ok := meta.(int); ok {
					outC <- value + 2
				}
				wg.Done()
			}(ch)
		}
		wg.Wait()
		close(outC)
	}()
}
func link_handler_3(node *fantree.Node, inC chan chan interface{}, outC chan interface{}) {
	var wg sync.WaitGroup

	go func() {
		for in := range inC {
			ch := in
			wg.Add(1)
			go func(ch chan interface{}) {
				meta := <-ch
				if value, ok := meta.(int); ok {
					outC <- value + 3
				}
				wg.Done()
			}(ch)
		}
		wg.Wait()
		close(outC)
	}()
}
func link_handler_4(node *fantree.Node, inC chan chan interface{}, outC chan interface{}) {
	var wg sync.WaitGroup

	go func() {
		for in := range inC {
			ch := in
			wg.Add(1)
			go func(ch chan interface{}) {
				meta := <-ch
				if value, ok := meta.(int); ok {
					outC <- value + 4
				}
				wg.Done()
			}(ch)
		}
		wg.Wait()
		close(outC)
	}()
}
func link_handler_5(node *fantree.Node, inC chan chan interface{}, outC chan interface{}) {
	var wg sync.WaitGroup

	go func() {
		for in := range inC {
			ch := in
			wg.Add(1)
			go func(ch chan interface{}) {
				meta := <-ch
				if value, ok := meta.(int); ok {
					outC <- value + 5
				}
				wg.Done()
			}(ch)
		}
		wg.Wait()
		close(outC)
	}()
}
func link_handler_6(node *fantree.Node, inC chan chan interface{}, outC chan interface{}) {
	var wg sync.WaitGroup

	go func() {
		for in := range inC {
			ch := in
			wg.Add(1)
			go func(ch chan interface{}) {
				meta := <-ch
				if value, ok := meta.(int); ok {
					outC <- value + 6
				}
				wg.Done()
			}(ch)
		}
		wg.Wait()
		close(outC)
	}()
}
func TestLinkPipeline(t *testing.T) {
	nodeList := []*fantree.Node{}
	nodeList = append(nodeList, fantree.NewNode("一",
		fantree.SetNodeCommand("一", link_handler_1)))
	nodeList = append(nodeList, fantree.NewNode("二",
		fantree.SetNodePreName("一"),
		fantree.SetNodeCommand("二", link_handler_2)))
	nodeList = append(nodeList, fantree.NewNode("三",
		fantree.SetNodePreName("一"),
		fantree.SetNodeCommand("三", link_handler_3)))
	nodeList = append(nodeList, fantree.NewNode("四",
		fantree.SetNodeNextName("六"),
		fantree.SetNodeCommand("四", link_handler_4)))
	nodeList = append(nodeList, fantree.NewNode("五",
		fantree.SetNodeNextName("六"),
		fantree.SetNodeCommand("五", link_handler_5)))
	nodeList = append(nodeList, fantree.NewNode("六",
		fantree.SetNodePreName("四"),
		fantree.SetNodeCommand("六", link_handler_6)))

	link, err := fantree.NewLink(nodeList)
	if err != nil {
		t.Error(err)
	}
	output, err := link.Pipeline(10)
	if err != nil {
		t.Error(err)
	}

	if o, ok := output.(int); ok {
		if o != 31 {
			t.Error("Result should be 31")
		}
	} else {
		t.Error(output)
	}
}
