package models

type List []*ListNode

type ListNode struct {
	Orderd  bool
	Body    string
	SubList List
}
