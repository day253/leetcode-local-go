# leetcode-local-go

LeetCode Golang 本地调试环境

## 使用步骤

### 1.配置leetcode插件

使用`ext install leetcode.vscode-leetcode`安装插件

配置leetcode插件，这样在sidebar点击题目会自动下载src目录

```
"leetcode.workspaceFolder": "/PathTo/leetcode-local-go/leetcode",
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

## 参考文档

[Leetcode插件](https://marketplace.visualstudio.com/items?itemName=leetcode.vscode-leetcode)

[Leetcode 分类顺序表](https://cspiration.com/leetcodeClassification)

[Leetcode 提交历史](https://leetcode.com/submissions)
