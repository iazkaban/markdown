package models

type Target uint8

const (
	TargetSelf  Target = iota //本页面打开
	TargetBlank               //新页面打开
)

//连接
type Link struct {
	Title  string //title
	URL    string //连接地址
	Target Target //目标打开方式
}

type ContentLink struct {
	Content Content //内容
	Link    *Link   //连接
}
