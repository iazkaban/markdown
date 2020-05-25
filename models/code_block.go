package models

type LanguageType uint16

const (
	LanguageTypeUnknown = iota
	LanguageTypeGolang
	LanguageTypeC
	LanguageTypeCPlus
	LanguageTypeJava
	LanguageTypeJavaScript
	LanguageTypeCSharp
	LanguageTypeRust
)

type CodeBlock struct {
	LanguageType LanguageType
	LanguageName string
	Body         string
}
