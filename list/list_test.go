package list

import "testing"

func TestList(t *testing.T) {
	l := NewList(1)
	l.Add(2)
	l.PrintList()
}
