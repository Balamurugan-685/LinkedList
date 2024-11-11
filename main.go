package main

import (
	list "linkedList/src"
	"log"
)

func main() {
	node := &list.Node{Value: 3, Next: &list.Node{Value: 5, Next: &list.Node{Value: 8, Next: nil}}}
	input := &list.LinkedList{Node: node}
	result := list.ReverseLinkedList(input)
	for result.Node != nil {
		log.Println(result.Node.Value)
		result.Node = result.Node.Next
	}
}
