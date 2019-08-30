package main

type Stack struct {
	room    [100]int32
	head    int
	MAXSIZE int
}

func ini() *Stack {
	s := new(Stack)
	s.head = -1
	s.MAXSIZE = 100
	return s
}

func (s *Stack) Gethead() int32 {
	if s.head == -1 {
		return -1
	}
	return s.room[s.head]
}

func (s *Stack) Push(p int32) {
	s.head = s.head + 1
	s.room[s.head] = p
}

func (s *Stack) Pop(p int32) {
	s.room[s.head] = p
	s.head = s.head - 1
}

func longestValidParentheses(s string) int {
	S := ini()
	var ct int
	ct = 0
	for _, ch1 := range s {
		if ch1 == '(' {
			S.Push(ch1)
		} else if ch1 == ')' {
			//fmt.Print(S.Gethead())
			S.Pop(0)
			ct += 2
		}

	}

	return ct

}
