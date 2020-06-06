package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

/*
给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
示例 1:
输入: "Let's take LeetCode contest"
输出: "s'teL ekat edoCteeL tsetnoc"
*/

func reverseWords(s string) string {
	//1.先按空格分割
	//2.对每个字符数组反序
	//3.按空格拼接
	strArr := strings.Split(s, " ")
	for i := 0; i < len(strArr); i++ {
		//反序
		strArr[i] = reverseStrArr(strArr[i])
	}
	res := strings.Join(strArr, " ")
	return res
}
func reverseStrArr(s string) string {
	i, j := 0, len(s)-1
	runes := []rune(s)
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}
func TestFivetyOne(t *testing.T) {
	fmt.Println(reverseWords("abbc anc 45"))
}
