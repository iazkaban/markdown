package models

type CellAlign uint8

const (
	CellAlignCentor CellAlign = iota
	CellAlignLeft
	CellAlignRight
)

//表格
type Table struct {
	Header []*Cell
	Rows   []*Row
}

//行
type Row []*Cell

//单元格
type Cell struct {
	Content  Content
	Align    CellAlign
	FontType FontType
}
