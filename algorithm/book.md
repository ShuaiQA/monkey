

# 动态规划

## 分析

- 首先是刻画最有结构(我们希望我们设置的结构可以由其子问题的最优结构所表达)
- 证明最优解的第一个组成部分是做出一个选择，例如钢条第一次切割位置的选择
- 在其可能的选择中，假设我们知道哪一种选择才是最优的解
- 根据当前的最优的解来进行选择，将当前的问题进行划分成多个子问题

课本上证明思想：作为构成原问题的最优解的组成部分，每个子问题的解就是它本身的最优解，反证法：假设子问题的解不是其自身的最优解，那么我们就可以从原问题的解中删除这些非最优的解，重新选择子问题的最优解，从而得到原问题的最优解，这与假设相矛盾。

我的理解是：首先我们需要从代码中得到的信息是，我们的最优解是有子问题的最优解所构成的，以钢管为例来说，我们最终的最优解是由倒数第二次进行选择之后进行解出来的。上面的证明思想是：最终的最优解，是由我们倒数第二次的最优解所构成的(如果我们知道子问题的最优解是正确的，我们选择每一处钢管的位置进行最大利润的选择就是最终的最优解)，那么我们思考凭什么我们说我们的子问题的最优解是正确的，反证法：我们需要结合程序进行思考，我们反的到底是什么，子问题不是最优的解，结合程序我们发现子问题最优的解是由子问题最优的解进行构成的，最终下去我们就会发现，最小的子问题就是我们初始化的dp边界问题的子问题的最优解，上面的反证法的意思是，如果当前的子问题不是最优的解(也就是说你初始化的时候思考边界要正确)，那么我们会不断的向上反馈直到原问题也不是最优的解(结合代码)，所以这与假设相矛盾。



写出相应的问题表达式之后，需要思考两个边界问题：

- 问题的边界，也就是dp的数组边界
- 划分选择的边界，思考如何对问题的划分能够将大的问题划分到边界问题上面，进行求解出来



## 钢管最大利润

钢管进行划分成相应的子钢管获取最大的利润

```go
// 写一下动态规划算法的套路
// 首先我们需要考虑的是对于一个大问题或分成一些小问题
// 我们在划分小问题的时候需要注意，作出一定的选择
// 并根据作出的选择选出我们需要的相应的最终结果
// arr := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30
// [0 1 5 8 10 13 17 18 22 25 30] [0 1 2 3 2 2 6 1 2 3 10]
func max_earn(e []int, length int) ([]int, []int) {
	//返回的是需要当前的长度的最大盈利，以及的切割的情况
	//e的下标代表当前的切割长度，值代表的是切割长度的盈利值
	arr := make([]int, length+1) //最大盈利的切割数
	dp := make([]int, length+1)  //dp[i]代表的是长度是i的最大盈利值
	dp[0] = 0
	var dfs func(cur int) int //根据当前的长度返回最大盈利值
	dfs = func(cur int) int {
		if cur == 0 {
			return 0
		}
		if dp[cur] != 0 {
			return dp[cur]
		}
		//将问题进行拆分成小问题
		for i := 1; i <= cur; i++ {
			// 根据拆分成的小问题来做出不同的选择
			if dp[cur-i]+e[i] > dp[cur] { 
				arr[cur] = i
				//注意当前的求解的值需要进一步递归来进行求解
				dp[cur] = dfs(cur-i) + e[i]
			}
		}
		return dp[cur]
	}
	dfs(length)
	return dp, arr
}
```



## 矩阵相乘

### debug递归

```go
// 调试
var count = 0

func debug(count int) {
	for i := 0; i < count; i++ {
		fmt.Print("   ")
	}
}	

//根据输入的矩阵的相应的行和列，我们输出的是对应的乘积的最小值和对应的划分区间情况
func min_cnt(arr [][]int) ([][]int, [][]int) {
	length := len(arr) //length = 6
	//dp[i][j] 代表的是[i,j]矩阵乘积的最小次数，注意左闭右闭，注意如果[i,i]只有一个矩阵那么是0
	dp := make([][]int, length)
	rr := make([][]int, length)
	for i := 0; i < length; i++ {
		rr[i] = make([]int, length)
		dp[i] = make([]int, length)
	}
	// 代表这[l,r]个的矩阵求最小乘积值
	var dfs func(l, r int) int
	dfs = func(l, r int) int {
		debug(count)
		fmt.Println(l, r)
		count++
		if l == r {
			count--
			debug(count)
			fmt.Println(l, r, "return")
			return 0
		}
		dp[l][r] = 99999999
		for j := l; j < r; j++ { //进一步求解[l,j],[j+1,r],需要有一个临界值,为了跳出dfs的情况
			temp := dfs(l, j) + dfs(j+1, r) + arr[l][0]*arr[j][1]*arr[r][1]
			if dp[l][r] > temp {
				rr[l][r] = j+1
				dp[l][r] = temp
			}
		}
		count--
		debug(count)
		fmt.Println(l, r, dp[l][r])
		return dp[l][r]
	}
	dfs(0, length-1)
	return dp, rr
}
```



## 最长公共子序列

"abcbdab", "bdcaba"              

```go
//返回lcs长度和字符串
func lcsLength(s1, s2 string) (int, string) {
	//dp[i][j]的值代表s1[i],s2[j]的时候lcs的最大长度
	//s1[i]==s2[j]   dp[i][j] = dp[i-1][j-1]+1
	//s1[i]!=s2[j]   dp[i][j] = max(dp[i-1][j],dp[i][j-1])
	//思考边界，我们考虑当字符串的长度是0的时候dp[0][j] = 0,dp[i][0] = 0
	//最大的长度应该是 dp[len(s1)][len(s2)]   0<=i<=len(s1)  0<=j<=len(s2)
	l1, l2 := len(s1), len(s2)
	//dp记录了长度，arr记录了应该向哪里走，1代表上，2代表左上，3代表左
	dp, arr := make([][]int, l1+1), make([][]int, l1+1)

	for i := 0; i < len(dp); i++ {
		dp[i], arr[i] = make([]int, l2+1), make([]int, l2+1)

	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				arr[i][j] = 2
			} else {
				if dp[i-1][j] >= dp[i][j-1] { //上面的优先
					arr[i][j] = 1
					dp[i][j] = dp[i-1][j]
				} else {
					arr[i][j] = 3
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	fmt.Println(dp, arr)
	return dp[l1][l2], printOneLcs(arr, l1, l2, s1)
}

// 输出的是逆序的情况
func printOneLcs(arr [][]int, i, j int, s1 string) string {
	if i == 0 || j == 0 {
		return ""
	}
	ans := ""
	if arr[i][j] == 2 { //代表左上
		ans += string(s1[i-1]) + printOneLcs(arr, i-1, j-1, s1)

	} else if arr[i][j] == 1 { //左
		ans += printOneLcs(arr, i-1, j, s1)
	} else { //上
		ans += printOneLcs(arr, i, j-1, s1)
	}
	return ans
}
```



## 构建最小期望的树

```go
//根据概率构建一个期望值最小的二叉树,返回的是该节点所在的层数
func BST(s []float32) ([][]float32, [][]int) { //	s := []float64{0.15, 0.1, 0.05, 0.1, 0.2}  len = 5
	//生成三个数组e[i][j]代表当前的[i,j]的节点的期望值
	//e[i][j],w[i][j],root[i][j]分别代表着,以[i,j]树的最小期望值,[i,j]的概率之和,[i,j]代表着树根是什么
	d := s[len(s)/2:] //len = 6
	w := computerW(s) //确定了相应的w[i][j] 代表叶子节点和非叶子节点的权重
	//接下来开始计算权重e,以及记录下root的数组
	//思考权重e怎么求,怎么对树进行子树的划分,划分之后需要加上什么才能够构成原有的情况
	//e[i][j] = e[i][k-1] + e[k+1][j] + w[i][j]
	//e[i][i-1] = w[i][i-1]  初始化边界
	//将k选择为根摘离出来,那么子树的所有节点就会降低一层,所以需要加上相应的权重
	//下面开始思考e的返回也就是边界, 根据公式可以知道e[1][5]就是答案,
	//len(e[0]) = len(d)  e[6][5] = w[6][5], len(e) = len(d)+1
	//root应与e类似
	e, root := make([][]float32, len(d)+1), make([][]int, len(d)+1)
	for i := 0; i < len(e); i++ {
		e[i] = make([]float32, len(d))
		root[i] = make([]int, len(d))
		if i > 0 {
			e[i][i-1] = w[i][i-1]
		}
	}
	var dfs func(i, j int) float32
	dfs = func(i, j int) float32 {
		if j == i-1 {
			return e[i][j]
		}
		var temp float32 = 999999 //设置一个不可能超过的期望
		for k := i; k <= j; k++ {
			cur := dfs(i, k-1) + dfs(k+1, j) + w[i][j]
			if cur < temp {
				temp = cur
				root[i][j] = k //当前节点 [i][j]以k为根节点
			}
		}
		e[i][j] = temp
		return temp
	}
	dfs(1, len(d)-1)
	return e, root
}

func computerW(s []float32) [][]float32 {
	k := s[:len(s)/2] //len = 5
	d := s[len(s)/2:] //len = 6
	w := make([][]float32, len(d)+1)
	//w[i][j] = sum(k[i]--k[j])+sum(d[i-1]--d[j])
	//思考w边界,当j = i-1  w[i][i-1] = d[i-1]
	//需要建立w[1][5]的情况就是总概率1，为什么不是w[0][4]节约空间 可以得到len(w[0]) = 6
	//w[6][5] = d[5]   //相关的len(w) = 7
	//因为w[1][0] 需要设置成 == d[0]便于求解相关的w
	for i := 0; i < len(w); i++ {
		w[i] = make([]float32, len(d))
		if i > 0 {
			w[i][i-1] = d[i-1] //初始化边界
		}
	}
	for l := 0; l < 5; l++ {
		for i := 1; i+l < len(w[i]); i++ {
			j := l + i
			w[i][j] = w[i][j-1] + k[j-1] + d[j]
		}
	}
	return w
}
```





# 贪心算法











































