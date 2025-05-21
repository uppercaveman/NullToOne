package main

/*
最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。

提示：

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 如果非空，则仅由小写英文字母组成

*/

func main() {
	println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	println(longestCommonPrefix([]string{"a"}))
	println(longestCommonPrefix([]string{"ab", "a"}))
	println(longestCommonPrefix([]string{"flower", "flower", "flower"}))
}

// 纵向扫描法
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := range len(strs[0]) {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
