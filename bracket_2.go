package main

import "container/list"

func longestValidParentheses_2(s string) int {
	L := list.New()
	var ct int
	ct = 0
	for _, ch1 := range s {
		if L.Len() == 0 {
			L.PushBack(ch1)
			continue
		}
		if ch1 == ')' {
			//fmt.Print(S.Gethead())
			if L.Back().Value == '(' {
				L.Remove(L.Back())
				ct += 2
				continue
			}
		} else {
			this := L.PushBack(ch1)
			if this.Prev().Value == '(' {
				ct = 0
			}
		}

	}

	return ct

}
