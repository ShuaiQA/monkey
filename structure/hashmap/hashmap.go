package hashmap

// 实现一个可扩展的hash
type Pair struct {
	key int
	val int
}

type Bucket struct {
	dep   int
	pairs []Pair
}

// 根据传入的key设置value的值
func (b *Bucket) find(key int) (int, bool) {
	for _, v := range b.pairs {
		if v.key == key {
			return v.key, true
		}
	}
	return 0, false
}

func (b *Bucket) remove(key int) bool {
	for i := range b.pairs {
		if b.pairs[i].key == key {
			b.pairs = append(b.pairs[:i], b.pairs[:i+1]...)
			return true
		}
	}
	return false
}

func (b *Bucket) insert(key, val int) {
	// 找到进行修改
	for i := range b.pairs {
		if b.pairs[i].key == key {
			b.pairs[i].val = val
			return
		}
	}
	// 没有找到进行添加
	b.pairs = append(b.pairs, Pair{key: key, val: val})
}

type ExtendHashTable struct {
	dep         int
	bucket_size int
	num_bucket  int
	dir         []*Bucket
}

func NewExtendHashTable(bucket_size int) *ExtendHashTable {
	return &ExtendHashTable{dep: 0, bucket_size: bucket_size, num_bucket: 1, dir: make([]*Bucket, 1)}
}

// 选择hash函数的低dep位当作dir的下标
func (e *ExtendHashTable) hash(key int) int {
	mask := (1 << e.dep) - 1
	hash := 10
	return hash & mask
}
