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

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil || f == nil {
		return
	}
	for i := 0; ; i++ {
		if !recursiveBFS(root, i, f) {
			return
		}
	}
}

func recursiveBFS(root *TreeNode, index int, f func(...interface{}) (int, error)) bool {
	if root == nil {
		return false
	}
	if index == 0 {
		f(root.Data)
		return true
	} else {
		tmp1, tmp2 := recursiveBFS(root.Left, index-1, f), recursiveBFS(root.Right, index-1, f)
		return tmp1 || tmp2
	}
}
