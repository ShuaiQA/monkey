package tree

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

type Tree struct {
	Root *TreeNode
}

func New() *Tree {
	return new(Tree)
}

// 返回节点的值是val的节点
func (t *Tree) Find(val int) *TreeNode {
	c := t.Root
	for c != nil && c.Val != val {
		if val < c.Val {
			c = c.Left
		} else {
			c = c.Right
		}
	}
	return c
}

// 并没有解决树中可能重复的值
func (t *Tree) Insert(z *TreeNode) {
	// y代表的是插入节点z的父节点,x代表的是插入的节点,如果树中没有节点z相同的值则x for循环之后==nil
	var y *TreeNode = nil
	x := t.Root
	for x != nil {
		y = x
		if z.Val < x.Val {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	// 设置z的父节点就是y
	z.Parent = y
	if y == nil { // 当前的树是空的,设置树的根节点
		t.Root = z
	} else if z.Val < y.Val { // 决定当前的z节点是y节点的左右孩子
		y.Left = z
	} else {
		y.Right = z
	}
}

// 删除分为4种情况
// 1、左节点是空的使用右节点进行替换(右节点可以是空的)
// 2、右节点是空的使用左节点进行替换
// 左右节点都不是空的,使用minnum函数找到后继节点(后继节点的性质决定了没有左孩子)
// 3、若后继节点的父节点不是删除的节点,需要将后继节点的右孩子替换掉后继节点,然后执行4
// 4、后继节点y的父节点z是删除的节点,直接将y节点替换掉z节点,然后设置y节点的左孩子节点和左孩子节点的父节点
func (t *Tree) Delete(z *TreeNode) {
	if z.Left == nil {
		t.transplant(z, z.Right)
	} else if z.Right == nil {
		t.transplant(z, z.Left)
	} else {
		y := minnum(z.Right)
		if y.Parent != z {
			t.transplant(y, y.Right)
			// 设置y后继节点的孩子节点,以及孩子节点的父节点
			// 因为在transplant并没有考虑孩子节点
			y.Right = z.Right
			y.Right.Parent = y
		}
		t.transplant(z, y)
		y.Left = z.Left
		y.Left.Parent = y
	}
}

// 将以v为根的树,来替换以u为根的树,注意以u为根的树的节点并没有修改
// 将原本u的父节点的孩子是u,使用v来代替,并设置v的父节点是u的父节点
// 只是设置了v的父节点和父节点的孩子是v,没有设置v的孩子是什么,也没有说被替换之后的u应该做什么
func (t *Tree) transplant(u, v *TreeNode) {
	if u.Parent == nil {
		t.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	// 当前的节点v可能是空的,例如没有左孩子使用右孩子来代替,但是右孩子可能是空的
	if v != nil {
		v.Parent = u.Parent
	}
}

func minnum(x *TreeNode) *TreeNode {
	c := x
	for c.Left != nil {
		c = c.Left
	}
	return c
}
