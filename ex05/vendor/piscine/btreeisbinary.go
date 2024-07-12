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

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	tmp1, tmp2 := BTreeIsBinary(root.Left), BTreeIsBinary(root.Right)
	if tmp1 && tmp2 {
		if (root.Left == nil || root.Left.Data <= root.Data) &&
			(root.Right == nil || root.Right.Data >= root.Data) {
			return true
		}
	}
	return false
}
