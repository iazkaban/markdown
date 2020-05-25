package models

type LanguageType uint16

const (
	LanguageTypeUnknown    LanguageType = iota //未知语言
	LanguageTypeGolang                         //golang
	LanguageTypeC                              //c
	LanguageTypeCPlus                          //c++
	LanguageTypeJava                           //java
	LanguageTypeJavaScript                     //javascript
	LanguageTypeCSharp                         //c#
	LanguageTypeRust                           //rust
)

type CodeBlock struct {
	LanguageType LanguageType //语言编号
	LanguageName string       //语言名称（其实是```之后跟的字符串，即使是Unknown，也可能会有这个字段内容）
	Content      Content      //内容
}
