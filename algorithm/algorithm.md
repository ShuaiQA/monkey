# 递归回溯

目的：递归找到全解

递归出口符合条件的可行解

找到集合，根据规则选择可行解，添加进继续递归，推出递归删除之前添加。

## 排列

- 路径：也就是已经做出的选择。 
- 选择列表：也就是你当前可以做的选择。 
- 结束条件：也就是到达决策树底层，⽆法再做选择的条件。

```go
//全排列
func permute(nums []int) [][]int {
   ans := make([][]int, 0)
   //记录当前的决策
   temp := make([]int, 0)
   //查看当前节点是否已经做完决策
   vis := make([]bool, len(nums))
   var dfs func()
   dfs = func() {
      //思考什么时候可以不用做出决策,将该方案加入最后
      if len(temp) == len(nums) {
         a := make([]int, len(temp))
         copy(a, temp)
         ans = append(ans, a)
         return
      }
      //决策列表可以选择的集合
      for i := 0; i < len(nums); i++ {
         if !vis[i] { //判断当前节点是否满足作决策的地方
            //标记当前节点做出决策了
            temp = append(temp, nums[i])
            vis[i] = true 
            //继续向下
            dfs()
            //删除之前所作的决策
            vis[i] = false
            temp = temp[:len(temp)-1]
         }
      }
   }
   dfs()
   return ans
}
```





## 组合(待优化)

[划分为k个相等的子集](https://leetcode.cn/problems/partition-to-k-equal-sum-subsets/)

- 将n个元素拿出k个进行排列
- 将n个元素拿出k个进行组合

考虑顺序问题

**视⻆⼀，如果我们切换到这 n 个数字的视⻆，每个数字都要选择进⼊到 k 个桶中的某⼀个**

**视⻆⼆，如果我们切换到这 k 个桶的视⻆，对于每个桶，都要遍历 nums 中的 n 个数字，然后选择是否将当前遍历到的数字装进⾃⼰这个桶⾥。**

```go
//划分子集将数组nums集合划分k的子集使每一个子集之和相同
func canPartitionKSubsets(nums []int, k int) bool {
   ans := make([][]int, k)
   //将nums中的数字选择加入到k个桶中
   bucket := make([]int, k)
   var target int
   for i := 0; i < len(nums); i++ {
      target += nums[i]
   }
   target /= k
   var dfs func(cur int) bool
   dfs = func(cur int) bool {
      if cur == len(nums) {
         for i := 0; i < len(bucket); i++ {
            if bucket[i] != target {
               return false
            }
         }
         return true
      }
      //将当前第cur数字分别放入k个桶中
      for i := 0; i < len(bucket); i++ {
         if bucket[i]+nums[cur] > target {
            continue
         }
         //添加标记
         bucket[i] += nums[cur]
         ans[i] = append(ans[i], nums[cur])
         if dfs(cur + 1) {
            return true
         }
         //撤销选择
         bucket[i] -= nums[cur]
         ans[i] = ans[i][:len(ans[i])-1]
      }
      //nums[cur]装入哪一个桶中都不可以
      return false
   }
   v := dfs(0)
   fmt.Println(ans)
   return v
}
```



在桶的角度上进行思考

```go
//划分子集将数组nums集合划分k的子集使每一个子集之和相同
func canPartitionKSubsets(nums []int, k int) bool {
   ans := make([][]int, k)
   vis := make([]bool, len(nums))
   //将nums中的数字选择加入到k个桶中
   bucket := make([]int, k)
   var target int
   for i := 0; i < len(nums); i++ {
      target += nums[i]
   }
   target /= k
   var dfs func(k int) bool
   dfs = func(k int) bool {
      if k == 0 {
         return true
      }
      //当前桶已经满了找下一个
      if bucket[k-1] == target {
         return dfs(k - 1)
      }
      for i := 0; i < len(nums); i++ {
         if vis[i] { //当前的nums[i]已经被加入别的桶里面了
            continue
         }
         if nums[i]+bucket[k-1] > target { //当前的桶里面放不下
            continue
         }
         vis[i] = true
         bucket[k-1] += nums[i]
         ans[k-1] = append(ans[k-1], nums[i])
         if dfs(k) {
            return true
         }
         //需要进行递归回溯k所以不能直接在递归出口出直接k--
         vis[i] = false
         bucket[k-1] -= nums[i]
         ans[k-1] = ans[k-1][:len(ans[k-1])-1]
      }
      return false
   }
   v := dfs(k)
   fmt.Println(ans)
   return v
}
```





## 找套路

- 元素⽆重不可复选，即 nums 中的元素都是唯⼀的，每个元素最多只能被使⽤⼀次。

  以组合为例，如果输⼊ nums = [2,3,6,7]，和为 7 的组合应该只有 [7]

- 元素可重不可复选，即 nums 中的元素可以存在重复，每个元素最多只能被使⽤⼀次。

​		以组合为例，如果输⼊ nums = [2,5,2,1,2]，和为 7 的组合应该有两种 [2,2,2,1] 和 [5,2]。

- 元素⽆重可复选，即 nums 中的元素都是唯⼀的，每个元素可以被使⽤若⼲次。

​		以组合为例，如果输⼊ nums = [2,3,6,7]，和为 7 的组合应该有两种 [2,2,3] 和 [7]。



## 元素无重复不可复选

### 组合

给你一个整数数组 `nums` ，数组中的元素 **互不相同** 。返回该数组所有可能的子集（幂集）。

解集 **不能** 包含重复的子集。你可以按 **任意顺序** 返回解集。

```
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
```



```go
func subsets(nums []int) [][]int {
   ans := make([][]int, 0)
   temp := make([]int, 0)
   var dfs func(cur int)
   dfs = func(cur int) {
      //先序位置每一个节点都是一个子集,此处收集的是所有的节点(每一层的)
      //如果说返回k个数字的组合只需要让函数记录下dfs层数，在此处加入if判断就好了
      c := make([]int, len(temp))
      copy(c, temp)
      ans = append(ans, c)
      //选择cur后面的元素
      for i := cur; i < len(nums); i++ {
         temp = append(temp, nums[i])
         dfs(i + 1)
         temp = temp[:len(temp)-1]
      }
   }
   dfs(0)
   return ans
}

//输出：[ [] [1] [1 2] [1 2 3] [1 3] [2] [2 3] [3] ]
```

### 排列

给定一个不含重复数字的数组 `nums` ，返回其 所有可能的全排列。你可以 **按任意顺序** 返回答案。

```
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
```



```go
//全排列
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	//记录当前的决策
	temp := make([]int, 0)
	//查看当前节点是否已经做完决策
	vis := make([]bool, len(nums))
	var dfs func()
	dfs = func() {
		//思考什么时候可以不用做出决策
        //如果返回k个数据的排列在此处修改代码就好了
		if len(temp) == len(nums) {
			a := make([]int, len(temp))
			copy(a, temp)
			ans = append(ans, a)
			return
		}
		//决策列表可以选择的集合
		for i := 0; i < len(nums); i++ {
			if vis[i] {
				continue
			}
            //做出选择
			temp = append(temp, nums[i])
			vis[i] = true
            //进入下一层递归
			dfs()
			//删除所作的决策
			vis[i] = false
			temp = temp[:len(temp)-1]
		}
	}
	dfs()
	return ans
}
```



## 元素可重不可复选



### 组合

给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

```
输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
```



```go
func subsetsWithDup(nums []int) [][]int {
   ans := make([][]int, 0)
   temp := make([]int, 0)
   //排序让重复的元素在一起
   sort.Ints(nums)
   var dfs func(cur int)
   dfs = func(cur int) {
      //先序位置每一个节点都是一个子集,此处收集的是所有的节点(每一层的)
      //如果说返回k个数字的组合只需要让函数记录下dfs层数，在此处加入if判断就好了
      c := make([]int, len(temp))
      copy(c, temp)
      ans = append(ans, c)
      //选择cur后面的元素
      for i := cur; i < len(nums); i++ {
         if i > cur && nums[i] == nums[i-1] {
            continue
         }
         temp = append(temp, nums[i])
         dfs(i + 1)
         temp = temp[:len(temp)-1]
      }
   }
   dfs(0)
   return ans
}
```



### 排列

```go
//全排列
func permute(nums []int) [][]int {
   ans := make([][]int, 0)
   //记录当前的决策
   temp := make([]int, 0)
   vis := make([]bool, len(nums))
   sort.Ints(nums)
   var dfs func()
   dfs = func() {
      //思考什么时候可以不用做出决策
      //如果返回k个数据的排列在此处修改代码就好了
      if len(temp) == len(nums) {
         a := make([]int, len(temp))
         copy(a, temp)
         ans = append(ans, a)
         return
      }
      //决策列表可以选择的集合
      for i := 0; i < len(nums); i++ {
         if vis[i] {
            continue
         }
         //当前遍历的元素大于0后面判断才能继续
         //当前元素等于前一个元素，并且以前一个元素为开始的已经使用完
         if i > 0 && nums[i] == nums[i-1] && !vis[i-1] {
            continue
         }
         //做出选择
         temp = append(temp, nums[i])
         vis[i] = true
         //进入下一层递归
         dfs()
         //删除所作的决策
         vis[i] = false
         temp = temp[:len(temp)-1]
      }
   }
   dfs()
   return ans
}
```



## 元素无重可复选(修改第一个)

### 组合

给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。

candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

```
输入: candidates = [2,3,5], target = 8
输出: [[2,2,2,2],[2,3,3],[3,5]]
```

```go
func combinationSum(nums []int, target int) [][]int {
   ans := make([][]int, 0)
   temp := make([]int, 0)
   var dfs func(cur, sum int)
   dfs = func(cur, sum int) {
      if sum == target {
         c := make([]int, len(temp))
         copy(c, temp)
         ans = append(ans, c)
         return
      }
      //添加一个出口
      if sum > target {
         return
      }
      //选择cur后面的元素
      for i := cur; i < len(nums); i++ {
         temp = append(temp, nums[i])
         //此处递归可以从i当前重复的地方继续
         dfs(i, sum+nums[i])
         temp = temp[:len(temp)-1]
      }
   }
   dfs(0, 0)
   return ans
}
```





### 排列

```go
[1,2,3]的可重复的全排列：

[[1 1 1] [1 1 2] [1 1 3] [1 2 1] [1 2 2] [1 2 3] [1 3 1] [1 3 2] [1 3 3] [2 1 1]
 [2 1 2] [2 1 3] [2 2 1] [2 2 2] [2 2 3] [2 3 1] [2 3 2] [2 3 3] [3 1 1] [3 1 2]
 [3 1 3] [3 2 1] [3 2 2] [3 2 3] [3 3 1] [3 3 2] [3 3 3]]
```



```go
//全排列
func permute(nums []int) [][]int {
   ans := make([][]int, 0)
   //记录当前的决策
   temp := make([]int, 0)
   var dfs func()
   dfs = func() {
      if len(temp) == len(nums) {
         a := make([]int, len(temp))
         copy(a, temp)
         ans = append(ans, a)
         return
      }
      //决策列表可以选择的集合
      for i := 0; i < len(nums); i++ {
         temp = append(temp, nums[i])
         dfs()
         temp = temp[:len(temp)-1]
      }
   }
   dfs()
   return ans
}
```





# dfs

```go
next := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
vis := make([][]bool, len(grid))
//填充vis
var dfs func(x, y int)
dfs = func(x, y int) {
   if vis[x][y] { //已经遍历过了
      return
   }
   vis[x][y] = true
   for i := 0; i < len(next); i++ {
      nx := x + next[i][0]
      ny := y + next[i][1]
      if nx.ny边界判断 {
         dfs(nx, ny)
      }
   }
}
```





# bfs

```go
step := 0
queue := make([]*TreeNode, 0)
var bfs func(cur, target *TreeNode)
bfs = func(cur, target *TreeNode) {
   queue = append(queue, cur)
   for len(queue) != 0 {
      size := len(queue)
      for i := 0; i < size; i++ {
         poll := queue[0]
         if poll == target {
            return
         }
         queue = queue[1:]
         for 遍历下一个节点 {
            有时候需要判断当前节点是否已经被访问过了(可以使用数组或者set进行标记)
            将所有的节点添加到队列中
         }
      }
      step++
   }
}
```



# 动态规划

「状态」，「选择」，「dp 数组的定义」

动态规划是什么？解决动态规划问题有什么技巧？如何学习动态规划？

⾸先，动态规划问题的⼀般形式就是求最值。动态规划其实是运筹学的⼀种最优化⽅法。

既然是要求最值，核⼼问题是什么呢？求解动态规划的核⼼问题是穷举。因为要求最值，肯定要把所有可⾏ 的答案穷举出来，然后在其中找最值呗。⾸先，动态规划的穷举有点特别，因为这类问题存在「重叠⼦问题」，如果暴⼒穷举的话效率会极其低下， 所以需要「备忘录」或者「DP table」来优化穷举过程，避免不必要的计算。 ⽽且，动态规划问题⼀定会具备「最优⼦结构」，才能通过⼦问题的最值得到原问题的最值。

**重叠⼦问题、最优⼦结构、状态转移⽅程就是动态规划三要素**

明确 base case -> 明确「状态」-> 明确「选择」 -> 定义 dp 数组/函数的含义。

```
# 初始化 base case
dp[0][0][...] = base
# 进⾏状态转移
for 状态1 in 状态1的所有取值：
 	for 状态2 in 状态2的所有取值：
		 for ...
 			dp[状态1][状态2][...] = 求最值(选择1，选择2...)
```



**递归算法的时间复杂度怎么计算？就是⽤⼦问题个数乘以解决⼀个⼦问题需要的时间。**

**动态规划的另⼀个重要特性最优⼦结构。**

动态规划问题，因为它具有「最优⼦结构」的。要符合「最优⼦结构」，⼦问题间必须互相独⽴。



## 凑零钱

给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的

```go
输入：coins = [1, 2, 5], amount = 11
输出：3 
解释：11 = 5 + 5 + 1
```

首先我们使用暴力进行求解问题，以目标值为11进行求解所有的可行解，比如`[5,5,1]`或者`[2,2,2,2,2,1]`，将所有的可行解进行求出来然后求最小的长度，即为最少的硬币数目。**注意不能使用贪心从大的向小的开始进行比如说`coins = [5,4,1],amount = 12`，使用贪心可能的是`[5,5,1,1]`，但是应该是`[4,4,4]`。**

暴力求解：

```go
func coinChange(coins []int, amount int) int {
   sort.Slice(coins, func(i, j int) bool {
      return coins[i] > coins[j]
   })
   ans := make([][]int, 0)
   temp := make([]int, 0)
   var dfs func(cur, sum int)
   dfs = func(cur, sum int) {
      if sum == amount {
         c := make([]int, len(temp))
         copy(c, temp)
         ans = append(ans, c)
         return
      }
      for i := cur; i < len(coins); i++ {
         if sum+coins[i] <= amount {
            temp = append(temp, coins[i])
            dfs(i, sum+coins[i])
            temp = temp[:len(temp)-1]
         }
      }
   }
   dfs(0, 0)
   if len(ans) == 0 {
      return -1
   }
   max := math.MaxInt
   for i := 0; i < len(ans); i++ {
      if len(ans[i]) < max {
         max = len(ans[i])
      }
   }
   fmt.Println(ans)
   return max
}
```

使用dp

- 确定 base case，显然⽬标⾦额 amount 为 0 时算法返回 0，因为不需要任何硬币就已经凑 出⽬标⾦额了。 
- 确定「状态」，也就是原问题和⼦问题中会变化的变量。由于硬币数量⽆限，硬币的⾯额也是题⽬给定的，只有⽬标⾦额会不断地向 base case 靠近，所以唯⼀的「状态」就是⽬标⾦额 amount。
-  确定「选择」，也就是导致「状态」产⽣变化的⾏为。⽬标⾦额为什么变化呢，因为你在选择硬币，你每 选择⼀枚硬币，就相当于减少了⽬标⾦额。所以说所有硬币的⾯值，就是你的「选择」。 
- 明确 dp 函数/数组的定义。我们这⾥讲的是⾃顶向下的解法，所以会有⼀个递归的 dp 函数，⼀般来说函 数的参数就是状态转移中会变化的量，也就是上⾯说到的「状态」；函数的返回值就是题⽬要求我们计算的 量。就本题来说，状态只有⼀个，即「⽬标⾦额」，题⽬要求我们计算凑出⽬标⾦额所需的最少硬币数量。 
- 所以我们可以这样定义 dp 函数：dp(n) 表示，输⼊⼀个⽬标⾦额 n，返回凑出⽬标⾦额 n 所需的最少硬币 数量。

```go
func coinChange(coins []int, amount int) int {
   sort.Slice(coins, func(i, j int) bool {
      return coins[i] > coins[j]
   })
   dp := make([]int, amount+1)
   for i := 0; i < len(dp); i++ {
      dp[i] = math.MaxInt
   }
   //base
   dp[0] = 0
   for i := 1; i < len(dp); i++ { //i代表的是当前amount下标
      for k := 0; k < len(coins); k++ { //上面的i只有可能是由coins中的元素[i-1]构成的
         last := i - coins[k]                     //当前代表的是i-1
         if last < 0 || dp[last] == math.MaxInt { //超出边界或者是上一个i-1没有解当前i也就没有解
            continue
         }
         dp[i] = min(dp[i], dp[last]+1)
      }
   }
   if dp[amount] == math.MaxInt {
      return -1
   }
   return dp[amount]
}

func min(x, y int) int {
   if x > y {
      return y
   }
   return x
}
```







## 双字符串进行对比

给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

- 插入一个字符

- 删除一个字符
- 替换一个字符

```
输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
```

首先从最小的进行分析，如果两个字符串都是空的，那么直接返回0，如果一个是空的，返回另一个字符串的长度，当前就是初始化的dp数组，那么接下来就要想想怎么继续延伸下去。该`dp[i][j]代表的是s1到i以及s2到j最小的情况答案`，如果当前的`s[1]==s[2]`那么`dp[i][j] =dp[i-1][j-1]`，如果不相等那么我们需要考虑插入、删除、替换操作，在不相等的情况考虑替换那么`dp[i][j] = dp[i-1][j-1]+1`考虑删除和插入操作，就是`dp[i][j] = dp[i-1][j]+1或者dp[i][j-1]+1`取上面的最小的情况的值就是答案。

```go
func minDistance(s1 string, s2 string) int {
    dp:=make([][]int,len(s1)+1)
    for i:=0;i<len(dp);i++{
        dp[i] = make([]int,len(s2)+1)
    }
    //初始化base
    for i:=0;i<len(dp);i++{
        dp[i][0] = i
    }
    for j:=0;j<len(dp[0]);j++{
        dp[0][j] = j
    }
    for i:=1;i<len(dp);i++{
        for j:=1;j<len(dp[i]);j++{
            if s1[i-1]==s2[j-1]{
                dp[i][j] = dp[i-1][j-1]
            }else{
                dp[i][j] = min(min(dp[i-1][j],dp[i][j-1]),dp[i-1][j-1])+1
            }
        }
    }
    // fmt.Println(dp)
    return dp[len(s1)][len(s2)]
}

func min(x,y int)int{
    if x>y{
        return y
    }
    return x
}
```



## 正则表达式







## 背包问题（子集问题的优化）

常见的套路就是从数组中选择元素。

- 查看数组中的子集，凑成target值有多少种最大的情况
- 数组中的子集能不能构成目标值target
- 0-1背包问题
- 完全背包问题

常见的套路

```go
//0-1背包问题,当前的状态是选择当前元素不与选择当前元素组合
dp[i][j] = dp[i-1][j]与dp[i-1][j-nums[i-1]]
//完全背包问题，当前状态是选择当前元素不与选择当前元素组合
dp[i][j] = dp[i-1][j]与dp[i][j-nums[i-1]]
```







## 股票选择问题(状态转移)

[买卖股票的最佳时机](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/)

[买卖股票的最佳时机 II](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/)

[买卖股票的最佳时机 III](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/)

[买卖股票的最佳时机 IV](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/)



查找状态，可以发现每一天的状态有当前持有股票，处在第j次购买时。每一天的状态有buy(购买)，sell(出售)，rest(不变)

状态的转移：假设0代表当前没有持有股票，1代表当前持有股票

**思考当前第i天的状态能做出的选择有哪些。**

**第i天能够做出的选择有rest、sell、buy选择。**

**注意dp数组仅仅只是一个确定的状态，当前的每一个状态是由dp数组的下标来确定的，比如说`dp[i][j][0]`代表当前i天没有持有股票，当前没有持有股票代表已经当前已经做出了选择，思考当前的状态是由前一个什么状态经过什么操作构成的。**

````
dp[i][j][0]代表当前是处于第i天，当前是处在完成第j次购买情况，当前的状态是没有购买。
dp[i][j][1]代表当前是处在第i天，当前是处在完成第j次购买情况，当前的状态是已经购买完(没有抛出)。
状态转移方程
dp[i][j][0] = max(dp[i-1][j][0],dp[i-1][j][1]+price[i])
                 //今天的状态不变 //今天将之前购买过的股票进行抛出
dp[i][j][1] = max(dp[i-1][j][1],dp[i-1][j-1][0]-price[i])
             //今天的状态不买不买 //今天的状态是要购买，也就是需要花费一次交易，并减去花的股票钱

	//初始化base
	//最初-1天不可能持有股票,在没有一次交易的情况下也是0
	dp[0][...][0]==0   dp[...][0][0]==0
	//最初-1天，如果持有股票(此处是不可能的事件，算法要求最大值，当前赋值最小值)d
	dp[0][...][1]==math.MinInt   dp[...][0][1]==math.MinInt
	//选择0次，代表没有利润
````



```go
func maxProfit(k int, prices []int) int {
   //需要初始化第0天最初的状态，所以需要+1
   dp := make([][][]int, len(prices)+1)
   //初始化没有一次交易情况，并且需要交易k次所以需要K+1
   for i := 0; i < len(dp); i++ {
      dp[i] = make([][]int, k+1)
   }
   //每一天的状态，记录交易了几次，当前有没有持有股票
   for i := 0; i < len(dp); i++ {
      for j := 0; j < len(dp[i]); j++ {
         dp[i][j] = make([]int, 2)
      }
   }
   //dp[i][j][0]代表第i天，以选择j次，当前没持有股票
   //dp[i][j][1]代表第i天，以选择j次，当前持有股票
   //初始化base
   //最初-1天不可能持有股票，dp[0][k][0]==0
   //最初-1天，如果持有股票(此处是不可能的事件，算法要求最大值，当前赋值最小值)dp[0][k][1]==math.MinInt
   //选择0次，代表没有利润
   for j := 0; j < len(dp[0]); j++ {
      dp[0][j][0] = 0
      dp[0][j][1] = math.MinInt
   }
   for i := 0; i < len(dp); i++ {
      dp[i][0][0] = 0
      dp[i][0][1] = math.MinInt
   }

   for i := 1; i < len(dp); i++ {
      for j := 1; j < len(dp[i]); j++ {
         //状态转移方程
         dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i-1])
         dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i-1])
      }
   }
   return dp[len(dp)-1][k][0]
}
```



**k==1的时候的问题。**

给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

```go
func maxProfit(prices []int) int {
	//需要初始化第0天最初的状态，所以需要+1
	dp := make([][]int, len(prices)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	//初始化base
    dp[0][0] = 0
    dp[0][1] = math.MinInt
	for i := 1; i < len(dp); i++ {
		//对上面的状态转移方程进行化简。
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i-1])
		dp[i][1] = max(dp[i-1][1], -prices[i-1])
	}
	return dp[len(dp)-1][0]
}
```



**k==无穷(与k无关)**

给定一个数组 prices ，其中 prices[i] 表示股票第 i 天的价格。

在每一天，你可能会决定购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以购买它，然后在 同一天 出售。
返回 你能获得的 最大 利润 。

```go
func maxProfit(prices []int) int {
    dp:=make([][]int,len(prices)+1)
    for i:=0;i<len(dp);i++{
        dp[i] = make([]int,2)
    }
    dp[0][0] = 0
    dp[0][1] = -prices[0]
    for i:=1;i<len(dp);i++{
        dp[i][0] = max(dp[i-1][1]+prices[i-1],dp[i-1][0])
        dp[i][1] = max(dp[i-1][0]-prices[i-1],dp[i-1][1])
    }
    // fmt.Println(dp)
    return dp[len(prices)][0]
}
```

**k==无穷，含有三个状态的时候**

给定一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格 。

设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:

卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

```
输入: prices = [1,2,3,0,2]
输出: 3 
解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
```

```go
func maxProfit(prices []int) int {
	//需要初始化第0天最初的状态，所以需要+1
	dp := make([][]int, len(prices)+1)
	//初始化没有一次交易情况，并且需要交易k次所以需要K+1

	//每一天的状态，记录交易了几次，当前有没有持有股票
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 3)
	}
    //0代表没有持有，1代表买入，2代表冷冻期
	dp[0][0] = 0
	dp[0][1] = math.MinInt
	dp[0][2] = math.MinInt

	for i := 1; i < len(dp); i++ {
		//状态转移方程
		dp[i][0] = max(dp[i-1][0], dp[i-1][2])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i-1])
        //冷冻期没有自旋状态只有从1买入到达的状态
		dp[i][2] = dp[i-1][1]+prices[i-1]
	}
    // fmt.Println(dp)
   	//返回冷冻期或者没有持有股票的状态的最大值。
	return max(dp[len(dp)-1][0],dp[len(dp)-1][2])
}
```



## 打家劫舍问题

你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。



对于状态的选择，对于每一户人家我们有选择可以偷以及可以不偷，这是两种状态。再加上每一户人家。

我们设置这样的转移方程

**思考由当前的问题转换成子问题是怎么样的：**

- **首先当前问题的选择：当前第i户到底偷不偷**
- **选择当前问题为偷的时候转换成子问题，当前偷的话前一户必定不能偷(转成前一户的子问题)**
- **选择当前问题不偷的话转换成子问题，当前不偷的话，前一个可以偷也可以不偷求解最大值的子问题。**

```
dp[i][0]  代表当前第i户人家没有被偷的最大价值和
dp[i][1]  代表当前第i户人家被偷了的最大价值和
思考状态转移方程
dp[i][0] = max(dp[i-1][0],dp[i-1][1])
//当前用户没有被偷，选取前一个状态偷或者不偷的最大值
dp[i][1] = dp[i-1][0]+nums[i-1]
//当前用户被偷了，选择前一个状态没有被偷并加上偷取当前用户的金额

```

```go
func rob(nums []int) int {
    dp:=make([][]int,len(nums)+1)
    for i:=0;i<len(dp);i++{
        dp[i] = make([]int,2)
    }
    //初始化base
    dp[0][0] = 0
    dp[0][1] = math.MinInt
    for i:=1;i<len(dp);i++{
        //状态转移方程
        dp[i][0] = max(dp[i-1][1],dp[i-1][0])
        dp[i][1] = dp[i-1][0]+nums[i-1]
    }
    return max(dp[len(nums)][0],dp[len(nums)][1])
}
```



对于环形数组的情况，我们需要注意首尾元素的选取，首先需要判断选择起始位置的元素那么末尾就不能选择，反之亦然。这个问题就变成了求解`nums[1:] 、nums[:len(nums)-1]`的情况了

对于树状房子进行选择。[打家劫舍 III](https://leetcode.cn/problems/house-robber-iii/)

```go
func rob(root *TreeNode) int {
    var last func(root *TreeNode) (int,int)  //返回当前节点，偷和不偷的情况
    last = func(root *TreeNode) (int,int) {
        if root == nil{
            return 0 , 0
        }
        l1,r1 := last(root.Left)    //l代表偷，r代表不偷 
        l2,r2 := last(root.Right)
        //注意如果当前节点偷的话，那么子节点必定是不偷的状态
        //但是如果当前节点不偷的话，子节点不一定是偷的状态，子节点任意组合(选取最大值)
        return r1 + r2 + root.Val , max(l1 , r1) + max(l2, r2)
    }
    l1,l2:=last(root)
    return max(l1,l2)
}
```



## 两人博弈问题

Alice 和 Bob 用几堆石子在做游戏。一共有偶数堆石子，排成一行；每堆都有 正 整数颗石子，数目为 piles[i] 。

游戏以谁手中的石子最多来决出胜负。石子的 总数 是 奇数 ，所以没有平局。

Alice 和 Bob 轮流进行，Alice 先开始 。 每回合，玩家从行的 开始 或 结束 处取走整堆石头。 这种情况一直持续到没有更多的石子堆为止，此时手中 石子最多 的玩家 获胜 。

假设 Alice 和 Bob 都发挥出最佳水平，当 Alice 赢得比赛时返回 true ，当 Bob 赢得比赛时返回 false 。

```
输入：piles = [5,3,4,5]
输出：true
解释：
Alice 先开始，只能拿前 5 颗或后 5 颗石子 。
假设他取了前 5 颗，这一行就变成了 [3,4,5] 。
如果 Bob 拿走前 3 颗，那么剩下的是 [4,5]，Alice 拿走后 5 颗赢得 10 分。
如果 Bob 拿走后 5 颗，那么剩下的是 [3,4]，Alice 拿走后 4 颗赢得 9 分。
这表明，取前 5 颗石子对 Alice 来说是一个胜利的举动，所以返回 true 。
```



**先⼿在做出选择之后，就成了后⼿，后⼿在对⽅做完选择后，就变成了先⼿。这种 ⻆⾊转换使得我们可以重⽤之前的结果，典型的动态规划标志。**

数组`dp[i][j].a`的定义是区间长度是`[i,j]`先手获取的最大值。

数组`dp[i][j].b`的定义是区间长度是`[i,j]`后手获取的最大值。

**思考由当前的问题转换成子问题是怎么样的：**

- **首先状态的选取，对于现有的区间有先手和后手状态之分**
- **`dp[i][j]`数组结构体的a，b的定义已经代表选择了先后手**
- **那么思考当前的先手是由之前的什么子问题，经过什么操作构成的**
- **以`dp[i][j].a`为例代表的是求解区间`[i,j]`的先手的最大值，该问题由什么子问题构成，对于该区间进行的选择可以选择`i`以及后手区间`[i+1,j]`构成，或者是选择`j`后手区间`[i,j-1]`构成，因为当前区间进行选择完毕后，就会变成后手的。**

```
分析先手和后手是相互转换的
dp[i][j].a = max(dp[i+1][j].b+nums[i],dp[i][j-1].b+nums[j])
//注意此时的先手是由先选择完i然后就变成后手选择区间[i+1][j]构成的，以及先手选择j就变成后手选择区间[i][j-1]构成

//对于某个区间的后手问题，其实就是先手选择完毕后，剩下的区间的先手问题。 
```



```go
func stoneGame(nums []int) (int, int) {
   n := len(nums)
   //初始化数组
   dp := make([][]Pair, n)
   for i := 0; i < n; i++ {
      dp[i] = make([]Pair, n)
   }
   //初始化base
   for i := 0; i < n; i++ {
      dp[i][i].a = nums[i]
      dp[i][i].b = 0
   }
   //斜着递推数组
   for d := 1; d < n; d++ {
      for i := 0; i < n-d; i++ {
         j := i + d //i,j代表的是斜着递推数组
         //其中l与r代表的是选择左边和右边的最大的情况
         //选择左边的递推方程
         l := nums[i] + dp[i+1][j].b
         r := nums[j] + dp[i][j-1].b
         if l > r {
            dp[i][j].a = l
            //选择左边，当前的后手变成删去左边的区间的先手
            dp[i][j].b = dp[i+1][j].a
         } else {
            dp[i][j].a = r
            dp[i][j].b = dp[i][j-1].a
         }
      }
   }
   return dp[0][n-1].a, dp[0][n-1].b
}
```



## 鸡蛋掉落问题(待优化)

给你 k 枚相同的鸡蛋，并可以使用一栋从第 1 层到第 n 层共有 n 层楼的建筑。

已知存在楼层 f ，满足 0 <= f <= n ，任何从 高于 f 的楼层落下的鸡蛋都会碎，从 f 楼层或比它低的楼层落下的鸡蛋都不会破。

每次操作，你可以取一枚没有碎的鸡蛋并把它从任一楼层 x 扔下（满足 1 <= x <= n）。如果鸡蛋碎了，你就不能再次使用它。如果某枚鸡蛋扔下后没有摔碎，则可以在之后的操作中 重复使用 这枚鸡蛋。

请你计算并返回要确定 f 确切的值 的 最小操作次数 是多少？

**确定最大的没有碎的层数**

```
输入：k = 1, n = 2
输出：2
解释：
鸡蛋从 1 楼掉落。如果它碎了，肯定能得出 f = 0 。 
否则，鸡蛋从 2 楼掉落。如果它碎了，肯定能得出 f = 1 。 
如果它没碎，那么肯定能得出 f = 2 。 
因此，在最坏的情况下我们需要移动 2 次以确定 f 是多少。 
```



思考对于鸡蛋掉落问题最原始的方法是，从一楼掉落碎了，答案就是0。没有碎，那么继续向上第二层第二层碎了答案就是1，不断地继续。对于有无数个鸡蛋来说，我们可以使用二分来进行，以`[1,9]`层为例，我们可以二分查找取中间的5层进行以鸡蛋实践，如果该鸡蛋没有碎那么就向上查找`[6,9]`(不需要包含5，因为此处的6就代表第1层，如果在第1层碎了返回第0层的5)，如果碎了就像下进行查找`[1,4]`，然后不断的二分进行查找，但是这道题来说鸡蛋的数目并不是无数多个，这个时候就需要我们进行考虑的了。

那么我最初确定的状态就是`dp[i][j][k]`代表确定从i层到j层有k个鸡蛋确定值得最小操作数，修改为`dp[N][k]`。

因为这个题的要求并没有要求返回哪一个楼层，所以区间i到j可以写成有N层表示，使用k个鸡蛋来确定。



```
分析：
对于dp[N][j] = m，代表楼层有N层有j个鸡蛋，最后求解的最小操作数是m
对于状态转移来说，对于N层我们可以不断地进行尝试，从1层一直到n层扔鸡蛋进行尝试，对于每一次拆分楼梯的答案我们需要求解最小的值
拆分完楼梯后，进行扔鸡蛋的测试我们这个时候进行分析，如果在第k层扔鸡蛋碎了，那么j-1向下的楼层开始遍历，如果在第k层鸡蛋没有碎，那么j不变向上测试。对于每一次的测试我们需要取最大的值，因为如果选择最小的值，对于另一边就不能确定到底在第几层碎了。
```



```go
func superEggDrop(k int, n int) int {
    //初始化dp
    dp:=make([][]int,n+1)
    for i:=0;i<len(dp);i++{
        dp[i] = make([]int,k+1)
    }
    fill(dp)
    //初始化base
    for i:=1;i<len(dp);i++{
        dp[i][1] = i
    }
    for j:=1;j<len(dp[0]);j++{
        dp[1][j] = 1
    }
    for i:=2;i<len(dp);i++{  //i代表当前有多少楼层
        for j:=2;j<len(dp[0]);j++{   //j代表当前有多少个鸡蛋
            for k:=1;k<i;k++{   //k代表当前区分的楼层，以k来划分楼层 
                 dp[i][j] = min(dp[i][j],max(dp[k-1][j-1],dp[i-k][j])+1) //碎了向下遍历，没有碎向上遍历
            }
        }
    }
    // fmt.Println(dp)
    return dp[n][k]
}

```



根据单调函数对第三层for循环进行二分查找。

```go
func superEggDrop(k int, n int) int {
    //初始化dp
    dp:=make([][]int,n+1)
    for i:=0;i<len(dp);i++{
        dp[i] = make([]int,k+1)
    }
    fill(dp,math.MaxInt)
    //初始化base
    for i:=1;i<len(dp);i++{
        dp[i][1] = i
    }
    for j:=1;j<len(dp[0]);j++{
        dp[1][j] = 1
    }
    for i:=2;i<len(dp);i++{  //i代表当前有多少楼层
        for j:=2;j<len(dp[0]);j++{   //j代表当前有多少个鸡蛋
            l,r:=1,i
            for l<=r{
                mid:=(l+r)/2
                broke := dp[mid-1][j-1]
                not_broke := dp[i-mid][j]
                if broke>not_broke{
                   r = mid-1
                    dp[i][j] = min(dp[i][j],broke+1)
                }else{
                    l = mid+1
                    dp[i][j] = min(dp[i][j],not_broke+1)
                }
            }
        }
    }
    // fmt.Println(dp)
    return dp[n][k]
}
```





## [地下城游戏](https://leetcode.cn/problems/dungeon-game/description/)

思考如果是从左上向右下进行递推，那么我们需要进行标记两个数字，一个是到达当前位置的最小需求血量、还有到达当前位置还剩下多少血量，但是对于下面我们做选择的时候，就需要以来当前的两个的值进行递推了，并且不断地进行比较最小需求血量以及当前的血量。来进行下一步操作。

思考另一种简单的方式，如果是从当前的位置到最后的位置进行，标记最小需求血量来说的话，那么就不需要进行纠结了，可以方向递推，比如说当前的`[i,j]`可以由`[i+1,j]`以及`[i,j+1]`的最小需求血量来进行求解。

```go
func calculateMinimumHP(dungeon [][]int) int {
    n,m:=len(dungeon),len(dungeon[0])
    memo:=make([][]int,len(dungeon))
    for i:=0;i<len(dungeon);i++{
        memo[i] = make([]int,len(dungeon[0]))
        for j:=0;j<len(memo[i]);j++{
            memo[i][j] = math.MinInt
        }
    }
    var dp func(x,y int)int
    dp = func(x,y int)int{
        if x==n-1&&y==m-1{   //处于右下角
            if dungeon[x][y]>=0{
                return 0
            }else{
                return -dungeon[x][y]
            }
        }
        if x==n||y==m{  //超出边界返回最大值
            return math.MaxInt
        }
        //已经存在memo中返回
        if memo[x][y]!=math.MinInt{
            return memo[x][y]
        }
        //转移方程
        res:= min(dp(x,y+1),dp(x+1,y))-dungeon[x][y]
        //进行备忘
        if res<=0{
            memo[x][y] = 0
        }else{
            memo[x][y] = res
        }
        return memo[x][y]
    }
    return dp(0,0)+1
}
```



[自由之路](https://leetcode.cn/problems/freedom-trail/)

```go
func findRotateSteps(ring string, key string) int {
    m:=make(map[byte][]int)
    for i:=0;i<len(ring);i++{
        m[ring[i]] = append(m[ring[i]],i)
    }
    memo:=make([][]int,len(ring))
    for i:=0;i<len(memo);i++{
        memo[i] = make([]int,len(key))
    }
    var dp func(i,j int)int
    dp = func(i,j int) int {
        if j==len(key){
            return 0
        }
        if memo[i][j]!=0{
            return memo[i][j]
        }
        res := math.MaxInt
        for _,k:=range m[key[j]]{
            del:= abs(k-i)
            //选择顺逆波动最小值
            del = min(del,len(ring)-del)
            //将指针波动到k上面，使j+1,进行子问题的查找
            subProblem:=dp(k,j+1)
            //将子问题以及当前的旋转次数，拼接次数添加到当前的结果上
            res = min(res,1+del+subProblem)
        }
        memo[i][j] = res
        return res
    }
    // fmt.Println(m)
    return dp(0,0)
}
```



## 中转k次最小路径问题

[K 站中转内最便宜的航班](https://leetcode.cn/problems/cheapest-flights-within-k-stops/)

```go
type Edge struct{
    from,length int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    memo:=make([][]int,k+2)
    for i:=0;i < len(memo);i++{
        memo[i] = make([]int,n)
    }
    graph:=make([][]Edge,n)
    //建立一个to到from的图
    for i:=0;i<len(flights);i++{
        from:=flights[i][0]
        to:=flights[i][1]
        length:=flights[i][2]
        graph[to] = append(graph[to],Edge{from,length})
    }
    var dp func(k,to int)int
    dp = func(k,to int)int{
        if to==src{
            return 0
        }
        if k==0&&to!=src{
            return -1
        }
        if memo[k][to]!=0{
            return memo[k][to]
        }
        res := math.MaxInt
        for _,e:=range graph[to]{
            from := e.from
            length :=e.length
            //-1代表没有中转了，math.MaxInt代表没有路径
            if dp(k-1,from)!=-1&&dp(k-1,from)!=math.MaxInt{
                res = min(res,dp(k-1,from)+length)
            }
        }
        memo[k][to] = res
        return res 
    }
    if dp(k+1,dst)==math.MaxInt{
        return -1
    }
    return dp(k+1,dst)
}

```





