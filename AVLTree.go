package main
import "fmt"



type tree struct{
	left node
	right node
	root node

}

type node struct{
	val int
}
//TODO: add in pointers to increase efficiency for leftBalance and rightBalancefunc

func rightBalance(AVLTree tree, addNode node) (t tree) {
	if addNode.val > AVLTree.right.val{
		temp := AVLTree.right
		AVLTree.right = addNode
		return rightBalance(AVLTree,temp)
	}else{
		return leftBalance(AVLTree,addNode)
	}
	return AVLTree
}
func leftBalance(AVLTree tree, addNode node) (t tree) {
	if addNode.val < AVLTree.left.val{
		temp := AVLTree.left
		AVLTree.left = addNode
		return leftBalance(AVLTree,temp)
	}else{
		return rightBalance(AVLTree, addNode)
	}

	return AVLTree
}

func insert(AVLTree tree, addNode node) (t tree) {
	if AVLTree.root.val > addNode.val{
		leftBalance(AVLTree,addNode)
	}else{
		rightBalance(AVLTree, addNode)
	}
	return AVLTree


}

func main(){
	AVLTree := tree{root: node{val:5},left:node{val:3},right:node{val: 2}}
	insert(AVLTree, node{val:6})
	fmt.Println(AVLTree)

}
