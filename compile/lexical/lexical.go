package lexical

import (
	"fmt"
	"go/format"
	"regexp"
)

// ^字符串的开始
// (?:非捕获组
// _[A-Za-z0-9]匹配_或任何 A-Za-z0-9
// |或者
// [A-Z-a-z]匹配字符 A-Za-z
// )关闭群组
// \w*匹配 0+ 次一个单词 char
// $字符串结束

type N int

const (
	IDENTIFIER = iota
	EQUAL
	NUM
)

type rule struct {
	s string
	t int
}

var rules = [...]rule{{`(?:_[A-Za-z0-9]|[A-Z-a-z])\w*`, IDENTIFIER}, {`==`, EQUAL}, {`[0-9]+`, NUM}}
var regs = make([]*regexp.Regexp, len(rules))
var token = make([]rule, 0)

func init() {
	for _, v := range rules {
		regs = append(regs, regexp.MustCompile(v.s))
	}
}

func MakeToken(s string) {
	for len(s) != 0 {
		i := 0
		for ; i < len(regs); i++ {
			if loc := regs[i].FindStringIndex(s); len(loc) != 0 && loc[0] == 0 {
				token = append(token, rule{s[:loc[1]], rules[i].t})
				s = s[loc[1]:]
				break
			}
		}
		if i == len(regs) {
			fmt.Errorf("no match")
		}
	}
}
