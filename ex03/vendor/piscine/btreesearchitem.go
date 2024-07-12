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
