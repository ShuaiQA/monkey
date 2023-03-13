package heap

// 设计分析首先优先队列需要使用数组进行实现的,完全二叉树
// 对于优先队列来说需要实现初始化操作,给定一个数组进行初始化
// 该接口需要实现比较操作

type Interface interface {
	Len() int
	Less(i, j int) bool // arr[i] < arr[j] ?
	Swap(i, j int)
	Push(x any)
	Pop() any
}

func Init(i Interface) {
	n := i.Len()
	// 从非叶子节点开始进行下沉操作
	for k := n/2 - 1; k >= 0; k-- {
		down(i, k, n)
	}
}

// 对于添加的数值x,添加到数组的最后一位的下标,需要对添加的值进行上浮操作
func Push(i Interface, x any) {
	i.Push(x)
	up(i, i.Len()-1)
}

// 删除当前的数组下标i的优先队列,主要的操作是,将i和n进行交换
// 然后需要下移操作
func Remove(h Interface, i int) any {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		down(h, i, n)
	}
	return h.Pop()
}

// 获取数组最后的下标值,然后和0下标进行交换,下移0下标操作(注意此处的n的范围)
// 注意此处并没有删除最后一个元素,只是down的范围没有最后一个元素
// 然后调用h.Pop(),一般h.Pop()是获取数组的最后一位元素(),然后数组缩小1
func Pop(h Interface) any {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

// 将位置是k的地方的值,进行比较操作,然后确定是否移动操作
func down(i Interface, k, n int) {
	// 如果当前的k没有左右子结点或者是当前的k<0,不需要down
	if 2*k+1 >= n || k < 0 {
		return
	}
	j0 := k*2 + 1
	if j1 := j0 + 1; j1 < n && i.Less(j1, j0) { // 当j1小于j0的时候执行下面的
		j0 = j1
	}
	if i.Less(j0, k) { // 当选取的j0小于k的时候调换
		i.Swap(k, j0)
		down(i, j0, n)
	}
}

func up(i Interface, j int) {
	if j == 0 {
		return
	}
	f := (j - 1) / 2
	if i.Less(j, f) { // 当前的节点比父节点还要小,需要交换
		i.Swap(j, f)
		up(i, f)
	}
}
