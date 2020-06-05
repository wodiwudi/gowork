package leetcode

/*
给定只含 "I"（增大）或 "D"（减小）的字符串 S ，令 N = S.length。
返回 [0, 1, ..., N] 的任意排列 A 使得对于所有 i = 0, ..., N-1，都有：
如果 S[i] == "I"，那么 A[i] < A[i+1]
如果 S[i] == "D"，那么 A[i] > A[i+1]

示例 1：
输出："IDID"
输出：[0,4,1,3,2]
示例 2：

输出："III"
输出：[0,1,2,3]
示例 3：

输出："DDI"
输出：[3,2,0,1]
*/
func diStringMatch(S string) []int {
	//1.长度为len(s)+1
	//如果第一个是I就从最小开始  D开始就是从最大开始
	//之后如果是D就从剩余数组里找最大的，I就从数组里找最小的
	length := len(S)
	if length == 0 {
		return nil
	}
	min, max := 0, length
	res := make([]int, length+1)
	for i := 0; i < length; i++ {
		if S[i] == 'I' {
			res[i] = min
			min++
		} else {
			res[i] = max
			max--
		}
	}
	res[length] = min
	return res
}
