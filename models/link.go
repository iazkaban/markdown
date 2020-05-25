package models

type Target uint8

const (
	TargetUnknown = iota
	TargetBlank
	TargetSelf
)

type Link struct {
	Title   string
	Link    string
	Content string
	Target  Target
}
