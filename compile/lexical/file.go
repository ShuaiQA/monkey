package lexical

type File struct {
	name  string // file name as provided to AddFile
	size  int    // file size as provided to AddFile
	lines []int  // 标记每一行的起始偏移量
}

// 将新的偏移加到lines中,确保当前新的偏移需要大于前一个偏移,小于file的size
func (f *File) AddLine(offset int) {
	if i := len(f.lines); (i == 0 || f.lines[i-1] < offset) && offset < f.size {
		f.lines = append(f.lines, offset)
	}
}
