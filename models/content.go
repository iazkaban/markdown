package models

type FontType uint8

const (
	FontTypeUnkown = iota
	FontTypeI      //斜体
	FontTypeBlod   //粗体
	FontTypeIBlod  //粗体+斜体
)

type Contest struct {
	Body       string
	FontType   FontType
	IsDeleted  bool //删除线
	IsDownLine bool //下划线
	IsUpTag    bool //是否是上标
	IsDownTag  bool //是否是下标
}
