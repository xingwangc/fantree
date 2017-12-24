// Copyright 2017 Simon Cai <xingwangc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fantree_test

import (
	"sync"
	"testing"

	"github.com/xingwangc/fantree"
)

func fb_handler_1(node *fantree.Node, inC chan chan interface{}, outC chan interface{}) {
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
		close(outC)
	}()
}

func fb_handler_2(node *fantree.Node, inC chan chan interface{}, outC chan interface{}) {
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
		close(outC)
	}()
}
func fb_handler_3(node *fantree.Node, inC chan chan interface{}, outC chan interface{}) {
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
		close(outC)
	}()
}
func fb_handler_4(node *fantree.Node, inC chan chan interface{}, outC chan interface{}) {
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
		close(outC)
	}()
}
func fb_handler_5(node *fantree.Node, inC chan chan interface{}, outC chan interface{}) {
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
		close(outC)
	}()
}
func fb_handler_6(node *fantree.Node, inC chan chan interface{}, outC chan interface{}) {
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
		close(outC)
	}()
}
func BenchmarkForestPipeline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nodeList := []*fantree.Node{}
		nodeList = append(nodeList, fantree.NewNode("一",
			fantree.SetNodeHandler(fb_handler_1)))
		nodeList = append(nodeList, fantree.NewNode("二",
			fantree.SetNodePreName("一"),
			fantree.SetNodeHandler(fb_handler_2)))
		nodeList = append(nodeList, fantree.NewNode("三",
			fantree.SetNodePreName("一"),
			fantree.SetNodeHandler(fb_handler_3)))
		nodeList = append(nodeList, fantree.NewNode("四",
			fantree.SetNodeNextName("六"),
			fantree.SetNodeHandler(fb_handler_4)))
		nodeList = append(nodeList, fantree.NewNode("五",
			fantree.SetNodeNextName("六"),
			fantree.SetNodeHandler(fb_handler_5)))
		nodeList = append(nodeList, fantree.NewNode("六",
			fantree.SetNodePreName("四"),
			fantree.SetNodeHandler(fb_handler_6)))

		forest, err := fantree.NewForest(nodeList)
		if err != nil {
			b.Error(err)
		}
		forest.Pipeline()
	}
}
