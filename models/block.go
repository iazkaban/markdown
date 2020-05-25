package models

type Block struct {
	Body  string
	Block *Block
}
