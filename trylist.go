package main

import (
	"container/list"
	"fmt"
)

func main() {
	intlist := list.New()
	intlist.PushBack(11)
	intlist.PushBack(22)
	intlist.PushBack(33)

	for e := intlist.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}
