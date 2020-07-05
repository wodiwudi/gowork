package leetcode

import (
	"fmt"
	"testing"
)

/*
leetcode 14
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:
输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。

*/

func longestCommonPrefix(strs []string) string {
	//先把数组前两个比较 把公共部分赋值给一个变量，然后从3到最后一个依次与公共部分比较
	length := len(strs)
	if length == 1 {
		return string(strs[0])
	}
	if length == 0 {
		return ""
	}
	s1, s2 := strs[0], strs[1]
	public := getPublic(s1, s2)
	if length == 2 {
		return public
	} else {
		for i := 2; i < length; i++ {
			public = getPublic(public, strs[i])
			if public == "" {
				return ""
			}
		}
		return public
	}
}

func getPublic(s1, s2 string) string {
	ps1, ps2 := 0, 0
	public := []byte{}
	//寻找公共部分，双指针一起走，一旦遇到不一样的直接break
	for ps1 < len(s1) && ps2 < len(s2) {
		if s1[ps1] == s2[ps2] {
			public = append(public, s1[ps1])
			ps1++
			ps2++
		} else {
			break
		}
	}
	return string(public)
}

func TestFivetyEight(t *testing.T) {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
