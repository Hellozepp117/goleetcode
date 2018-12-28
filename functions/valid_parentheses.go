package functions

import "container/list"

func isValid(s string) bool {
	target := []rune(s)
	l := list.New()
	for _, v := range target {
		switch v {
		case '(', '[', '{':
			l.PushBack(v)
		case ')':
			if l.Back() != nil && l.Back().Value == '(' {
				l.Remove(l.Back())
			} else {
				return false
			}
		case ']':
			if l.Back() != nil && l.Back().Value == '[' {
				l.Remove(l.Back())
			} else {
				return false
			}
		case '}':
			if l.Back() != nil && l.Back().Value == '{' {
				l.Remove(l.Back())
			} else {
				return false
			}
		}
	}
	if l.Len() == 0 {
		return true
	}
	return false
}
