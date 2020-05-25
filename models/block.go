package models

//引用块
type Block struct {
	Content Content //内容
	Block   *Block  //如果有子块的话，不为nil
	List    *List   //块中列表
}
