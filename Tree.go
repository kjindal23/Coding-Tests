package main

import "fmt"

type node struct{
	left *node
	data int
	right *node
}

type tree struct{
	root *node
}

func (n *node) insert(d int){

	if (n == nil){
		return
	}else if(d <= n.data){
		if(n.left == nil){
			n.left = &node{left:nil,data:d,right:nil}
		}else{
			n.left.insert(d)
		}
	}else {
		if(n.right == nil){
			n.right = &node{left:nil,data:d,right:nil}
		}else {
			n.right.insert(d)
		}
	}
	}

func (t *tree) insert(d int){

	if(t.root == nil){
		t.root = &node{left:nil,data:d,right:nil}
	}else{
		t.root.insert(d)
	}
}
func (n *node)inorder(){
	
	if(n.left != nil){
		n.left.inorder()
	}
	fmt.Println("IN ORDER " ,n.data)
	if(n.right != nil){
		n.right.inorder()
	}
}
func (n *node)preorder(){
	
	fmt.Println("PRE ORDER" ,n.data)
	if(n.left != nil){
		n.left.preorder()
	}
	
	if(n.right != nil){
		n.right.preorder()
	}
}
func (n *node)postorder(){
	
	if(n.left != nil){
		n.left.postorder()
	}

	if(n.right != nil){
		n.right.postorder()
	}
	fmt.Println("POSTORDER" , n.data)
}
func (t *tree) print(){
	if(t.root == nil){
		fmt.Println("Empty tree")
	}else {
		t.root.inorder()
		t.root.preorder()
		t.root.postorder()
	}
}

func main(){
	var t1 tree
	t1.insert(20)
	t1.insert(15)
	t1.insert(25)
	t1.insert(10)
	t1.insert(21)
	
	t1.insert(12)
	t1.insert(13)
	t1.insert(11)
	t1.insert(40)
	t1.insert(50)
	t1.print()
}