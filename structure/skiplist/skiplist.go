package skiplist

import (
	"fmt"
	"math/rand"
)

const maxLevel = 10
const pFactor = 0.5

// 初始化当前的root节点,标注当前跳表的最大高度
type Skiplist struct {
	root *Node
}

// 当前节点包含val(存储的值)
// 随机化当前的node节点有多少个next的指针
type Node struct {
	val  int
	next []*Node
}

// func (s *Skiplist) String() string {
// 	ss := ""
// 	for i := maxLevel - 1; i >= 0; i-- {
// 		cur := s.root
// 		for cur != nil {
// 			ss = fmt.Sprintf("%s  %d", ss, cur.val)
// 			cur = cur.next[i]
// 		}
// 		ss += "\n"
// 	}
// 	return ss
// }

func (s *Skiplist) randomLevel() int {
	lv := 1
	for lv < maxLevel && rand.Float64() < pFactor {
		lv++
	}
	fmt.Println(lv)
	return lv
}

// 构建一个Skiplist节点,设置跳表的最大的层数
// 构建一个root节点,该节点含有一个无效的-1值,并且含有maxLevel个next节点,现在next节点全部是nil
func Constructor() *Skiplist {
	r := make([]*Node, maxLevel)
	return &Skiplist{root: &Node{val: -1, next: r}}
}

// 获取最接近num的节点的前一个节点,有助于查找、添加、删除操作
// 为了防止是nil需要初始化为root的节点,也就是说无论是查找、添加、删除都是在root之后的节点
func (s *Skiplist) getPreNode(num int) []*Node {
	pre := make([]*Node, maxLevel)
	for i := range pre {
		pre[i] = s.root
	}
	cur := s.root
	for i := maxLevel - 1; i >= 0; i-- {
		// 如果当前层的下一个节点不为空,并且下一个节点值小于当前查找数
		// (若下一个节点等于当前数,那么下一个节点就是target)向后移动将记录不了pre了
		for cur.next[i] != nil && cur.next[i].val < num {
			cur = cur.next[i]
		}
		pre[i] = cur // pre获取的是当前的节点,如果配合查找、添加、删除还需要结合下标操作
	}
	return pre
}

func (s *Skiplist) Search(target int) bool {
	pre := s.getPreNode(target)
	node := pre[0].next[0]
	return node != nil && target == node.val
}

func (s *Skiplist) Add(num int) {
	pre := s.getPreNode(num)
	lv := s.randomLevel()
	add := &Node{val: num, next: make([]*Node, lv)}
	for i, v := range pre[:lv] {
		add.next[i] = v.next[i]
		v.next[i] = add
	}
}

func (s *Skiplist) Erase(num int) bool {
	pre := s.getPreNode(num)
	cur := pre[0].next[0]
	if cur == nil || num != cur.val { // 没有找到num
		return false
	}
	// 先查找当前的num的节点含有多少个next的指针,然后对pre的对应的指针设置num节点的next的指针的值
	for i := 0; i < len(cur.next); i++ {
		// 修改前面的节点的next层的节点指向后续的next
		pre[i].next[i] = cur.next[i]
	}
	return true
}
