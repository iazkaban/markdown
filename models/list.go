package models

type List []*ListNode

//列表行
type ListNode struct {
	Orderd  bool    //是否有编号
	Content Content //内容
	SubList *List   //子列表
	Block   *Block  //列表中块
}
