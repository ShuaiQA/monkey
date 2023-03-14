package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	Root *TreeNode
}

func New() *Tree {
	return new(Tree)
}

func (t *Tree) Find(val int) bool {
	if t.Root.Val == val {
		return true
	}
	_, target := t.find(val)
	return target != nil && target.Val == val
}

func (t *Tree) Insert(val int) {
	pa := t.Root
	if pa == nil {
		t.Root = &TreeNode{Val: val}
		return
	}
	parent, target := t.find(val)
	if target == nil { // 没有找到对应的值
		if val > parent.Val {
			parent.Right = &TreeNode{Val: val}
		} else {
			parent.Left = &TreeNode{Val: val}
		}
	}
}

func (t *Tree) Delete(val int) bool {
	pp, p := t.find(val)
	if p == nil {
		return false
	}
	if p.Left != nil && p.Right != nil {
		minpp, minp := t.findRightMin(p)
		p.Val = minp.Val
		minpp.Left = minp.Right
		return true
	}
	var child *TreeNode
	if p.Left != nil {
		child = p.Left
	} else {
		child = p.Right
	}
	if pp == nil { // 删除的是根节点的val
		t.Root = child
	} else if pp.Left == p {
		pp.Left = child
	} else {
		pp.Right = child
	}
	return true
}

// 根据当前的val的值,进行向下查找,找到的val的节点放在target中,没有找到val,target节点是空的
// parent 记录着向下查找的过程也就是target的父节点
// 如果查找的节点是根节点,其parent设置为空
func (t *Tree) find(val int) (parent, target *TreeNode) {
	parent = nil
	target = t.Root
	for target != nil {
		if target.Val > val {
			parent = target
			target = target.Left
		} else if target.Val < val {
			parent = target
			target = target.Right
		} else {
			break
		}
	}
	return parent, target
}

// 查找当前节点的右面最小的节点,返回右面最小的节点以及父节点
func (t *Tree) findRightMin(node *TreeNode) (parent, target *TreeNode) {
	parent = node
	target = node.Right
	for target.Left != nil {
		parent = target
		target = target.Left
	}
	return
}
