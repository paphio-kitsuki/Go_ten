package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	tmp := &TreeNode{Data: data}
	if root == nil {
		return tmp
	}
	for now := root; ; {
		if data <= now.Data {
			if now.Left == nil {
				now.Left = tmp
				tmp.Parent = now
				break
			}
			now = now.Left
		} else {
			if now.Right == nil {
				now.Right = tmp
				tmp.Parent = now
				break
			}
			now = now.Right
		}
	}
	return root
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}
	if recursiveSearchNode(root, node) {
		if node.Left == nil && node.Right == nil {
			replaceChild(node, nil)
			if node == root {
				return nil
			}
		} else if node.Left == nil || node.Right == nil {
			var tmpNode *TreeNode
			if node.Left == nil {
				tmpNode, node.Right.Parent = node.Right, node.Parent
			} else {
				tmpNode, node.Left.Parent = node.Left, node.Parent
			}
			replaceChild(node, tmpNode)
			if node == root {
				return tmpNode
			}
		} else {
			tmpNode := node.Right
			for tmpNode.Left != nil {
				tmpNode = tmpNode.Left
			}
			replaceChild(tmpNode, nil)
			replaceChild(node, tmpNode)
			tmpNode.Left, tmpNode.Right = node.Left, node.Right
			if node == root {
				return tmpNode
			}
		}
	}
	return root
}

func replaceChild(node, src *TreeNode) {
	if node.Parent != nil {
		if node.Parent.Left == node {
			node.Parent.Left = src
		} else if node.Parent.Right == node {
			node.Parent.Right = src
		}
	}
}

func recursiveSearchNode(root, node *TreeNode) bool {
	if root == nil {
		return false
	}
	if root == node {
		return true
	}
	return recursiveSearchNode(root.Left, node) || recursiveSearchNode(root.Right, node)
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil || f == nil {
		return
	}
	recursiveApply(root, f)
}

func recursiveApply(node *TreeNode, f func(...interface{}) (int, error)) {
	if node == nil {
		return
	}
	recursiveApply(node.Left, f)
	f(node.Data)
	recursiveApply(node.Right, f)
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	return recursiveSearch(root, elem)
}

func recursiveSearch(node *TreeNode, elem string) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Data == elem {
		return node
	} else if node.Data > elem {
		return recursiveSearch(node.Left, elem)
	} else {
		return recursiveSearch(node.Right, elem)
	}
}
