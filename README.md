# fantree
fantree is a package which provide the implementation to concurrently execute user defined event.

## Install

	go get github.com/xingwangc/fantree

## Doc

Please refer to [GoDoc](https://godoc.org/github.com/xingwangc/fantree) for more information.

## Overview

When we programming to process data, execute computation, or schedule tasks. there are kind of cases, like using commands to process data stream, parts of commands may depends on others, others are not. And sometimes, even command depends on others, it is not necessary to immediately following them.I called the relationship of these commands, tasks as weak coupling here. And I use the style which is used to defined the service relations in systemd to define the relationship between commands(or tasks) here. For example, if there is a command `B` depends on `A`, and a `C` depends on `B`, which means `C` should execute after `B`, and `B` is After `A`。 So use a table to list the relations is as bellow :

| command | before | after |
|---------|--------|-------|
| A       | B      |       |
| B       | C      | A     |
| C       |        | B     |

* before: means *execute before*, so in a full sentence here should be "A execute before B". If we use a C style double linked list to express this, it should be A->next = B, B->previous = A
* after: means *execute after*, full exprecession here should be "B execute after A".

And then go on our example, if comes another command `D`, it does not rely on previous 3 commands. So we can get the conclusion, that A, B and C should be executed synchronously, but D can be executed concurrently with them.

Then how about a new command `E` is defined to just depends on `A`, and the command `C` is changed to depend on both `B` and a new command `F`? This is what fantree designed to handle.

## fantree

The package *fantree* is designed to provide a middle layer to constract a special forest for commands defined like we discussed previously. When users defining the commands, they also can define a handler for every commands depends on the function they want. If there isn't a handler provided for command(task), package provides a default one, but it do nothing except closing the out channel as signal.

Then user can invoke the `Pipeline` method of forest to execute pipeline. Which will do the pipeline as user defined automatically. Interdependent commands will be executed synchronously, irrelevant commands will be executed concurrently.

*fantree* use the **fan-out**,**fan-in** mode, and it setup a forest based on a directed graph, so it is named *fantree*. Besides the **fan-out**, **fan-in**, it also used the **future/promise** mode.

### Construct the forest

Let's continue our example. Let's list all commands in a table again.

| command | before | after |
|---------|--------|-------|
| A       | B      |       |
| B       | C      | A     |
| C       |        | B     |
| D       |        |       |
| E       |        | A     |
| F       | C      |       |

it should be constructed as a special forest(actually it is a set of directed graphs) like:

	    E
	  /
	A
	  \
	    B
	      \
	        C
	      /
	    F
	D

From the figure, it is so easy to understand that: A, F and D can be executed concurrently; E and B can concurrently execute after A is complete; at last C will execute until both B and F is complete.

### fan-out

The result of A will be fan-out to E and B.

### fan-in

The results of B and F will be fan-in to C.

### future/promise

All commands start by seperate goroutines to concurrently executing with a promise. But some will be blocked until the future comes.

### single link

If all your commands will process a single copy of concurrency unsafely data like map. But they also have the weak coupling feature in line with the description, you also can use `NewLink()` function to help you organize them in a single link, and then use its Pipeline to execute command one by one.

## Example

Detail examples please refer to the examples and tests in package.

## Copyright

Copyright 2017 Simon Cai <xingwangc@gmail.com>. All rights reserved.
 Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.
