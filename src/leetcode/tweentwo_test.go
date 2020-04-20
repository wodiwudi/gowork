package leetcode

import (
	"errors"
	"fmt"
	"testing"
)

/*
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/
func TestTweenTwo(t *testing.T) {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	var result []string
	var generate func(min, max int, tmp string)
	generate = func(min, max int, tmp string) {
		if min >= max {
			result = append(result, tmp)
			return
		}
		generate(min+1, max, tmp+"(")
		generate(min+1, max, tmp+")")
	}
	generate(0, 2*n, "")
	//合法性校验  栈操作
	fin := []string{}
	for _, v := range result {
		if judge(v) {
			fin = append(fin, v)
		}
	}

	return fin
}

func judge(s string) bool {
	lens := len(s)
	stack := make([]rune, 0, lens)

	push := func(x rune) error {
		length := len(stack)
		if cap(stack) <= length {
			return errors.New("stack is full")
		}
		stack = stack[:length+1]
		stack[length] = x
		return nil
	}

	pop := func() (x rune, err error) {
		length := len(stack)
		if length == 0 {
			return 0, errors.New("stack is empty")
		}
		x = stack[length-1]
		stack = stack[:length-1]
		return x, nil
	}
	for _, v := range s {
		if v == '(' {
			err := push(v)
			if err != nil {
				return false
			}
		} else {
			_, err := pop()
			if err != nil {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
func TestJudge(t *testing.T) {
	test := "{{{}}{"
	//测试栈操作
	fmt.Println(judge(test))
}
