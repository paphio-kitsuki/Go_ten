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

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil || f == nil {
		return
	}
	recursiveApply(root, f)
}

func recursiveApply(node *TreeNode, f func(...interface{}) (int, error)) {
	if node == nil {
		return
	}
	f(node.Data)
	recursiveApply(node.Left, f)
	recursiveApply(node.Right, f)
}
