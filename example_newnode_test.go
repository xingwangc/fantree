// Copyright 2017 Simon Cai <xingwangc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fantree_test

import (
	"fmt"
	"sync"

	"github.com/xingwangc/fantree"
)

func handler(self *fantree.Node, inC chan chan interface{}, outC chan interface{}) {
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
		fmt.Println("hello")
		close(outC)
	}()
}

func Example_newNode() {
	inC := make(chan chan interface{})
	outC := make(chan interface{})

	node := fantree.NewNode("test_node",
		fantree.SetNodePreName("previous"),
		fantree.SetNodeNextName("next"),
		fantree.SetNodeCommand("test", handler),
	)

	go node.Cmd.Handler(node, inC, outC)

	close(inC)
	<-outC
	// Output: hello
}
