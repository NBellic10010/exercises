package main

import (
	"container/list"
)

func converttoInt(e interface{}) int {
	switch v := e.(type) {
	case int:
		return v
	default:
		return -1
	}

}

func longestValidParentheses_2(s string) int {
	L := list.New()
	var length int
	length = 0
	L.PushBack(0)
	for _, ch1 := range s {
		if ch1 == ')' {
			//fmt.Print(S.Gethead())
			if L.Len() == 1 {
				L.Remove(L.Back())
				L.PushBack(0)
			}
			p := converttoInt(L.Back().Value)
			p = p + 2
			L.Remove(L.Back())
			q := converttoInt(L.Back().Value)
			q = q + p
			L.Remove(L.Back())
			L.PushBack(q)
			if converttoInt(L.Back().Value) > length {
				length = converttoInt(L.Back().Value)
				continue
			}
		} else {
			L.PushBack(0)
		}

	}
	return length
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if i+j == target {
				a := []int{i, j}
				return a
			}
		}

	}
	return []int{}
}
