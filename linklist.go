package main

import "fmt"

type Node struct{
	next *Node
	key interface{}
}

type List struct{
	head *Node
}

func (L *List)insert(key interface{}){
	tempnode := &Node{
		key : key,
		next : nil,
	}
	if L.head == nil{
		L.head = tempnode;
	}else{
		l := L.head
		for l.next != nil{
			l = l.next
			}
			l.next = tempnode;

		}
	}

func (L *List )display () {
	head := L.head;

	for head != nil {
		fmt.Printf("%+v" , head.key)
		head = head.next;
		fmt.Println()
	}
}

func main(){

	list := List{}

	list.insert(5)
	list.insert(6)
	list.insert(7)
	list.insert(8)
	list.insert(9)

	list.display()
	
}