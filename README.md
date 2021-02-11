# leetcode-local-go

[![Build Status](https://travis-ci.com/day253/leetcode-local-go.svg?branch=master)](https://travis-ci.com/day253/leetcode-local-go)

LeetCode Golang 本地调试环境

## 使用步骤

### 1.配置leetcode插件

使用`ext install leetcode.vscode-leetcode`安装插件

配置leetcode插件，这样在sidebar点击题目会自动下载src目录

```
"leetcode.workspaceFolder": "/home/codespace/workspace/leetcode-local-go",
"leetcode.defaultLanguage": "golang"
```

### 3.本地调试

添加`package leetcode`，添加测试用例，根据需求修改函数名、测试用例输入输出等等，例子如下

```golang
package leetcode

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}

// @lc code=end
```

### 4.提交到leetcode测试更多case

`>LeetCode: Submit to LeetCode`提交答案到leetcode，如果不通过，添加到本地case直到测试通过再次提交。

## Rroblem Rank

* Top 400:
  * [Array](docs/array.md)
  * [Backtracking](docs/back_tracking.md)
  * [Binary Search](docs/binary_search.md)
  * [Bit Manipulation](docs/bit_manipulation.md)
  * [Design](docs/design.md)
  * [DFS & BFS](docs/dfs_bfs.md)
  * [Dynamic Programming](docs/dynamic_programming.md)
  * [Graph](docs/graph.md)
  * [LinkedList](docs/linked_list.md)
  * [Math](docs/math.md)
  * [Matrix](docs/matrix.md)
  * [PriorityQueue](docs/priorityqueue.md)
  * [Random](docs/random.md)
  * [Stack](docs/stack.md)
  * [String](docs/string.md)
  * [Topological Sort](docs/topological_sort.md)
  * [Tree](docs/tree.md)
  * [Trie](docs/trie.md)
  * [Union Find](docs/union_find.md)
* [Top 250](docs/top250.md)
* [Data Science](docs/data_science.md)

## 参考文档

[Leetcode插件](https://marketplace.visualstudio.com/items?itemName=leetcode.vscode-leetcode)

[Leetcode 分类顺序表](https://cspiration.com/leetcodeClassification)

[Leetcode 提交历史](https://leetcode.com/submissions)

[Leetcode-cn 提交历史](https://leetcode-cn.com/progress/)
