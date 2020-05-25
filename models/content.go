package models

type FontType int16

const (
	FontTypeNone        FontType = 0         //无特殊字体
	FontTypeItalics     FontType = 1 << iota //斜体
	FontTypeBlod                             //粗体
	FontTypeDeleted                          //删除线
	FontTypeUnderLine                        //下划线
	FontTypeSuperScript                      //上标
	FontTypeSubScript                        //下标
)

//内容
type Content struct {
	Content  string   //内容
	FontType FontType //字体设定
}
