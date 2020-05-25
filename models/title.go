package models

//标题(对应HTML的h1/h2/h3/h4/h5/h6)
type Title struct {
	Level   uint8   //级别，数字越大，字体越小
	Content Content //内容
}
