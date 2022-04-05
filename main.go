package main

import (
	"fmt"
	"github.com/selfcheer/leetcode/linked_list"
)

func main() {
	myLL := linked_list.Constructor()
	//myLL.AddAtHead(1)
	myLL.AddAtTail(3)
	//myLL.AddAtIndex(1, 2)
	//fmt.Println(myLL.Get(1))
	//myLL.DeleteAtIndex(0)
	fmt.Println(myLL.Get(0))
}